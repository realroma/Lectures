// Package crawler реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package crawler

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Crawler представляет собой поисковый робот для обхода сайтов
type Crawler struct {
	Url   string
	Depth int
}

// New создает новый экземпляр типа Crawler с указанными параметрами
func New(url string, depth int) *Crawler {
	c := Crawler{
		Url:   url,
		Depth: depth,
	}
	return &c
}

// Scan осуществляет рекурсивный обход ссылок сайта, указанного в Url,
// с учётом глубины перехода по ссылкам, указанной в Depth.
func (c *Crawler) Scan() (data map[string]string, err error) {
	data = make(map[string]string)
	parse(c.Url, c.Url, c.Depth, data)
	return data, nil
}

// parse рекурсивно обходит ссылки на странице, переданной в url.
// Глубина рекурсии задаётся в depth.
// Каждая найденная ссылка записывается в ассоциативный массив
// data вместе с названием страницы.
func parse(url, baseurl string, depth int, data map[string]string) error {
	if depth == 0 {
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		return err
	}
	page, err := html.Parse(response.Body)
	if err != nil {
		return err
	}

	data[url] = pageTitle(page)
	links := pageLinks(nil, page)
	for _, link := range links {
		if data[link] == "" && strings.HasPrefix(link, baseurl) {
			parse(link, baseurl, depth-1, data)
		}
	}

	return nil
}

// pageTitle осуществляет рекурсивный обход HTML-страницы и возвращает значение элемента <tittle>.
func pageTitle(n *html.Node) string {
	var title string
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title = pageTitle(c)
		if title != "" {
			break
		}
	}
	return title
}

// pageLinks рекурсивно сканирует узлы HTML-страницы и возвращает все найденные ссылки без дубликатов.
func pageLinks(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if !sliceContains(links, a.Val) {
					links = append(links, a.Val)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = pageLinks(links, c)
	}
	return links
}

// sliceContains возвращает true если массив содержит переданное значение
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

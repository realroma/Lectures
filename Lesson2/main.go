package main

import (
	"flag"
	"fmt"
	"strings"
	"text/scanner"

	"project/Lectures/Lesson2/crawler"
)

func pri(m map[string]string) {
	for i, v := range m {
		fmt.Println("Key: ", i)
		fmt.Println("Hash: ", v)
	}
}

// В чём суть?: Мы принимаем любое количество значений типа map в виде списка. Распаковываем его и все данные посредством цикла заливаем в новую оболочку.
func add(m ...map[string]string) map[string]string {
	mr := make(map[string]string)
	for _, n := range m {
		for i, v := range n {
			mr[i] = v
		}
	}
	return mr
}

func er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Sc(m map[string]string, w string) {
	var r scanner.Scanner
	for _, v := range m {
		r.Init(strings.NewReader(v))
		for tok := r.Scan(); tok != scanner.EOF; tok = r.Scan() {
			if r.TokenText() == w {
				fmt.Println(v)
			}
		}
	}
}

func parseFlag() string {
	f := flag.String("s", "", "Ссылка на сайт о котором надо получить информацию.") //Создаём ссылку на флаг(Типо -g при использовании команд.) s название флага, "" - значение по умолчанию, комментарий.
	flag.Parse()                                                                    //Без этого флаг будет выдаватся по умолчанию.
	return *f
}
func main() {
	word := parseFlag()

	a, err := crawler.Scan("https://go.dev", 3)

	er(err)
	b, err := crawler.Scan("http://habr.com", 2)
	er(err)
	c, err := crawler.Scan("https://html5book.ru/hyperlinks-in-html/", 4)
	er(err)
	m := add(a, b, c)
	pri(m)
	Sc(m, word)
}

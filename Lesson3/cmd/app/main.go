package main

import (
	"flag"
	"fmt"

	"project/Lectures/Lesson3/pkg/crawler"
	"project/Lectures/Lesson3/pkg/sorter"
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

func parseFlag() string {
	f := flag.String("s", "", "Ссылка на сайт о котором надо получить информацию.") //Создаём ссылку на флаг(Типо -g при использовании команд.) s название флага, "" - значение по умолчанию, комментарий.
	flag.Parse()                                                                    //Без этого флаг будет выдаватся по умолчанию.
	return *f
}
func main() {
	word := parseFlag()

	a := crawler.New("https://go.dev", 3)
	b := crawler.New("http://habr.com", 2)
	c := crawler.New("https://html5book.ru/hyperlinks-in-html/", 4)
	ma, err := a.Scan()
	er(err)
	mb, err := b.Scan()
	er(err)
	mc, err := c.Scan()
	er(err)

	m := add(ma, mb, mc)
	pri(m)
	sorter.Sorter(m, word)

}

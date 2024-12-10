package main

import (
	"flag"
	"fmt"

	"project/Lectures/Lesson2/crawler"
)

func pri(m map[string]string) {
	for i, v := range m {
		fmt.Println("Key: ", i)
		fmt.Println("Hash: ", v)
	}
}

func add(m map[string]string, m2 map[string]string) map[string]string {
	for i, v := range m2 {
		_, b := m[v]
		if b == false {
			m[i] = m2[v]
			m[i] = m2[i]
		}
	}
	return m
}

func er(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func parseFlag() {
	f := flag.String("s", "", "Ссылка на сайт о котором надо получить информацию.") //Создаём ссылку на флаг(Типо -g при использовании команд.) s название флага, "" - значение по умолчанию, комментарий.
	flag.Parse()                                                                    //Без этого флаг будет выдаватся по умолчанию.
	fmt.Println(*f)
}
func main() {
	parseFlag()
	//r := bufio.NewReader()
	//fmt.Println(r)
	a, err := crawler.Scan("https://go.dev", 3)
	er(err)
	b, err := crawler.Scan("http://habr.com", 2)
	er(err)
	c, err := crawler.Scan("https://html5book.ru/hyperlinks-in-html/", 4)
	er(err)
	m := add(a, b)
	m2 := add(m, c)
	pri(m2)
}

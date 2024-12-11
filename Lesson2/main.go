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

func parseFlag() {
	f := flag.String("s", "", "Ссылка на сайт о котором надо получить информацию.") //Создаём ссылку на флаг(Типо -g при использовании команд.) s название флага, "" - значение по умолчанию, комментарий.
	flag.Parse()                                                                    //Без этого флаг будет выдаватся по умолчанию.
	fmt.Println(*f)
}
func main() {
	parseFlag()
	a, err := crawler.Scan("https://go.dev", 3)
	er(err)
	b, err := crawler.Scan("http://habr.com", 2)
	er(err)
	c, err := crawler.Scan("https://html5book.ru/hyperlinks-in-html/", 4)
	er(err)
	m := add(a, b, c)
	// r := bufio.NewScanner()
	// fmt.Println(r)
	pri(m)
}

package main

import (
	"fmt"
	"flag"

	"project/Lectures/crawler"
)

func pri(m map[string]string) {
	for i, v := range m {
		fmt.Println("Key: ",i)
		fmt.Println("Hash: ",v)
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

func main() {
	f := flag.String("s", "",  "Ссылка на сайт о котором надо получить информацию.")//Создаём ссылку на флаг(Типо -g при использовании команд.) s название флага, "" - значение по умолчанию, комментарий.
	flag.Parse()//Без этого флаг будет выдаватся по умолчанию.
	fmt.Println(*f)
	a, err := crawler.Scan("https://go.dev", 3)
	if err != nil {
		fmt.Println(err)
	}
	b, err := crawler.Scan("http://habr.com", 2)
	if err != nil {
		fmt.Println(err)
	}
	m := add(a, b)
	pri(m)
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"project/Lectures/Lesson5/pkg/crawler"
	"project/Lectures/Lesson5/pkg/filer"
	"project/Lectures/Lesson5/pkg/index"
	"project/Lectures/Lesson5/pkg/sorter"
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

func Scan(m map[string]string) map[string]string {
	// Получаем сслки.
	a := crawler.New("https://go.dev", 1)
	b := crawler.New("http://habr.com", 2)
	c := crawler.New("https://html5book.ru/hyperlinks-in-html/", 3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	ma, err := a.Scan()
	er(err)
	mb, err := b.Scan()
	er(err)
	mc, err := c.Scan()
	er(err)
	m = add(ma, mb, mc)
	pri(m)
	return m
}

func main() {
	//Делаем флаг для поиска по словам в консоли.
	word := parseFlag()
	var m map[string]string

	//Объединяем ссылки в единое.
	//Тут наверное надо переписать на ссылки чтобы не задействовало много памяти в будущем.

	file := filer.New("")
	filer.OpenFile()
	fmt.Println(file.Name())

	m = filer.Read(file)
	for i, v := range m {
		fmt.Println("i:", i)
		fmt.Println("v:", v)
	}

	byteMap, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	filer.Write(file, byteMap)

	itd := index.Indexer(m)
	fmt.Println(itd)

	sorter.Sorter(m, word)
	if word != "" {
		index.Read(word, itd)
	}
}

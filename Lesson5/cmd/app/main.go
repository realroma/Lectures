package main

import (
	"encoding/json"
	"fmt"
	"log"

	"project/Lectures/Lesson5/pkg/crawler"
	"project/Lectures/Lesson5/pkg/filer"
	"project/Lectures/Lesson5/pkg/index"
	"project/Lectures/Lesson5/pkg/lecture_flag"
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
		log.Fatal(err)
	}
}

func Scan(m map[string]string) map[string]string {
	// Получаем сслки.
	fmt.Println("Site")
	a := crawler.New("https://go.dev", 1)
	b := crawler.New("http://wikipedia.org", 1)
	c := crawler.New("https://html5book.ru/hyperlinks-in-html/", 3)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	fmt.Println("Add")
	ma, err := a.Scan()
	fmt.Println(m)
	er(err)
	mb, err := b.Scan()
	fmt.Println(m)
	er(err)
	mc, err := c.Scan()
	er(err)
	m = add(ma, mb, mc)
	pri(m)
	return m
}

type s struct {
	field  string
	field2 string
}

func main() {
	//Делаем флаг для поиска по словам в консоли.
	word := lecture_flag.ParseFlag()
	var site map[string]string

	//Объединяем ссылки в единое.
	name := "Link.txt"
	filer.New(name)

	f, err := filer.OpenFile(name)
	fmt.Println(f.Name())
	er(err)

	fmt.Println("Scan")
	moll := Scan(site)

	fmt.Println("Marshal")
	b, err := json.Marshal(moll)
	er(err)

	fmt.Println("Write")
	f.Write(b)

	fmt.Println("Read")
	_, err = filer.Read(f, b)

	f.Read(b)

	// fmt.Println("Read b:", b)

	m := make(map[string]string)
	err = json.Unmarshal(b, &m)
	er(err)
	fmt.Println(m)

	itd := index.Indexer(moll)
	fmt.Println(itd)

	sorter.Sorter(moll, word)
	if word != "" {
		index.Read(word, itd)
	}
}

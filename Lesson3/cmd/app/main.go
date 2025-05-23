package main

import (
	"fmt"

	"project/Lectures/Lesson3/pkg/crawler"
	"project/Lectures/Lesson3/pkg/index"
	"project/Lectures/Lesson3/pkg/lecture_flag"
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
		fmt.Println(n)
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

func main() {
	//Делаем флаг для поиска по словам в консоли.
	word := lecture_flag.ParseFlag()

	//Получаем сслки.
	a := crawler.New("https://go.dev", 2)
	b := crawler.New("http://habr.com", 2)
	c := crawler.New("https://html5book.ru/hyperlinks-in-html/", 2)
	fmt.Println(a, b, c)
	//ma, err := crawler.New("https://go.dev", 2).Scan()
	ma, err := a.Scan()
	er(err)
	mb, err := b.Scan()
	er(err)
	mc, err := c.Scan()
	er(err)
	fmt.Println(ma, mb, mc)

	//Объединяем ссылки в единое.

	//Тут наверное надо переписать на ссылки чтобы не задействовало много памяти в будущем.
	m := add(ma, mb, mc)
	pri(m)

	//Тут та же проблема. Пустой массив. Надо доделать.
	itd := index.Indexer(m)

	sorter.Sorter(m, word)
	if word != "" {
		index.Read(word, itd)
	}
}

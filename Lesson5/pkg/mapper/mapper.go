package mapper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"project/Lectures/Lesson5/pkg/crawler"
)

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

func pri(m map[string]string) {
	for i, v := range m {
		fmt.Println("Key: ", i)
		fmt.Println("Hash: ", v)
	}
}

func Scan() map[string]string {
	// Получаем сслки.
	a := crawler.New("https://go.dev", 1)
	b := crawler.New("http://wikipedia.org", 1)
	c := crawler.New("https://html5book.ru/hyperlinks-in-html/", 3)
	ma, err := a.Scan()
	if err != nil {
		log.Fatal(err)
	}
	mb, err := b.Scan()
	if err != nil {
		log.Fatal(err)
	}
	mc, err := c.Scan()
	if err != nil {
		log.Fatal(err)
	}
	m := add(ma, mb, mc)
	pri(m)
	return m
}

func MapWrite(m map[string]string, name string) {

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(name, b, os.FileMode(os.O_WRONLY))
	if err != nil {
		log.Fatal(err)
	}
}

func MapRead(name string) map[string]string {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}

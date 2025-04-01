package filer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func New(path string) *os.File {
	if path == "" {
		path = "Link.txt"
	}
	f, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func CheckFile() *[]byte {
	var b []byte
	f, err := os.OpenFile("Link.txt", os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g, err := f.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(g)

	return &b
}

// Возвращать он ничего не должен.
func Write(f *os.File, b []byte) {
	err := os.WriteFile(f.Name(), b, 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func Read(f *os.File) map[string]string {
	//Читаем файл и получаем массив байт.
	file, err := os.ReadFile(f.Name())
	if err != nil {
		log.Fatal()
	}
	fmt.Println(file)
	//Создаём переменную в которую данные поместятся.
	m := make(map[string]string)

	err = json.Unmarshal(file, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

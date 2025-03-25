package filer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Message struct {
	Link string
}

func CheckSave() {
	f, err := os.OpenFile("Link.txt", os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//os.WriteFile()
}

func WriteFile() string {
	m := map[string]string{
		"Some": "value",
	}
	bm, _ := json.Marshal(m)

	err := os.WriteFile("Link.txt", bm, 0666)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := os.ReadFile("Link.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
	return string(bytes)
}

func ReadFile() map[string]string {
	//Читаем файл и получаем массив байт.
	rw, err := os.ReadFile("Link.txt")
	if err != nil {
		log.Fatal(err)
	}

	//Создаём переменную в которую данные поместятся.
	m := make(map[string]string)

	err = json.Unmarshal(rw, &m)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

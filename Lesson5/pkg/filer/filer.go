package filer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func CheckSave() {
	f, err := os.OpenFile("Link.txt", os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//os.WriteFile()
}

func WriteFile() string {
	m := map[string]int{
		"Some": 1,
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

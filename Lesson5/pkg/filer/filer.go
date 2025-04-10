package filer

import (
	"log"
	"os"
)

func New(name string) *os.File {
	gw, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Реализация если используется vsl. Сделать потом проверку для рабочей директории
	err = os.Chdir(gw + "/Lectures/Lesson5/cmd/app")
	if err != nil {
		log.Fatal(err)
	}

	if name == "" {
		name = "Link.txt"
	}

	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

package filer

import (
	"fmt"
	"io"
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

func OpenFile(s string) (f *os.File, err error) {
	return os.OpenFile(s, os.O_WRONLY, 0666)
}

func Write(w io.Writer, b []byte) (i int, err error) {
	fmt.Println("Writing")
	return w.Write(b)
}

func Read(r io.Reader, b []byte) (i int, err error) {
	return r.Read(b)
}

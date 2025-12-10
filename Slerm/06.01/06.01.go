package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"project/module/server"
)

func EnvCreate() {
	os.Setenv("APP_LOGFILE_PATH", "Log/log.txt")
}

func Find() string {
	_, was := os.LookupEnv("APP_LOGFILE_PATH")
	if was == false {
		os.Setenv("APP_LOGFILE_PATH", "log.txt")
	}
	return os.Getenv("APP_LOGFILE_PATH")
}

func main() {
	EnvCreate()
	fmt.Println(Find())
	Server()
}

func Write(text string) {
	file, _ := os.Create(os.Getenv("APP_LOGFILE_PATH"))
	defer file.Close()
	file.WriteString(text)
	file.Sync()
}

func Post(w http.ResponseWriter, req *http.Request) {
	scanner := bufio.NewScanner(req.Body)
	scanner.Scan()
	a := scanner.Text()
	Write(a)
	fmt.Println(a)
}

func Server() {
	http.HandleFunc("/", Post)
	server.HTTPServer()
}

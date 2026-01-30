package main

import (
	"fmt"
	"net/http"
	"os"
	"project/Lectures/Slerm/03.03/server"
)

func json(w http.ResponseWriter, req *http.Request) {
	file, _ := os.ReadFile("DataBase/db.json")
	a := string(file)
	fmt.Println(a)
	fmt.Fprintf(w, a)
}

func main() {
	http.HandleFunc("/health", json)
	server.HTTPServer()
}

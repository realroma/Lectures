package main

import (
	hello "Project/Lectures/ProdvinutiyGolang/Archetecture/internal"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}

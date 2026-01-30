package main

import (
	"Project/Lectures/ProdvinutiyGolang/Archetecture/configs"
	hello "Project/Lectures/ProdvinutiyGolang/Archetecture/internal"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}

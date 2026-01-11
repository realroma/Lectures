package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func getStatus(codeCh chan string) {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		log.Fatal(err)
	}
	codeCh <- resp.Status
}

func main() {
	code := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getStatus(code)
			wg.Done()
		}()

	}
	go func() {
		wg.Wait()
		close(code)
	}()
	for res := range code {
		fmt.Printf("Code: %v \n", res)
	}
}

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type Upstruc struct {
	Status     string
	Service_id string
	Checks     checks
}

type checks struct {
	Ping_mysql ping_mysql
}

type ping_mysql struct {
	Component_id   string
	Component_type string
	Status         string
}

func main() {
	url := flag.String("url", "", "Path to server")
	flag.Parse()
	HTTPClient(*url)
}

func unmarshal(text string) {
	var link Upstruc
	json.Unmarshal([]byte(text), &link)
	fmt.Printf("Overall status is %v, with service_id %v mesql component is %v\n", link.Status, link.Service_id, link.Checks.Ping_mysql.Status)
}

func HTTPClient(url string) {
	resp, _ := http.Get(url)
	scanner := bufio.NewScanner(resp.Body)
	scanner.Scan()
	a := scanner.Text()
	unmarshal(a)
}

//How to use? Write with "-url http://localhost:8080/health".

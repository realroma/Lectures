package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ansver := 0
	numGorutines := 3
	ch := make(chan int, numGorutines)

	need := len(arr) / numGorutines

	for i := 0; i < numGorutines; i++ {
		firstId := i * need
		second := firstId + need
		go calculate(arr[firstId:second], ch)
	}

	for i := 0; i < numGorutines; i++ {
		ansver += <-ch
	}
	fmt.Println(ansver)
}

func calculate(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

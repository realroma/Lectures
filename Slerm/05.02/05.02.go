package main

import (
//	"errors"
	"fmt"
	"time"
)

func write(ch chan<- int, t *time.Timer) {
	fmt.Println("Writing in chanel")
	time.Sleep(1 * time.Second)
	select {
	case <-t.C:
		fmt.Println("Func is stop.")
	default:
		for i := 0; i <= 3; i++ { //Записываем в небуферизированный канал значения.
		ch <- i
		}
	}
	close(ch) //
}

func read(i int) {
	fmt.Printf("Done %v.\n", i)
}

func MyFunc () {
	ch := make(chan int)
	timer := time.NewTimer(1 * time.Second)
	go write(ch, timer)
	for i := range ch {
		go read(i)
	}
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("MyFunc is starting")
	MyFunc()
}

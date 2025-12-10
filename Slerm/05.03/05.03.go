package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func went(ctx context.Context, i int) {
	time.Sleep(time.Duration(i) * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("Time is out.")
	default:
		fmt.Println("Ready.")
	}
}

func main() {
	parent := context.Background()
	ctx, err := context.WithTimeout(parent, 1500*time.Millisecond)
	if err != nil {
		log.Fatal(err)
	}
	go went(ctx, 1)
	go went(ctx, 2)
	fmt.Println("Wait")
	time.Sleep(3 * time.Second)
	fmt.Println("Main is ends.")
}

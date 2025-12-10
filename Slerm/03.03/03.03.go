package main

import "fmt"

func main () {
	var slise []string = []string{"1", "3", "2", "4"}
	var slise2 []string
	for i := range slise {
		if slise2 != nil {
			for f := range slise2 {
				fmt.Println(slise [i], slise[i] > slise2[f], slise[f])
			}
		}
		slise2 = append(slise2, slise[i])
	}
	fmt.Println(slise, slise2)
}

package main

import (
	"fmt"
	)

func main () {
	var str string = "съешь ещё этих мягких французких булок, да выпей чаю"
	var maps map[string]int = map[string]int{}
	s := []rune(str)
	for _, f := range s {
		g := string(f)
		switch {
			case g == " ":
			case g == ",":
			default:
				maps[string (f)]++
		}
	}
	fmt.Println(maps)
}

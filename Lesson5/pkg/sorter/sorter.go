package sorter

import (
	"fmt"
	"strings"
	"text/scanner"
)

func Sorter(m map[string]string, w string) {
	for _, v := range m {
		arr := Sort(v)
		for _, v := range arr {
			if v == w {
				fmt.Println(v)
			}
		}
	}
}

func Elements(arr1 []int, arr2 []int) []int {
	fmt.Println("Arr: ", arr1)
	fmt.Println("Arr2: ", arr2)
	for _, v := range arr1 {
		for _, v2 := range arr2 {
			if v != v2 {
				arr1 = append(arr1, v2)
			}
		}
	}
	if len(arr1) < 1 {
		return nil
	}
	return arr1
}

func Sort(str string) []string {
	var r scanner.Scanner
	var arr []string
	r.Init(strings.NewReader(str))
	for tok := r.Scan(); tok != scanner.EOF; tok = r.Scan() {
		arr = append(arr, r.TokenText())
	}
	return arr
}

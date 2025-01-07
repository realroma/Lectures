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

func Sort(str string) []string {
	var r scanner.Scanner
	var arr []string
	r.Init(strings.NewReader(str))
	for tok := r.Scan(); tok != scanner.EOF; tok = r.Scan() {
		arr = append(arr, r.TokenText())
	}
	return arr
}

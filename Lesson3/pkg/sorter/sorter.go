package sorter

import (
	"fmt"
	"strings"
	"text/scanner"
)

func Sorter(m map[string]string, w string) {
	var r scanner.Scanner
	for _, v := range m {
		r.Init(strings.NewReader(v))
		for tok := r.Scan(); tok != scanner.EOF; tok = r.Scan() {
			if r.TokenText() == w {
				fmt.Println(v)
			}
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	t.Content = strings.ReplaceAll(t.Content, "  ", " ")

	var message string

	r := []rune(t.Content)

	for i, v := range r {
		//Закрывашка на случаи первого и последнего элемента.
		if v == '-' && i != 0 && i != len(r)-1 {
			fmt.Println(i, "", len(r))
			r[i-1], r[i+1] = r[i+1], r[i-1]
			t.Content = string(r)
		}
	}

	for _, v := range t.Content {
		switch {
		case v == '-':
			message = (message + "")
		case v == '+':
			message = (message + "!")
		default:
			message = (message + string(v))
		}
	}

	t.Content = message

	var integer [10]rune = [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var intbuf int

	for _, v := range t.Content {
		for i, va := range integer {
			if v == va {
				t.Content = strings.Replace(t.Content, string(va), "", 1)
				intbuf = intbuf + i
			}
		}
	}

	if intbuf != 0 {
		t.Content = t.Content + " " + strconv.Itoa(intbuf)
	}

	fmt.Println(t.Content)

}

func main() {
	text := &Text{}
	fmt.Println("Ведите строку:")
	scanner := bufio.NewScanner(os.Stdin)

	//	fmt.Println("Ведите строку:.")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}

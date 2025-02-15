// Реализуация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		//Здесь присзодит замыкание кода, если он имеет ссылку на себя.
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	e := l.root.next
	//Меняем начало списка на последующий элемент.
	l.root.next = e.next
	l.root.prev = e
	//Стираем связь с последующим элементом если он есть.
	if e.next != nil {
		e.next = e.prev
	}
	return l
}

func Swap(e *Elem) {
	next := e.next
	prev := e.prev
	//fmt.Printf("befor\n next: %v,\n prev: %v\n", next, prev)
	e.prev = next
	e.next = prev
	//fmt.Printf("after\n next: %v,\n prev: %v\n", e.next, e.prev)
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	Swap(l.root)
	l.root = l.root.prev
	Swap(l.root)
	l.root = l.root.prev
	Swap(l.root)
	l.root = l.root.prev
	Swap(l.root)
	e := l.root.prev
	e.next = l.root
	l.root = e

	return l
	//Почему тут не выдаётся значений? скорее всего после входа в нулевой поинтер он не может выйти. Как решить? Сделать функцию с особым буфером,
	//в котором при появлении в следующем блоке значиния nil будет условие и самозамыкание.
}

//Тут я столкнулся с ситуацией самозамыкания и бесконечного цикла. У Element есть ссылка на ячейку 0xc000100060 которую он ставит в начале и конце списка.
//Как только мы попадаем в эту ячейку он постоянно переходит в себя. Образуется бесконечный цикл.

//Я не могу найти решения проблемы пустого листа. В моих мыслях был алгоритм по типу:
//1)Входим в первый элемент через element = l.root.next
//2)Меняем значения  l.root.next и l.root.prev местами используя буферную переменную. Но кажется, я не правильно вставлял в буфер значение.
//3)Входим в следующие элементы используя l.root = l.root.prev
//4)Как только получаем значение l.Val == nil выходим из листа и меняем местами последние значения.

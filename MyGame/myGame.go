package main

import (
	"bufio"
	"fmt"
	"os"
)

type Player struct {
	Balance int
	Buisnes *Busines
}

type Busines struct {
	Name  string
	Pay   int
	Cost  int
	Grade string
}

func NewP() Player {
	b := DefaultB()
	var p Player = Player{0, b}
	return p
}

// Поинтеры нужны для того чтобы при изменении баланса у структуры, в мапе они так же изменялись, и не подвергались копированию.
func (p *Player) Buy(b *Busines) {
	if p.Balance >= b.Cost {
		p.Balance, p.Buisnes = p.Balance-b.Cost, b
		fmt.Println("Sucksess")
	}
	fmt.Println("Dont enough money")
}

func (p *Player) Work() {
	p.Balance++
}

func (p *Player) Pay() *Player {
	p.Balance = p.Balance + p.Buisnes.Pay
	return p
}

func NewB(name string, cost int, pay int) *Busines {
	b := Busines{Name: name, Cost: cost, Pay: pay, Grade: ""}
	return &b
}

func DefaultB() *Busines {
	if DefaultBusines.Name == "" {
		return NewB("Farm", 20, 2)
	}
	return &DefaultBusines
}

var DefaultBusines Busines

func main() {
	//Создаём бизнес в доступе.
	p := NewP()
	b := DefaultB()
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write command.")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "stop" {
			break
		} else if text == "work" {
			fmt.Println("before balance ", p.Balance)
			p.Work()
			fmt.Println("after balance ", p.Balance)
		} else if text == "skip" {
			p.Pay()
			fmt.Println("Day left. Your pay on come.")
			fmt.Println(p.Balance)
		}
	}
}

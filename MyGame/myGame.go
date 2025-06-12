package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func NewP() Player {
	b := DefaultB()
	var p Player = Player{0, b, time.Now()}
	return p
}

type Player struct {
	Balance  int
	Buisnes  *Busines
	Previous time.Time
}

func (p *Player) Work() {
	p.Balance++
}

func (p *Player) Pay() {
	now := time.Now()
	if now.Sub(p.Previous) > 20 {
		p.Balance = p.Balance + p.Buisnes.Pay
		p.Previous = now
	} else {
		fmt.Println("Time less than 10.")
	}
	return
}

// Поинтеры нужны для того чтобы при изменении баланса у структуры, в мапе они так же изменялись, и не подвергались копированию.
func (p *Player) Buy(b *Busines) {
	if p.Balance >= b.Cost {
		p.Balance, p.Buisnes = p.Balance-b.Cost, b
		fmt.Println("Sucksess")
	}
	fmt.Println("Dont enough money")
}

func NewB(name string, cost int, pay int) *Busines {
	b := Busines{Name: name, Cost: cost, Pay: pay, Grade: ""}
	return &b
}

type Busines struct {
	Name  string
	Pay   int
	Cost  int
	Grade string
}

func DefaultB() *Busines {
	if DefaultBusines.Name == "" {
		return NewB("Farm", 20, 2)
	}
	return &DefaultBusines
}

var DefaultBusines Busines

func handler(scan *bufio.Scanner, p *Player) {
	for scan.Scan() {
		text := strings.ToLower(scan.Text())
		if text == "exit" {
			fmt.Println("Exiting..")
			return
		} else if text == "work" {
			fmt.Println("Player balance before: ", p.Balance)
			p.Work()
			fmt.Println("Player balance: ", p.Balance)
		} else if text == "skip" {
			fmt.Println("Player balance today: ", p.Balance)
			p.Pay()
			fmt.Println("Balance player: ", p.Balance)
		} else {
			fmt.Println("Wrong command.")
		}
	}
}

func main() {
	//Создаём бизнес в доступе.
	p := NewP()
	b := DefaultB()
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write command.")

	handler(scanner, &p)

	fmt.Println(p.Balance, "\n", &p.Buisnes)
}

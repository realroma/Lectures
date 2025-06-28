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
	var p Player = Player{0, *b, time.Now()}
	return p
}

type Player struct {
	Balance  int
	Buisnes  Busines
	Previous time.Time
}

func (p *Player) Work() {
	p.Balance++
}

func (p *Player) Pay() {
	//Узнаём время сейчас.
	now := time.Now()

	//Получаем разницу между сейчас и прошлым разом.
	diffrent := now.Sub(p.Previous)

	hours, _ := time.ParseDuration("10s")
	// hours, _ := time.ParseDuration("1h")
	if diffrent > hours {
		fmt.Println("Diffrent: ", diffrent)
		fmt.Println("Diffrent Round: ", diffrent.Seconds())
		p.Balance = p.Balance + p.Buisnes.Pay*int(diffrent.Seconds())
		p.Previous = now
		fmt.Println("Time: ", now)
	} else {
		fmt.Println("Ten seconds don't has been left.")
	}
}

// Поинтеры нужны для того чтобы при изменении баланса у структуры, в мапе они так же изменялись, и не подвергались копированию.
func (p *Player) Buy(b Busines) {
	if p.Balance >= b.Cost {
		p.Balance, p.Buisnes = p.Balance-b.Cost, b
		fmt.Println("Sucksess")
	}
	fmt.Println("Dont enough money")
}

func NewB(name string, cost int, pay int, grade string) *Busines {
	b := Busines{Name: name, Cost: cost, Pay: pay, Grade: grade}
	return &b
}

type Busines struct {
	Name  string
	Pay   int
	Cost  int
	Grade string
}

func (b *Busines) ChangeGrade(grade string) {
	b.Grade = grade
}

func DefaultB() *Busines {
	if DefaultBusines.Name == "" {
		return NewB("Farm", 20, 2, "None")
	}
	return &DefaultBusines
}

var DefaultBusines Busines

func handler( /*parent context.Context,*/ scan *bufio.Scanner, p *Player) {
	// ctx, cancel := context.WithTimeout(parent, 10 * time.Second)
	// Поступило предложение сделать всё через behavior tree. Оно будет работать через контекст. Интересное предположение.
	// Ещё заново перебрать контексты.
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
		} else if text == "grade" {
			fmt.Println("Widtch grade do you want?")
			WriteGrade(p)
			p.Buisnes.ChangeGrade("Not none.")
		} else {
			fmt.Println("Wrong command.")
		}
	}
}

func WriteGrade(p *Player) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if sc.Text() == "exit" {
			fmt.Println("Exiting..")
			return
		} else if sc.Text() == "home" {
			p.Buisnes.ChangeGrade("home")
			fmt.Println(p.Buisnes)
			return
		} else {
			fmt.Println("Wrong command.")
			return
		}
	}
}

func main() {
	//Создаём бизнес в доступе.
	// parent := context.Background()
	p := NewP()
	b := DefaultB()
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Write command.")

	handler( /*parent,*/ scanner, &p)

	fmt.Println(p.Balance, "\n", p.Buisnes)
}

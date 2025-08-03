package player

import (
	"fmt"
	b "project/Lectures/MyGame/internal/buisnes"

	"time"
)

type Player struct {
	Id       int       //`json: "id"`
	Balance  int       //`json: "balance"`
	Buisnes  b.Buisnes //`json: "busines"`
	Previous time.Time //`json: "previous"`
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
		p.Balance = p.Balance + (p.Buisnes.Pay * int(diffrent.Seconds()))
		p.Previous = now
		fmt.Println("Time: ", now)
	} else {
		fmt.Println("Ten seconds don't has been left.")
	}
}

// Поинтеры нужны для того чтобы при изменении баланса у структуры, в мапе они так же изменялись, и не подвергались копированию.
func (p *Player) Buy(buisnes b.Buisnes) {
	if p.Balance >= buisnes.Cost {
		p.Balance, p.Buisnes = p.Balance-buisnes.Cost, buisnes
		fmt.Println("Sucksess")
	}
	fmt.Println("Dont enough money")
}

func PPB() {
	b := b.New("None", 10, 1, "none")
	fmt.Println(b)
}

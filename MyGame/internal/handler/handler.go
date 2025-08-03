package handler

import (
	"bufio"
	"fmt"
	"os"
	p "project/Lectures/MyGame/internal/player"
	"strings"
)

func textFromConsole() string {
	fmt.Println("Write command.")
	sc := bufio.NewScanner((os.Stdin))
	if sc.Scan() {
		return strings.ToLower(sc.Text())
	}
	return ""
}

func menu(text string, p *p.Player) {
	switch {
	case text == "work":
		fmt.Println("Player balance before: ", p.Balance)
		p.Work()
		fmt.Println("Player balance after: ", p.Balance)
		return
	case text == "exit" || text == "":
		return
	case text == "skip":
		fmt.Println("Player balance: ", p.Balance)
		p.Pay()
		fmt.Println("Player balance today: ", p.Balance)
		return
	case text == "grade":
		additionalMenu(textFromConsole(), p)
	default:
		return
	}
}

func additionalMenu(text string, p *p.Player) {
	switch {
	case text == "exit" || text == "":
		return
	case text == "home":
		fmt.Println("Player: ", p, " Want chanhe his: ", p.Buisnes)
		p.Buisnes.ChangeGrade("home")
		fmt.Println("Player: ", p, " Change business: ", p.Buisnes)
	default:
		return
	}
}

func Handler(p *p.Player) {
	menu(textFromConsole(), p)
}

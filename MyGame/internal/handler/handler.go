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
		fmt.Println("Can't found any coincidence.")
		return
	}
}

func additionalMenu(text string, p *p.Player) {
	fmt.Println("Open menu grade for: ", p.Id)
	switch {
	case text == "exit" || text == "":
		return
	case text == "home":
		fmt.Println("Player: ", p, " Want change this: ", p.Buisnes)
		p.Buisnes.ChangeGrade("home")
		fmt.Println("Player: ", p, " Changed business: ", p.Buisnes)
	case text == "cancel":
		fmt.Println("Canceled.")
	default:
		fmt.Println("Can't found any coincidence.")
		return
	}
}

func Handler(p *p.Player) {
	menu(textFromConsole(), p)
}

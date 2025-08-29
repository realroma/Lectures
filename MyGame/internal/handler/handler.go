package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"project/Lectures/MyGame/internal/database"
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

func menu(text string, p *p.Player, db *sql.DB) int8 {
	switch {
	case text == "balance":
		request, err := db.Exec("SELECT balance FROM users WHERE id = $1", p.Id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("DB: ", request)
	case text == "work":
		fmt.Println("Player balance before: ", p.Balance)
		p.Work()
		// request, err := db.Exec("UPDATE TABLE users WHERE id = $1 VALUE balance = $2", p.Id, p.Balance)
		fmt.Println("Player balance after: ", p.Balance)
	case text == "exit" || text == "":
		return 1
	case text == "skip":
		fmt.Println("Player balance: ", p.Balance)
		p.Pay()
		fmt.Println("Player balance today: ", p.Balance)
	case text == "grade":
		additionalMenu(textFromConsole(), p)
	default:
		fmt.Println("Can't found any coincidence.")
		return 1
	}
	return 0
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

func Handler(p *p.Player, db *sql.DB) {
	var exitStatus int8 = 0
	for exitStatus == 0 {
		exitStatus = menu(textFromConsole(), p, db)
	}
}

func Login(db *sql.DB, id int) *p.Player {
	fmt.Println(db, id)
	database.Select(db, id)
	p := p.New(id)
	return p
}

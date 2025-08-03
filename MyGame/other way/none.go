package none

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	b "project/Lectures/MyGame/internal/buisnes"
	p "project/Lectures/MyGame/internal/player"

	bt "github.com/joeycumines/go-behaviortree"
)

func GetText() string {
	sc := bufio.NewScanner((os.Stdin))
	if sc.Scan() {
		return strings.ToLower(sc.Text())
	}
	return ""
}

var (
	work = func(p *p.Player) bt.Node {
		return bt.New(
			func(children []bt.Node) (bt.Status, error) {
				fmt.Println("Player balance before: ", p.Balance)
				p.Work()
				return bt.Success, nil
			},
		)
	}
)

func skip(p *p.Player) bt.Node {
	return bt.New(
		func(children []bt.Node) (bt.Status, error) {
			fmt.Println("Player balance: ", p.Balance)
			p.Pay()
			fmt.Println("Player balance today: ", p.Balance)
			return bt.Success, nil
		},
	)
}

func handler2(text string, p *p.Player) bt.Node {
	return bt.New(
		bt.Selector,
		bt.New(
			bt.Sequence,
			bt.New(
				func(children []bt.Node) (bt.Status, error) {
					if text == "work" {
						return bt.Success, nil
					}
					return bt.Failure, nil
				},
			),
			work(p),
		),
		bt.New(
			bt.Sequence,
			bt.New(

				func(children []bt.Node) (bt.Status, error) {
					if text == "skip" {
						return bt.Success, nil
					}
					return bt.Failure, nil
				},
			),
			skip(p),
		),
		// bt.New(
		// bt.Sequence,
		// func(children []bt.Node) (bt.Status, error) {
		// 	if text == "grade" {
		// 		fmt.Println("Widtch grade do you want?")
		// 		bt.New(
		// 			bt.Selector,
		// 			bt.New(
		// 				func(children []bt.Node) (bt.Status, errors) {
		// 					if text == "exit" {
		// 						return bt.Success, nil
		// 					}
		// 					return bt.Failure, nil
		// 				},
		// 			),
		// 			bt.New(
		// 				func(children []bt.Node) (bt.Status, errors) {
		// 					if text == "home" {
		// 						fmt.Println("Player: ", p, " Want chanhe his: ", p.Buisnes)
		// 						p.Buisnes.ChangeGrade("home")
		// 						fmt.Println("Player: ", p, " Change business: ", p.Buisnes)
		// 					}
		// 				},
		// 			),
		// 		)
		// return bt.Failure, nil
		// }
		// },
		// ),
	)
}

func handler(text string) {
	sc := bufio.NewScanner(os.Stdin)
	if sc.Scan() {
		if text == "exit" || text == "" {

		}
	}
}

// ctx, cancel := context.WithTimeout(parent, 10 * time.Second)
// Поступило предложение сделать всё через behavior tree. Оно будет работать через контекст. Интересное предположение.
// Ещё заново перебрать контексты.

func main() {

	//Создаём бизнес в доступе.
	// parent := context.Background()
	b := b.New("None", 0, 0, "none")
	p := p.New(000)
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	fmt.Println("Write command.")

	ticker := bt.NewTickerStopOnFailure(
		context.Background(),
		time.Millisecond,
		handler2(GetText(), &p),
	)
	<-ticker.Done()

	fmt.Println(p.Balance, "\n", p.Buisnes)
}

// sc := reader
// for sc.Scan() {
// if...
//	for s2.Scan() {}
//}
//

// раз

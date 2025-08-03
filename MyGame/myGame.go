package main

import (
	"fmt"

	b "project/Lectures/MyGame/internal/buisnes"
	h "project/Lectures/MyGame/internal/handler"
	p "project/Lectures/MyGame/internal/player"

	"github.com/ilyakaznacheev/cleanenv"
)

//Предложена идея создать мапу с значениями. Какие проблемы могут быть? Проблемы с использованием функций. Надо попробовать функции в мапы.

// ctx, cancel := context.WithTimeout(parent, 10 * time.Second)
// Поступило предложение сделать всё через behavior tree. Оно будет работать через контекст. Интересное предположение.
// Ещё заново перебрать контексты.

// Создать соединение с базой данных. Сделать методы создания, редактирования, удаления строк пользователей.
// Создать временные файлы.
func main() {
	type config struct {
		User     string `env:"USER"`
		Password string `env:"PASSWORD"`
		DBNAME   string `env:"DBNAME"`
	}

	//Создаём бизнес в доступе.
	// parent := context.Background()
	var cfg2 config
	cleanenv.ReadConfig("/home/cpp/project/module/Env/.env", &cfg2)
	fmt.Println(cfg2)

	b := b.New("None", 0, 0, "none")
	p := p.New(000)
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	h.Handler(&p)

	fmt.Println(p.Balance, "\n", p.Buisnes)
}

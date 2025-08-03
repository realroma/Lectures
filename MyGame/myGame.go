package main

import (
	"fmt"

	b "project/Lectures/MyGame/internal/buisnes"
	daba "project/Lectures/MyGame/internal/database"
	h "project/Lectures/MyGame/internal/handler"
	p "project/Lectures/MyGame/internal/player"

	_ "github.com/lib/pq"
)

//Предложена идея создать мапу с значениями. Какие проблемы могут быть? Проблемы с использованием функций. Надо попробовать функции в мапы.

// ctx, cancel := context.WithTimeout(parent, 10 * time.Second)
// Поступило предложение сделать всё через behavior tree. Оно будет работать через контекст. Интересное предположение.
// Ещё заново перебрать контексты.

// Создать соединение с базой данных. Сделать методы создания, редактирования, удаления строк пользователей.
// Создать временные файлы.

func main() {
	//Создаём бизнес в доступе.
	// parent := context.Background()

	b := b.New("None", 0, 0, "none")
	p := p.New(000)
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	cfg := daba.Createconfig()
	db := daba.Opendb(cfg)

	daba.Exe(db, p.Id, *b)
	daba.Del(db, p.Id)
	h.Handler(p)

	fmt.Println(p.Balance, "\n", p.Buisnes)
}

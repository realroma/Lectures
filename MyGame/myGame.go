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

	var player *p.Player
	b := b.New("None", 0, 0, "none")
	player = p.New(1)
	player = nil
	id := 0

	//Открытие и закрытие базы данных.
	cfg := daba.Createconfig()
	db := daba.Opendb(cfg)
	defer db.Close()

	// Получение id и возвращение профиля игрока из базы данных.
	fmt.Println("Before login: ", player)
	if player == nil {
		player = h.Login(db, id)
	} else {
		if *player == *p.New(id) {
			fmt.Println("Login")
			player = h.Login(db, id)
		}
	}
	fmt.Println("After login: ", player)

	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, player)

	//Запись/удаление данных.
	// daba.Insert(db, &p.Id, *b, p.Balance)
	//daba.Delete(db, &p.Id)

	//Обработка событий.
	h.Handler(player, db)

	fmt.Println(player.Balance, "\n", player.Buisnes)
}

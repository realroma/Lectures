package main

import (
	"fmt"

	b "project/Lectures/MyGame/internal/buisnes"
<<<<<<< HEAD
	daba "project/Lectures/MyGame/internal/database"
	h "project/Lectures/MyGame/internal/handler"
	p "project/Lectures/MyGame/internal/player"

	_ "github.com/lib/pq"
=======
	h "project/Lectures/MyGame/internal/handler"
	p "project/Lectures/MyGame/internal/player"

	"github.com/ilyakaznacheev/cleanenv"
>>>>>>> 31e8e98fb86b970935d45d13dc767d9c2a0c7f2c
)

//Предложена идея создать мапу с значениями. Какие проблемы могут быть? Проблемы с использованием функций. Надо попробовать функции в мапы.

// ctx, cancel := context.WithTimeout(parent, 10 * time.Second)
// Поступило предложение сделать всё через behavior tree. Оно будет работать через контекст. Интересное предположение.
// Ещё заново перебрать контексты.
<<<<<<< HEAD

// Создать соединение с базой данных. Сделать методы создания, редактирования, удаления строк пользователей.
// Создать временные файлы.
=======
>>>>>>> 31e8e98fb86b970935d45d13dc767d9c2a0c7f2c

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
<<<<<<< HEAD

	var player *p.Player
	b := b.New("None", 0, 0, "none")
	player = p.New(1)
	player = nil
	id := 0

	//Открытие и закрытие базы данных.
	cfg := daba.Createconfig()
	db := daba.Opendb(cfg)
	defer db.Close()
=======
	var cfg2 config
	cleanenv.ReadConfig("/home/cpp/project/module/Env/.env", &cfg2)
	fmt.Println(cfg2)

	b := b.New("None", 0, 0, "none")
	p := p.New(000)
	fmt.Printf("Busines default: %v \nPlayer: %v \n", b, p)

	h.Handler(&p)
>>>>>>> 31e8e98fb86b970935d45d13dc767d9c2a0c7f2c

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

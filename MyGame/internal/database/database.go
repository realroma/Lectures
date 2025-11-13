package database

import (
	"database/sql"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DBName   string `env:"DBNAME"`
}

func Createconfig() config {
	var cfg config
	cleanenv.ReadConfig("/home/cpp/project/Lectures/MyGame/env/.env", &cfg)
	fmt.Println(cfg)

	return cfg
}

func Opendb(cfg config) (db *sql.DB) {
	var dbstr = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cfg.User, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", dbstr)
	if err != nil {
		panic(err)
	}

	//Закыть таблицу тут, через defer, мы не можем. Надо это делать на уровне выше, откуда идёт вызов таблицы.
	return db
}

// Передаём строку в базу данных.
func Insert(db *sql.DB, id *int, balance int) {
	fmt.Println("Insert into Database")
	request, err := db.Exec("INSERT INTO users (id, balance) VALUES ($1, $2, $3)", id, balance)
	if err != nil {
		fmt.Println("Error in package database:")
		panic(err)
	}
	fmt.Println(request)
}

// Удаляем из базы данных по id
func Delete(db *sql.DB, id *int) {
	fmt.Println("Delete from database")
	result, err := db.Exec("DELETE FROM users WHERE id = $1", &id)
	if err != nil {
		fmt.Println("Error in package database:")
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}

// Находим пользователя по id
func Select(db *sql.DB, id int) sql.Result {
	fmt.Println("Found user.")
	request, err := db.Exec("SELECT FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println("Error in package database:")
		panic(err)
	}
	fmt.Println("DB: ", request)
	return request
}

// Изменяем данные пользователя
func Alter(db *sql.DB) {
	fmt.Println("Alter in database user: ")
	request, err := db.Exec("ALTER TABLE users ")
	if err != nil {
		fmt.Println("Error in package database:")
		panic(err)
	}
	fmt.Println("DB: ", request)
}

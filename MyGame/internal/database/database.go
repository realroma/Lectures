package database

import (
	"database/sql"
	"fmt"
	"log"

	b "project/Lectures/MyGame/internal/buisnes"

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
	defer db.Close()
	return db
}

func Exe(db *sql.DB, id int, b b.Buisnes) {
	fmt.Println("Insert into Database")
	request, err := db.Exec("INSERT INTO users (id, busineses) VALUES ($1, $1)", id, b)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(request)
}

func Del(db *sql.DB, id int) {
	fmt.Println("Delete from database")
	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.RowsAffected())
}

func Found(db *sql.DB, id int) {
	fmt.Println("Found user.")
	request, err := db.Exec("SELECT FROM users WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(request)
}

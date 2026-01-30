package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Prinln("Error loading .env file. Use default config.")
	}
	//Чтобы не было возникновения неизвестных параметров вместо простого os.Getenv("DSN"), мы явно вернём структуру.
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}

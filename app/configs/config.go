package configs

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	Host    string `env:"APP_HOST"`
	Port    int    `env:"APP_PORT"`
	LogPath string `env:"APP_LOG_PATH"`
}

type DatabaseConfig struct {
	Host     string `env:"DATABASE_HOST"`
	Port     int    `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Password string `env:"DATABASE_PASSWORD"`
	Name     string `env:"DATABASE_NAME"`
}

type config struct {
	App AppConfig
	DB  DatabaseConfig
}

var (
	App AppConfig
	DB  DatabaseConfig
)

func InitConfig(path string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file: " + err.Error())
	}

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Error parse env data to struct: " + err.Error())
	}

	App = cfg.App
	DB = cfg.DB
}

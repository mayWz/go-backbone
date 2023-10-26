package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ENV struct {
	ENV              string
	DatabaseHost     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabasePort     string
	SecertKey        string
}

func SetupENV() ENV {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: can't load .env file")
	}

	// Setup the loaded enviroment
	env := os.Getenv("ENV")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	secertKey := os.Getenv("SECRET_KEY")

	return ENV{
		ENV:              env,
		DatabaseHost:     dbHost,
		DatabaseUsername: dbUsername,
		DatabasePassword: dbPassword,
		DatabaseName:     dbName,
		DatabasePort:     dbPort,
		SecertKey:        secertKey,
	}
}

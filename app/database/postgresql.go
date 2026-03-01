package database

import (
	"log"
	"os"

	pg "github.com/go-pg/pg/v10"
	// PostgreSQL driver
)

var DB *pg.DB

func ConnectPostgreSql() error {
	// 	export DB_HOST=localhost
	// export DB_PORT=5400
	// export DB_USER=admin
	// export DB_PASSWORD=admin@123
	// export DB_NAME=whatsappcrm
	// export APP_ENV=docker

	// go run main.go
	appEnv := "docker"
	// appEnv := os.Getenv("APP_ENV")
	var options *pg.Options
	if appEnv == "local" {
		options = &pg.Options{
			Network:  "tcp",
			Addr:     "localhost" + ":" + "5432",
			User:     "admin",
			Password: "admin@123",
			Database: "whatsappcrm",
		}
	}
	if appEnv == "docker" {
		options = &pg.Options{
			Network:  "tcp",
			Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		}
	}
	log.Printf("Connecting with options: Addr=%s, User=%s, Database=%s\n", options.Addr, options.User, options.Database)

	db := pg.Connect(options)
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatal("failed to connect with pgsql with error: ", err)

		return err
	}
	log.Println("Database is connected")
	DB = db
	return nil
}

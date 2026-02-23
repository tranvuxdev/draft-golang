package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tranvux/draft-go/cmd/router"
	"github.com/tranvux/draft-go/internal/model"
	"github.com/tranvux/draft-go/internal/seed"
	"github.com/tranvux/draft-go/pkg/database"
)

func main() {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("main: cannot load .env file", err)
	}

	// get db
	db := database.Connect()

	// connect db
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("main: cannot get env file", err)
	}

	fmt.Println("main: connect successfully")

	// ping db
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("main: database cannot ping", err)
	}

	// check db
	db.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto")

	// auto migrate
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
		&model.Tag{})

	// seed data
	shouldSeed := flag.Bool("seed", false, "run seed data")
	flag.Parse()
	if *shouldSeed {
		seed.Run(db)
	}

	r := router.Setup(db)
	r.Run(os.Getenv("SERVER_PORT"))
}

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tranvux/learn-structs/cmd/router"
	"github.com/tranvux/learn-structs/internal/model"
	"github.com/tranvux/learn-structs/internal/seed"
	"github.com/tranvux/learn-structs/pkg/database"
)

func main() {
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
	r.Run(":8080")
}

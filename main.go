package main

import (
	"log"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}

	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var groups []entity.Group

	db.Debug().Where("id=?", "g-xya").Preload("Address").Preload("Properties").Find(&groups)
	log.Printf("%+v", groups)
}

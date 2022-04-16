package admin

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./../../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file: %s\n", err.Error())
	}
}

func TestFindByUsername(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo AdminRepository = NewAdminRepositoryImpl(db)
	if user, err := repo.FindByUsername(context.Background(), "admin"); err != nil {
		t.Log(err.Error())
	} else {
		t.Logf("%+v", user)
	}
}

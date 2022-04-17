package property

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./../../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file: %s\n", err.Error())
	}
}

func TestInsert(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	property := entity.Property{
		ID:          "p-xya2222",
		Name:        "Topeng Bujang Ganong",
		Description: "Topeng bujang ganong",
		Amount:      3,
		GroupID:     "g-xya",
	}

	var repo PropertyRepository = NewPropertyRepositoryImpl(db)

	if err := repo.Insert(context.Background(), property); err != nil {
		t.Log(err)
	}
}

func TestUpdate(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	property := entity.Property{
		ID:          "p-xya2222",
		Name:        "Topeng Bujang Ganong Update",
		Description: "Ini adalah deskripsi topeng bujang ganong",
		Amount:      5,
	}

	var repo PropertyRepository = NewPropertyRepositoryImpl(db)

	if err := repo.Update(context.Background(), property); err != nil {
		t.Log(err)
	}
}

func TestDelete(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo PropertyRepository = NewPropertyRepositoryImpl(db)

	if err := repo.Delete(context.Background(), "p-xya2222"); err != nil {
		t.Log(err)
	}
}

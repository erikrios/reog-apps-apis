package village

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./../../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file: %s\n", err.Error())
	}
}

func TestFindByID(t *testing.T) {
	var repo VillageRepository = NewVillageRepositoryImpl()
	if village, err := repo.FindByID("3502030007"); err != nil {
		t.Log(err)
	} else {
		t.Logf("%+v", village)
	}
}

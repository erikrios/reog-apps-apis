package address

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

func TestUpdate(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo AddressRepository = NewAddressRepositoryImpl(db)

	address := entity.Address{
		ID:           "g-xya",
		Address:      "Address 1 Update",
		VillageID:    "5321101010",
		VillageName:  "Pager",
		DistrictID:   "5321101",
		DistrictName: "Bungkal",
		RegencyID:    "5321",
		RegencyName:  "Ponorogo",
		ProvinceID:   "53",
		ProvinceName: "Jawa Timur",
	}

	if err := repo.Update(context.Background(), address); err != nil {
		t.Log(err)
	}
}

package address

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/address"
	"github.com/erikrios/reog-apps-apis/repository/village"
	"github.com/joho/godotenv"
)

var addressService AddressService

func init() {
	if err := godotenv.Load("./../../.env.local"); err != nil {
		log.Fatalf("Error loading .env.local file: %s\n", err.Error())
	}

	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)
	addressRepo := address.NewAddressRepositoryImpl(db)
	villageRepo := village.NewVillageRepositoryImpl()
	addressService = NewAddressServiceImpl(addressRepo, villageRepo)
}

func TestUpdate(t *testing.T) {
	payload := payload.UpdateAddress{
		Address:   "RT 02 RW 03 Dukuh Tengah",
		VillageID: "3502020019",
	}

	if err := addressService.Update(context.Background(), "g-eaI", payload); err != nil {
		t.Log("error:", err)
	}
}

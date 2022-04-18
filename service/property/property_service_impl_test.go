package property

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/property"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"github.com/joho/godotenv"
)

var propertyService PropertyService

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
	propertyRepo := property.NewPropertyRepositoryImpl(db)
	groupRepo := group.NewGroupRepositoryImpl(db)
	idGenerator := generator.NewNanoidIDGenerator()
	qrCodeGenerator := generator.NewQRCodeGeneratorImpl()
	propertyService = NewPropertyServiceImpl(propertyRepo, groupRepo, idGenerator, qrCodeGenerator)
}

func TestCreate(t *testing.T) {
	payload := payload.CreateProperty{
		Name:        "Dadak Merak",
		Description: "Ini adalah dadak merak",
		Amount:      1,
	}

	if id, err := propertyService.Create(context.Background(), "g-eaI", payload); err != nil {
		t.Log("error:", err)
	} else {
		t.Log("no error:", id)
	}
}

func TestUpdate(t *testing.T) {
	payload := payload.UpdateProperty{
		Name:        "Bujang Ganong Update",
		Description: "Ini adalah bujang ganong update",
		Amount:      10,
	}

	if err := propertyService.Update(context.Background(), "p-FJPDo81", payload); err != nil {
		t.Log("error:", err)
	}
}

func TestDelete(t *testing.T) {
	if err := propertyService.Delete(context.Background(), "p-FJPDo81"); err != nil {
		t.Log("error: ", err)
	}
}

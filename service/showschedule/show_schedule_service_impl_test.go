package showschedule

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/showschedule"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load("./../../.env.local"); err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
	}
	var err error
	db, err = config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}
}

func TestCreate(t *testing.T) {
	var service ShowScheduleService = NewShowScheduleServiceImpl(
		showschedule.NewShowScheduleRepositoryImpl(db),
		group.NewGroupRepositoryImpl(db),
		generator.NewNanoidIDGenerator(),
	)

	p := payload.CreateShowSchedule{
		GroupID:  "g-Nzo",
		Place:    "Lapangan Bungkal",
		StartOn:  "01 May 22 14:00 WIB",
		FinishOn: "01 May 22 17:00 WIB",
	}

	if id, err := service.Create(context.Background(), p); err != nil {
		t.Log("error:", err)
	} else {
		t.Logf("no error: %s", id)
	}
}

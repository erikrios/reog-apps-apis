package group

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/group"
	"github.com/erikrios/reog-apps-apis/repository/village"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"github.com/joho/godotenv"
)

var groupService GroupService

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
	groupRepo := group.NewGroupRepositoryImpl(db)
	villageRepo := village.NewVillageRepositoryImpl()
	idGenerator := generator.NewNanoidIDGenerator()
	groupService = NewGroupServiceImpl(groupRepo, villageRepo, idGenerator)
}

func TestCreate(t *testing.T) {
	payload := payload.CreateGroup{
		Name:      "Group Satu",
		Leader:    "Erik R",
		Address:   "RT 01 RW 01 Dukuh Bibis",
		VillageID: "3502030007",
	}

	if id, err := groupService.Create(context.Background(), payload); err != nil {
		t.Log("Result err: ", err)
	} else {
		t.Log("Result id: ", id)
	}
}

func TestGetAll(t *testing.T) {
	if responses, err := groupService.GetAll(context.Background()); err != nil {
		t.Log("Result err: ", err)
	} else {
		t.Logf("Result responses: %+v", responses)
	}
}

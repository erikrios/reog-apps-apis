package showschedule

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./../../.env.local"); err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
	}
}

func TestInsert(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repository ShowcheduleRepository = NewShowScheduleRepositoryImpl(db)

	showSchedule := entity.ShowSchedule{
		ID:       "s-An9LEb",
		GroupID:  "g-Nzo",
		Place:    "Lapangan Bungkal",
		StartOn:  time.Now().Add(5 * time.Hour),
		FinishOn: time.Now().Add(8 * time.Hour),
	}

	if err := repository.Insert(context.Background(), showSchedule); err != nil {
		t.Log("error:", err)
	} else {
		t.Log("no error:", showSchedule.ID)
	}
}

func TestFindAll(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repository ShowcheduleRepository = NewShowScheduleRepositoryImpl(db)

	if showSchedules, err := repository.FindAll(context.Background()); err != nil {
		t.Log("error:", err)
	} else {
		t.Logf("no error: %+v", showSchedules)
	}
}

func TestFindByGroupID(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repository ShowcheduleRepository = NewShowScheduleRepositoryImpl(db)

	if showSchedules, err := repository.FindByGroupID(context.Background(), "g-Nzo"); err != nil {
		t.Log("error:", err)
	} else {
		t.Logf("no error: %+v", showSchedules)
	}
}

func TestUpdate(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repository ShowcheduleRepository = NewShowScheduleRepositoryImpl(db)

	showSchedule := entity.ShowSchedule{
		Place:    "Lapangan Bungkal Update",
		StartOn:  time.Now().Add(6 * time.Hour),
		FinishOn: time.Now().Add(9 * time.Hour),
	}

	if err := repository.Update(context.Background(), "s-An9LEb", showSchedule); err != nil {
		t.Log("error:", err)
	} else {
		t.Log("no error:", showSchedule.Place)
	}
}

func TestDelete(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repository ShowcheduleRepository = NewShowScheduleRepositoryImpl(db)

	if err := repository.Delete(context.Background(), "s-An9LEb"); err != nil {
		t.Log("error:", err)
	} else {
		t.Log("no error:", "record deleted")
	}
}

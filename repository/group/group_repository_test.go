package group

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

	var repo GroupRepository = NewGroupRepositoryImpl(db)
	group := entity.Group{
		ID:     "g-xya",
		Name:   "Group Satu",
		Leader: "Erik R",
		Address: entity.Address{
			Address:      "Address 3",
			VillageID:    "5321101010",
			VillageName:  "Pager",
			DistrictID:   "5321101",
			DistrictName: "Bungkal",
			RegencyID:    "5321",
			RegencyName:  "Ponorogo",
			ProvinceID:   "53",
			ProvinceName: "Jawa Timur",
		},
	}

	if err := repo.Insert(context.Background(), group); err != nil {
		t.Log(err.Error())
	}
}

func TestInsertAll(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo GroupRepository = NewGroupRepositoryImpl(db)
	groups := []entity.Group{
		{
			ID:     "g-xyb",
			Name:   "Group Dua",
			Leader: "Erik R",
			Address: entity.Address{
				Address:      "Address 2",
				VillageID:    "5321101010",
				VillageName:  "Pager",
				DistrictID:   "5321101",
				DistrictName: "Bungkal",
				RegencyID:    "5321",
				RegencyName:  "Ponorogo",
				ProvinceID:   "53",
				ProvinceName: "Jawa Timur",
			},
		},
		{
			ID:     "g-xya",
			Name:   "Group Satu",
			Leader: "Erik S",
			Address: entity.Address{
				Address:      "Address Satu",
				VillageID:    "5321101010",
				VillageName:  "Pager",
				DistrictID:   "5321101",
				DistrictName: "Bungkal",
				RegencyID:    "5321",
				RegencyName:  "Ponorogo",
				ProvinceID:   "53",
				ProvinceName: "Jawa Timur",
			},
		},
		{
			ID:     "g-xyc",
			Name:   "Group Tiga",
			Leader: "Erik R",
			Address: entity.Address{
				Address:      "Address 3",
				VillageID:    "5321101010",
				VillageName:  "Pager",
				DistrictID:   "5321101",
				DistrictName: "Bungkal",
				RegencyID:    "5321",
				RegencyName:  "Ponorogo",
				ProvinceID:   "53",
				ProvinceName: "Jawa Timur",
			},
		},
	}

	if err := repo.InsertAll(context.Background(), groups); err != nil {
		t.Log(err)
	}
}

func TestFindAll(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo GroupRepository = NewGroupRepositoryImpl(db)

	if groups, err := repo.FindAll(context.Background()); err != nil {
		t.Log(err)
	} else {
		t.Logf("%+v", groups)
	}
}

func TestFindByID(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo GroupRepository = NewGroupRepositoryImpl(db)

	if group, err := repo.FindByID(context.Background(), "g-xya"); err != nil {
		t.Log(err)
	} else {
		t.Logf("%+v", group)
	}
}

func TestUpdate(t *testing.T) {
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo GroupRepository = NewGroupRepositoryImpl(db)

	newGroup := entity.Group{
		ID:     "g-xyb",
		Name:   "Group Dua Update",
		Leader: "Erik R",
		Address: entity.Address{
			Address:      "Address 2 Update",
			VillageID:    "5321101010",
			VillageName:  "Pager",
			DistrictID:   "5321101",
			DistrictName: "Bungkal",
			RegencyID:    "5321",
			RegencyName:  "Ponorogo",
			ProvinceID:   "53",
			ProvinceName: "Jawa Timur",
		},
	}

	if err := repo.Update(context.Background(), newGroup); err != nil {
		t.Log(err)
	}
}

func TestDelete(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	var repo GroupRepository = NewGroupRepositoryImpl(db)

	if err := repo.Delete(context.Background(), "g-xyb"); err != nil {
		t.Log(err)
	}
}

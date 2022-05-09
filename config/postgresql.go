package config

import (
	"fmt"
	"os"

	"github.com/erikrios/reog-apps-apis/entity"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLDatabase() (*gorm.DB, error) {
	var sslMode string
	if os.Getenv("ENV") == "production" {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), sslMode,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func MigratePostgreSQLDatabase(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Admin{}, &entity.Group{}, &entity.Address{}, &entity.Property{}, &entity.ShowSchedule{})
}

func SetInitialDataPostgreSQLDatabase(db *gorm.DB) error {
	idGenerator := generator.NewNanoidIDGenerator()
	passwordGenerator := generator.NewBcryptPasswordGenerator()

	id, err := idGenerator.GenerateAdminID()
	if err != nil {
		return err
	}

	password, err := passwordGenerator.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), 10)
	if err != nil {
		return err
	}

	admin := &entity.Admin{
		ID:       id,
		Username: os.Getenv("ADMIN_USERNAME"),
		Name:     os.Getenv("ADMIN_NAME"),
		Password: string(password),
	}

	result := db.Save(admin)
	return result.Error
}

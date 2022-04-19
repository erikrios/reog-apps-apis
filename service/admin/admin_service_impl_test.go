package admin

import (
	"context"
	"log"
	"testing"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/model/payload"
	"github.com/erikrios/reog-apps-apis/repository/admin"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var adminService AdminService

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
	adminRepo := admin.NewAdminRepositoryImpl(db)
	bcryptPasswordGenerator := generator.NewBcryptPasswordGenerator()
	jwtTokenGenerator := generator.NewJWTTokenGenerator(echo.New().NewContext(nil, nil))
	adminService = NewAdminServiceImpl(adminRepo, bcryptPasswordGenerator, jwtTokenGenerator)
}

func TestLogin(t *testing.T) {
	payload := payload.Credential{
		Username: "admin",
		Password: "erikrios",
	}

	if token, err := adminService.Login(context.Background(), payload); err != nil {
		t.Log("error:", err)
	} else {
		t.Log("no error:", token)
	}
}

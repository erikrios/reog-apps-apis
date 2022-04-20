package main

import (
	"log"
	"os"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/controller"
	_ "github.com/erikrios/reog-apps-apis/docs"
	"github.com/erikrios/reog-apps-apis/middleware"
	ar "github.com/erikrios/reog-apps-apis/repository/admin"
	gr "github.com/erikrios/reog-apps-apis/repository/group"
	vr "github.com/erikrios/reog-apps-apis/repository/village"
	as "github.com/erikrios/reog-apps-apis/service/admin"
	gs "github.com/erikrios/reog-apps-apis/service/group"
	"github.com/erikrios/reog-apps-apis/utils/generator"
	_ "github.com/erikrios/reog-apps-apis/validation"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title           Reog Apps API
// @version         1.0
// @description     API for Reog Group in Ponorogo
// @termsOfService  http://swagger.io/terms/

// @contact.name   Erik Rio Setiawan
// @contact.url    http://www.swagger.io/support
// @contact.email  erikriosetiawan15@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost
// @BasePath  /api/v1
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s\n", err.Error())
	}

	db, err := config.NewPostgreSQLDatabase()
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Successfully connected to database with instance address: %p\n", db)
	}

	config.MigratePostgreSQLDatabase(db)
	config.SetInitialDataPostgreSQLDatabase(db)

	port := ":" + os.Getenv("PORT")

	passwordGenerator := generator.NewBcryptPasswordGenerator()
	tokenGenerator := generator.NewJWTTokenGenerator()
	idGenerator := generator.NewNanoidIDGenerator()
	qrCodeGenerator := generator.NewQRCodeGeneratorImpl()

	adminRepository := ar.NewAdminRepositoryImpl(db)
	groupRepository := gr.NewGroupRepositoryImpl(db)
	villageRepository := vr.NewVillageRepositoryImpl()

	adminService := as.NewAdminServiceImpl(adminRepository, passwordGenerator, tokenGenerator)
	groupService := gs.NewGroupServiceImpl(groupRepository, villageRepository, idGenerator, qrCodeGenerator)

	adminsController := controller.NewAdminsController(adminService)
	groupsController := controller.NewGroupsController(groupService)

	e := echo.New()

	if os.Getenv("ENV") == "production" {
		middleware.BodyLimit(e)
		middleware.Gzip(e)
		middleware.RateLimiter(e)
		middleware.Recover(e)
		middleware.Secure(e)
		middleware.RemoveTrailingSlash(e)
	} else {
		middleware.Logger(e)
		middleware.RemoveTrailingSlash(e)
	}

	e.GET("/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1")
	adminsController.Route(g)
	groupsController.Route(g)

	e.Logger.Fatal(e.Start(port))
}

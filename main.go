package main

import (
	"log"
	"os"

	"github.com/erikrios/reog-apps-apis/config"
	"github.com/erikrios/reog-apps-apis/controller"
	_ "github.com/erikrios/reog-apps-apis/docs"
	"github.com/erikrios/reog-apps-apis/middleware"
	dr "github.com/erikrios/reog-apps-apis/repository/address"
	ar "github.com/erikrios/reog-apps-apis/repository/admin"
	gr "github.com/erikrios/reog-apps-apis/repository/group"
	pr "github.com/erikrios/reog-apps-apis/repository/property"
	ssr "github.com/erikrios/reog-apps-apis/repository/showschedule"
	vr "github.com/erikrios/reog-apps-apis/repository/village"
	ds "github.com/erikrios/reog-apps-apis/service/address"
	as "github.com/erikrios/reog-apps-apis/service/admin"
	gs "github.com/erikrios/reog-apps-apis/service/group"
	ps "github.com/erikrios/reog-apps-apis/service/property"
	sss "github.com/erikrios/reog-apps-apis/service/showschedule"
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

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @host      localhost:3000
// @BasePath  /api/v1
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %s\n", err.Error())
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
	addressRepository := dr.NewAddressRepositoryImpl(db)
	propertyRepository := pr.NewPropertyRepositoryImpl(db)
	showScheduleRepository := ssr.NewShowScheduleRepositoryImpl(db)

	adminService := as.NewAdminServiceImpl(adminRepository, passwordGenerator, tokenGenerator)
	groupService := gs.NewGroupServiceImpl(groupRepository, villageRepository, idGenerator, qrCodeGenerator)
	addressService := ds.NewAddressServiceImpl(addressRepository, villageRepository)
	propertyService := ps.NewPropertyServiceImpl(propertyRepository, groupRepository, idGenerator, qrCodeGenerator)
	showScheduleService := sss.NewShowScheduleServiceImpl(showScheduleRepository, groupRepository, idGenerator)

	adminsController := controller.NewAdminsController(adminService)
	groupsController := controller.NewGroupsController(groupService, propertyService, addressService)
	showSchedulesController := controller.NewShowSchedulesController(showScheduleService)

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
	showSchedulesController.Route(g)
	e.Logger.Fatal(e.Start(port))
}

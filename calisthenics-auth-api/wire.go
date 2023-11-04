//go:build wireinject
// +build wireinject

package main

import (
	"calisthenics-auth-api/api"
	"calisthenics-auth-api/config"
	"calisthenics-auth-api/data"
	"calisthenics-auth-api/data/repository"
	"calisthenics-auth-api/middleware"
	"calisthenics-auth-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GeneralSet = wire.NewSet(NewDatabase, InitRoutes)

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewProfileRepository,
)

var DomainServiceSet = wire.NewSet(
	service.NewUserService,
	service.NewProfileService,
)

var ServiceSet = wire.NewSet(
	service.NewAuthService,
	service.NewUserProfileService,
)

var ControllerSet = wire.NewSet(
	api.NewAuthController,
	api.NewUserProfileController,
	api.NewUserController,
)

func InitializeApp() *echo.Echo {
	wire.Build(GeneralSet, RepositorySet, DomainServiceSet, ServiceSet, ControllerSet)
	return &echo.Echo{}
}

func NewDatabase() *gorm.DB {
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)))
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = db.AutoMigrate(&data.User{}, data.Profile{})
	if err != nil {
		panic("Failed to migrate database")
	}
	return db
}

func InitRoutes(
	authController *api.AuthController,
	userProfileController *api.UserProfileController,
	userController *api.UserController) *echo.Echo {

	e := echo.New()
	e.Use(middleware.ServiceContextMiddleware)
	e.Validator = &config.CustomValidator{validator.New()}
	authController.InitAuthRoutes(e)

	authMiddleware := middleware.NewAuthenticationMiddleware()
	authMiddlewareGroup := e.Group("/v1/user", authMiddleware.AuthenticationMiddleware)
	userProfileController.InitUserProfileRoutes(authMiddlewareGroup)
	userController.InitUserRoutes(authMiddlewareGroup)

	return e
}

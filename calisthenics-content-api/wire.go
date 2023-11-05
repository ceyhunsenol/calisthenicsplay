//go:build wireinject
// +build wireinject

package main

import (
	"calisthenics-content-api/api"
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
	"calisthenics-content-api/data/repository"
	"calisthenics-content-api/middleware"
	"calisthenics-content-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GeneralSet = wire.NewSet(NewDatabase, InitRoutes, cache.NewCacheManager)

var RepositorySet = wire.NewSet(
	repository.NewContentRepository,
	repository.NewMediaRepository,
	repository.NewGenreRepository,
)

var DomainServiceSet = wire.NewSet(
	service.NewContentService,
	service.NewMediaService,
	service.NewGenreService,
)

var CacheServiceSet = wire.NewSet(
	//cache.NewMediaCacheService,
	cache.NewContentCacheService,
	//cache.NewGenreCacheService,
)

var ControllerSet = wire.NewSet(
	api.NewCacheController,
)

func InitializeApp() *echo.Echo {
	wire.Build(GeneralSet, CacheServiceSet, ControllerSet)
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
	err = db.AutoMigrate()
	if err != nil {
		panic("Failed to migrate database")
	}
	return db
}

func InitRoutes(cacheController *api.CacheController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.ServiceContextMiddleware)
	cacheController.InitCacheRoutes(e)
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	return e
}

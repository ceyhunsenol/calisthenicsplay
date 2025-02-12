// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"calisthenics-content-api/api"
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
	"calisthenics-content-api/data/repository"
	"calisthenics-content-api/integration/calisthenics"
	"calisthenics-content-api/job"
	middleware2 "calisthenics-content-api/middleware"
	"calisthenics-content-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeApp() *echo.Echo {
	db := NewDatabase()
	iMediaRepository := repository.NewMediaRepository(db)
	iMediaService := service.NewMediaService(iMediaRepository)
	iCacheService := cache.NewCacheManager()
	iMediaCacheService := cache.NewMediaCacheService(iCacheService)
	iContentRepository := repository.NewContentRepository(db)
	iContentService := service.NewContentService(iContentRepository)
	iMediaCacheOperations := service.NewMediaCacheOperations(iMediaService, iMediaCacheService, iContentService)
	iContentCacheService := cache.NewContentCacheService(iCacheService)
	iContentCacheOperations := service.NewContentCacheOperations(iContentService, iContentCacheService, iMediaService, iMediaCacheService)
	iGenreRepository := repository.NewGenreRepository(db)
	iGenreService := service.NewGenreService(iGenreRepository)
	iGenreCacheService := cache.NewGenreCacheService(iCacheService)
	iGenreCacheOperations := service.NewGenreCacheOperations(iGenreService, iGenreCacheService)
	iGeneralInfoCacheService := cache.NewGeneralInfoCacheService(iCacheService)
	iGeneralInfoRepository := repository.NewGeneralInfoRepository(db)
	iGeneralInfoService := service.NewGeneralInfoService(iGeneralInfoRepository)
	iGeneralInfoCacheOperations := service.NewGeneralInfoCacheOperations(iGeneralInfoCacheService, iGeneralInfoService)
	iContentAccessCacheService := cache.NewContentAccessCacheService(iCacheService)
	iContentAccessRepository := repository.NewContentAccessRepository(db)
	iContentAccessService := service.NewContentAccessService(iContentAccessRepository)
	iContentAccessCacheOperations := service.NewContentAccessCacheOperations(iContentAccessCacheService, iContentAccessService)
	iMediaAccessCacheService := cache.NewMediaAccessCacheService(iCacheService)
	iMediaAccessRepository := repository.NewMediaAccessRepository(db)
	iMediaAccessService := service.NewMediaAccessService(iMediaAccessRepository)
	iMediaAccessCacheOperations := service.NewMediaAccessCacheOperations(iMediaAccessCacheService, iMediaAccessService)
	ihlsEncodingCacheService := cache.NewHLSEncodingCacheService(iCacheService)
	iEncodingRepository := repository.NewEncodingRepository(db)
	iEncodingService := service.NewEncodingService(iEncodingRepository)
	ihlsEncodingCacheOperations := service.NewHLSEncodingCacheOperations(ihlsEncodingCacheService, iEncodingService)
	iTranslationCacheService := cache.NewTranslationCacheService(iCacheService)
	iTranslationRepository := repository.NewTranslationRepository(db)
	iTranslationService := service.NewTranslationService(iTranslationRepository)
	iTranslationCacheOperations := service.NewTranslationCacheOperations(iTranslationCacheService, iTranslationService)
	iLimitedCacheService := cache.NewLimitedCacheService(iCacheService)
	iJobService := job.NewJobService(iLimitedCacheService)
	iInitService := service.NewInitService(iContentCacheOperations, iGenreCacheOperations, iMediaCacheOperations, iContentAccessCacheOperations, iGeneralInfoCacheOperations, iMediaAccessCacheOperations, ihlsEncodingCacheOperations, iTranslationCacheOperations, iJobService)
	cacheController := api.NewCacheController(iMediaCacheOperations, iContentCacheOperations, iGenreCacheOperations, iGeneralInfoCacheOperations, iContentAccessCacheOperations, iMediaAccessCacheOperations, ihlsEncodingCacheOperations, iInitService)
	iGenreOperations := service.NewGenreOperations(iGenreService, iGenreCacheService)
	genreController := api.NewGenreController(iGenreOperations)
	iContentOperations := service.NewContentOperations(iContentService, iContentCacheService)
	contentController := api.NewContentController(iContentOperations)
	iCalisthenicsAuthService := calisthenics.NewCalisthenicsAuthService(iLimitedCacheService)
	iMediaPlayActionService := service.NewMediaPlayActionService(iGeneralInfoCacheService, iContentAccessCacheService, iMediaAccessCacheService, iCalisthenicsAuthService, iMediaCacheService, iTranslationCacheService)
	iParameterService := service.NewParameterService(iGeneralInfoCacheService)
	ihlsMediaService := service.NewHLSMediaService(ihlsEncodingCacheService, iMediaCacheService, iMediaPlayActionService, iParameterService)
	hlsMediaController := api.NewHLSMediaController(ihlsMediaService)
	hlsMediaTestController := api.NewHLSMediaTestController()
	echoEcho := InitRoutes(cacheController, genreController, contentController, hlsMediaController, hlsMediaTestController, iInitService)
	return echoEcho
}

// wire.go:

var GeneralSet = wire.NewSet(NewDatabase, InitRoutes, cache.NewCacheManager, job.NewJobService)

var IntegrationSet = wire.NewSet(calisthenics.NewCalisthenicsAuthService)

var RepositorySet = wire.NewSet(repository.NewContentRepository, repository.NewMediaRepository, repository.NewGenreRepository, repository.NewGeneralInfoRepository, repository.NewContentAccessRepository, repository.NewMediaAccessRepository, repository.NewEncodingRepository, repository.NewTranslationRepository)

var DomainServiceSet = wire.NewSet(service.NewContentService, service.NewMediaService, service.NewGenreService, service.NewInitService, service.NewGeneralInfoService, service.NewContentAccessService, service.NewMediaAccessService, service.NewEncodingService, service.NewTranslationService)

var CacheServiceSet = wire.NewSet(cache.NewMediaCacheService, cache.NewContentCacheService, cache.NewGenreCacheService, cache.NewGeneralInfoCacheService, cache.NewMediaAccessCacheService, cache.NewContentAccessCacheService, cache.NewLimitedCacheService, cache.NewHLSEncodingCacheService, cache.NewTranslationCacheService)

var ServiceSet = wire.NewSet(service.NewMediaOperations, service.NewContentOperations, service.NewGenreOperations, service.NewMediaCacheOperations, service.NewContentCacheOperations, service.NewGenreCacheOperations, service.NewGeneralInfoCacheOperations, service.NewContentAccessCacheOperations, service.NewMediaAccessCacheOperations, service.NewMediaPlayActionService, service.NewHLSEncodingCacheOperations, service.NewHLSMediaService, service.NewParameterService, service.NewTranslationCacheOperations)

var ControllerSet = wire.NewSet(api.NewCacheController, api.NewGenreController, api.NewContentController, api.NewHLSMediaTestController, api.NewHLSMediaController)

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

func InitRoutes(
	cacheController *api.CacheController,
	genreController *api.GenreController,
	contentController *api.ContentController,
	hlsMediaController *api.HLSMediaController,
	hlsMediaTestController *api.HLSMediaTestController,
	initService service.IInitService,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Static("/downloads/"))
	e.GET("/downloads/:file", func(c echo.Context) error {
		fileName := c.Param("file")
		return c.Attachment("downloads/"+fileName, fileName)
	})
	e.Use(middleware2.ServiceContextMiddleware)
	cacheController.InitCacheRoutes(e)
	genreController.InitGenreRoutes(e)
	hlsMediaController.InitHLSMediaRoutes(e)
	hlsMediaTestController.InitHLSMediaTestRoutes(e)
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	initService.InitCache()
	go initService.InitJob()
	return e
}

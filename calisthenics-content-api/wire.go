//go:build wireinject
// +build wireinject

package main

import (
	"calisthenics-content-api/api"
	"calisthenics-content-api/cache"
	"calisthenics-content-api/config"
	"calisthenics-content-api/data/repository"
	"calisthenics-content-api/integration/calisthenics"
	"calisthenics-content-api/job"
	"calisthenics-content-api/middleware"
	"calisthenics-content-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GeneralSet = wire.NewSet(NewDatabase, InitRoutes, cache.NewCacheManager, job.NewJobService)

var IntegrationSet = wire.NewSet(
	calisthenics.NewCalisthenicsAuthService,
)

var RepositorySet = wire.NewSet(
	repository.NewContentRepository,
	repository.NewMediaRepository,
	repository.NewGenreRepository,
	repository.NewGeneralInfoRepository,
	repository.NewContentAccessRepository,
	repository.NewMediaAccessRepository,
	repository.NewEncodingRepository,
	repository.NewTranslationRepository,
)

var DomainServiceSet = wire.NewSet(
	service.NewContentService,
	service.NewMediaService,
	service.NewGenreService,
	service.NewInitService,
	service.NewGeneralInfoService,
	service.NewContentAccessService,
	service.NewMediaAccessService,
	service.NewEncodingService,
	service.NewTranslationService,
)

var CacheServiceSet = wire.NewSet(
	cache.NewMediaCacheService,
	cache.NewContentCacheService,
	cache.NewGenreCacheService,
	cache.NewGeneralInfoCacheService,
	cache.NewMediaAccessCacheService,
	cache.NewContentAccessCacheService,
	cache.NewLimitedCacheService,
	cache.NewHLSEncodingCacheService,
	cache.NewTranslationCacheService,
)

var ServiceSet = wire.NewSet(
	service.NewMediaOperations,
	service.NewContentOperations,
	service.NewGenreOperations,
	service.NewMediaCacheOperations,
	service.NewContentCacheOperations,
	service.NewGenreCacheOperations,
	service.NewGeneralInfoCacheOperations,
	service.NewContentAccessCacheOperations,
	service.NewMediaAccessCacheOperations,
	service.NewMediaPlayActionService,
	service.NewHLSEncodingCacheOperations,
	service.NewHLSMediaService,
	service.NewParameterService,
	service.NewTranslationCacheOperations,
)

var ControllerSet = wire.NewSet(
	api.NewCacheController,
	api.NewGenreController,
	api.NewContentController,
	api.NewHLSMediaTestController,
	api.NewHLSMediaController,
)

func InitializeApp() *echo.Echo {
	wire.Build(GeneralSet, RepositorySet, DomainServiceSet, CacheServiceSet, ServiceSet, ControllerSet, IntegrationSet)
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

func InitRoutes(
	cacheController *api.CacheController,
	genreController *api.GenreController,
	contentController *api.ContentController,
	hlsMediaController *api.HLSMediaController,
	hlsMediaTestController *api.HLSMediaTestController,
	initService service.IInitService,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware2.Static("/downloads/"))
	e.GET("/downloads/:file", func(c echo.Context) error {
		fileName := c.Param("file")
		return c.Attachment("downloads/"+fileName, fileName)
	})
	e.Use(middleware.ServiceContextMiddleware)
	cacheController.InitCacheRoutes(e)
	genreController.InitGenreRoutes(e)
	hlsMediaController.InitHLSMediaRoutes(e)
	hlsMediaTestController.InitHLSMediaTestRoutes(e)
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	initService.InitCache()
	go initService.InitJob()
	return e
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"calisthenics-root-api/api/v1"
	"calisthenics-root-api/config"
	"calisthenics-root-api/data"
	"calisthenics-root-api/data/repository"
	"calisthenics-root-api/middleware"
	"calisthenics-root-api/service"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitializeApp() *echo.Echo {
	db := NewDatabase()
	iUserRepository := repository.NewUserRepository(db)
	iUserService := service.NewUserService(iUserRepository)
	iAuthService := service.NewAuthService(iUserService)
	authController := v1.NewAuthController(iAuthService)
	userController := v1.NewUserController(iUserService)
	iRoleRepository := repository.NewRoleRepository(db)
	iRoleService := service.NewRoleService(iRoleRepository)
	roleController := v1.NewRoleController(iRoleService)
	iPrivilegeRepository := repository.NewPrivilegeRepository(db)
	iPrivilegeService := service.NewPrivilegeService(iPrivilegeRepository)
	privilegeController := v1.NewPrivilegeController(iPrivilegeService)
	iContentRepository := repository.NewContentRepository(db)
	iContentService := service.NewContentService(iContentRepository)
	iHelperContentRepository := repository.NewHelperContentRepository(db)
	iHelperContentService := service.NewHelperContentService(iHelperContentRepository, iContentService)
	iHelperContentOperations := service.NewHelperContentOperations(iHelperContentService, iContentService)
	iRequirementContentRepository := repository.NewRequirementContentRepository(db)
	iRequirementContentService := service.NewRequirementContentService(iRequirementContentRepository)
	iRequirementContentOperations := service.NewRequirementContentOperations(iRequirementContentService, iContentService)
	contentController := v1.NewContentController(iContentService, iHelperContentOperations, iRequirementContentOperations)
	iMediaRepository := repository.NewMediaRepository(db)
	iMediaService := service.NewMediaService(iMediaRepository)
	mediaController := v1.NewMediaController(iMediaService)
	iGenreTypeRepository := repository.NewGenreTypeRepository(db)
	iGenreTypeService := service.NewGenreTypeService(iGenreTypeRepository)
	genreTypeController := v1.NewGenreTypeController(iGenreTypeService)
	iGenreRepository := repository.NewGenreRepository(db)
	iGenreService := service.NewGenreService(iGenreRepository)
	iGenreContentRepository := repository.NewGenreContentRepository(db)
	iGenreContentService := service.NewGenreContentService(iGenreContentRepository)
	iGenreContentOperations := service.NewGenreContentOperations(iGenreService, iGenreContentService, iContentService)
	genreController := v1.NewGenreController(iGenreContentOperations, iGenreService)
	iTranslationRepository := repository.NewTranslationRepository(db)
	iTranslationService := service.NewTranslationService(iTranslationRepository)
	translationController := v1.NewTranslationController(iTranslationService)
	iContentTranslationRepository := repository.NewContentTranslationRepository(db)
	iContentTranslationService := service.NewContentTranslationService(iContentTranslationRepository)
	contentTranslationController := v1.NewContentTranslationController(iContentTranslationService)
	echoEcho := InitRoutes(authController, userController, roleController, privilegeController, contentController, mediaController, genreTypeController, genreController, translationController, contentTranslationController, iUserService)
	return echoEcho
}

// wire.go:

var GeneralSet = wire.NewSet(NewDatabase, InitRoutes)

var RepositorySet = wire.NewSet(repository.NewUserRepository, repository.NewPrivilegeRepository, repository.NewRoleRepository, repository.NewContentRepository, repository.NewMediaRepository, repository.NewGenreTypeRepository, repository.NewGenreRepository, repository.NewHelperContentRepository, repository.NewRequirementContentRepository, repository.NewGenreContentRepository, repository.NewTranslationRepository, repository.NewContentTranslationRepository)

var DomainServiceSet = wire.NewSet(service.NewUserService, service.NewPrivilegeService, service.NewRoleService, service.NewContentService, service.NewMediaService, service.NewHelperContentService, service.NewRequirementContentService, service.NewGenreTypeService, service.NewGenreService, service.NewGenreContentService, service.NewTranslationService, service.NewContentTranslationService)

var ServiceSet = wire.NewSet(service.NewAuthService, service.NewHelperContentOperations, service.NewRequirementContentOperations, service.NewGenreContentOperations)

var ControllerSet = wire.NewSet(v1.NewAuthController, v1.NewUserController, v1.NewRoleController, v1.NewPrivilegeController, v1.NewContentController, v1.NewMediaController, v1.NewGenreTypeController, v1.NewGenreController, v1.NewTranslationController, v1.NewContentTranslationController)

func NewDatabase() *gorm.DB {
	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", dbUser, dbName, dbPassword)))
	if err != nil {
		panic("Failed to connect to the database")
	}
	err = db.AutoMigrate(&data.User{}, &data.Role{}, &data.Privilege{}, &data.Content{}, &data.Media{}, &data.Genre{}, &data.GenreType{}, &data.Translation{}, &data.ContentTranslation{})
	if err != nil {
		panic("Failed to migrate database")
	}
	return db
}

func InitRoutes(
	authController *v1.AuthController,
	userController *v1.UserController,
	roleController *v1.RoleController,
	privilegeController *v1.PrivilegeController,
	contentController *v1.ContentController,
	mediaController *v1.MediaController,
	genreTypeController *v1.GenreTypeController,
	genreController *v1.GenreController,
	translationController *v1.TranslationController,
	contentTanslationController *v1.ContentTranslationController,
	userService service.IUserService) *echo.Echo {

	e := echo.New()
	e.Use(middleware.ServiceContextMiddleware)
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	authController.InitAuthRoutes(e)
	privilegeMiddleware := middleware.NewPrivilegeMiddleware(userService)
	authMiddleware := middleware.NewAuthenticationMiddleware()
	middlewareGroup := e.Group("/v1/users", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	userController.InitUserRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/roles", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	roleController.InitRoleRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/privileges", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	privilegeController.InitPrivilegeRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/contents", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	contentController.InitContentRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/medias", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	mediaController.InitMediaRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/genre-types", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	genreTypeController.InitGenreTypeRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/genres", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	genreController.InitGenreRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/translations", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	translationController.InitTranslationRoutes(middlewareGroup)

	middlewareGroup = e.Group("/v1/content-translations", authMiddleware.MiddlewareFunc, privilegeMiddleware.MiddlewareFunc)
	contentTanslationController.InitContentTranslationRoutes(middlewareGroup)

	return e
}

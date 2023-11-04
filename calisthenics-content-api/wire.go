//go:build wireinject
// +build wireinject

package main

import (
	"calisthenics-content-api/config"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GeneralSet = wire.NewSet(NewDatabase, InitRoutes)

func InitializeApp() *echo.Echo {
	wire.Build(GeneralSet)
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

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}
	return e
}

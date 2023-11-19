//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
)

var GeneralSet = wire.NewSet(InitRoutes)

var IntegrationSet = wire.NewSet()

var RepositorySet = wire.NewSet()

var DomainServiceSet = wire.NewSet()

var CacheServiceSet = wire.NewSet()

var ServiceSet = wire.NewSet()

var ControllerSet = wire.NewSet()

func InitializeApp() *echo.Echo {
	wire.Build(GeneralSet)
	return &echo.Echo{}
}

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware2.Static("/medias/"))
	e.GET("/medias/:file", func(c echo.Context) error {
		fileName := c.Param("file")
		return c.Attachment("medias/"+fileName, fileName)
	})
	return e
}

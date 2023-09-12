package main

import (
	_ "github.com/aerosystems/recaptcha-service/docs"
	"github.com/aerosystems/recaptcha-service/internal/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (app *Config) NewRouter() *echo.Echo {
	e := echo.New()

	docsGroup := e.Group("/docs")
	docsGroup.Use(middleware.BasicAuth)
	docsGroup.GET("/*", echoSwagger.WrapHandler)

	e.GET("/ping", app.BaseHandler.Ping)
	e.GET("/v1/validate", app.BaseHandler.Validate)
	e.GET("/v1/validate/v3", app.BaseHandler.ValidateV3)

	return e
}

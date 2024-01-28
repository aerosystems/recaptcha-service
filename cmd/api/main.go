package main

import (
	"fmt"
	"github.com/aerosystems/recaptcha-service/internal/handlers"
	"github.com/aerosystems/recaptcha-service/pkg/captcha"
	"github.com/aerosystems/recaptcha-service/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"os"
)

const webPort = 80

// @title reCAPTCHA Service
// @version 1.0.0
// @description A part of microservice infrastructure, who responsible for reCAPTCHA v2/v3 validation

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host gw.verifire.dev/recaptcha
// @schemes https
// @BasePath /
func main() {
	log := logger.NewLogger(os.Getenv("HOSTNAME"))

	recaptchaService := captcha.NewRecaptchaService(os.Getenv("RECAPTCHA_V2_SECRET"), os.Getenv("RECAPTCHA_V3_SECRET"))

	app := Config{
		BaseHandler: handlers.NewBaseHandler(log.Logger, recaptchaService),
	}

	e := app.NewRouter()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			log.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

	e.Use(middleware.Recover())

	errChan := make(chan error)

	go func() {
		log.Infof("starting recaptcha-service HTTP server on port %d\n", webPort)
		errChan <- e.Start(fmt.Sprintf(":%d", webPort))
	}()

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

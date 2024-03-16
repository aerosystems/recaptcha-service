//go:build wireinject
// +build wireinject

package main

import (
	"github.com/aerosystems/recaptcha-service/internal/config"
	HttpServer "github.com/aerosystems/recaptcha-service/internal/infrastructure/http"
	"github.com/aerosystems/recaptcha-service/internal/infrastructure/http/handlers"
	"github.com/aerosystems/recaptcha-service/internal/repository/google"
	"github.com/aerosystems/recaptcha-service/internal/usecases"
	"github.com/aerosystems/recaptcha-service/pkg/logger"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

//go:generate wire
func InitApp() *App {
	panic(wire.Build(
		wire.Bind(new(handlers.RecaptchaUsecase), new(*usecases.RecaptchaUsecase)),
		wire.Bind(new(usecases.RecaptchaRepository), new(*google.Api)),
		ProvideApp,
		ProvideLogger,
		ProvideConfig,
		ProvideHttpServer,
		ProvideLogrusLogger,
		ProvideBaseHandler,
		ProvideValidateHandler,
		ProvideRecaptchaUsecase,
		ProvideRecaptchaApi,
	))
}

func ProvideApp(log *logrus.Logger, cfg *config.Config, httpServer *HttpServer.Server) *App {
	panic(wire.Build(NewApp))
}

func ProvideLogger() *logger.Logger {
	panic(wire.Build(logger.NewLogger))
}

func ProvideConfig() *config.Config {
	panic(wire.Build(config.NewConfig))
}

func ProvideHttpServer(log *logrus.Logger, validateHandler *handlers.ValidateHandler) *HttpServer.Server {
	panic(wire.Build(HttpServer.NewServer))
}

func ProvideLogrusLogger(log *logger.Logger) *logrus.Logger {
	return log.Logger
}

func ProvideBaseHandler(log *logrus.Logger, cfg *config.Config) *handlers.BaseHandler {
	return handlers.NewBaseHandler(log, cfg.Mode)
}

func ProvideValidateHandler(baseHandler *handlers.BaseHandler, recaptchaUsecase handlers.RecaptchaUsecase) *handlers.ValidateHandler {
	return handlers.NewValidateHandler(baseHandler, recaptchaUsecase)
}

func ProvideRecaptchaUsecase(recaptchaRepo usecases.RecaptchaRepository) *usecases.RecaptchaUsecase {
	panic(wire.Build(usecases.NewRecaptchaUsecase))
}

func ProvideRecaptchaApi(cfg *config.Config) *google.Api {
	return google.NewApi(cfg.RecaptchaSecretV2, cfg.RecaptchaSecretV3)
}

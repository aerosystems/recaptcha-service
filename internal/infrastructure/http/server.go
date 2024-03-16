package HttpServer

import (
	"fmt"
	"github.com/aerosystems/recaptcha-service/internal/infrastructure/http/handlers"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const webPort = 80

type Server struct {
	log             *logrus.Logger
	echo            *echo.Echo
	validateHandler *handlers.ValidateHandler
}

func NewServer(
	log *logrus.Logger,
	validateHandler *handlers.ValidateHandler,
) *Server {
	return &Server{
		log:             log,
		echo:            echo.New(),
		validateHandler: validateHandler,
	}
}

func (s *Server) Run() error {
	s.setupMiddleware()
	s.setupRoutes()
	s.setupValidator()
	s.log.Infof("starting HTTP server recaptcha-service on port %d\n", webPort)
	return s.echo.Start(fmt.Sprintf(":%d", webPort))
}

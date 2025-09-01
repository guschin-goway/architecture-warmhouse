package server

import (
	"github.com/labstack/echo/v4"

	"temperature/internal/config"
	"temperature/internal/logger"
	"temperature/internal/transport/http/api"
)

type Server struct {
	cfg    *config.Config
	logger *logger.Logger
	echo   *echo.Echo
}

func New(cfg *config.Config, logger *logger.Logger, handlers *APIHandlers, middlewares []echo.MiddlewareFunc) *Server {
	s := &Server{
		cfg:    cfg,
		logger: logger,
		echo:   echo.New(),
	}
	s.echo.Use(middlewares...)
	api.RegisterHandlers(s.echo, handlers)
	return s
}

// Start запускает HTTP сервер
func (s *Server) Start() error {
	if s.cfg.App.SSLEnable {
		return s.echo.StartTLS(s.cfg.App.AppPort(), s.cfg.App.SSLConfig.CertPath, s.cfg.App.SSLConfig.KeyPath)
	}
	return s.echo.Start(s.cfg.App.AppPort())
}

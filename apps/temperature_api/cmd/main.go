package main

import (
	"context"
	bLog "log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"temperature/internal/config"
	"temperature/internal/logger"
	"temperature/internal/service"
	"temperature/internal/transport/http/server"
	"temperature/internal/transport/http/temperature"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		bLog.Fatalf("Ошибка при загрузке конфигурационного файла")
	}

	// Initialize logger
	log, err := logger.New(logger.Config{
		Level:        cfg.Log.Level,
		Format:       cfg.Log.Format,
		ServiceName:  cfg.App.Name,
		Environment:  cfg.App.Env,
		EnableCaller: cfg.Log.EnableCaller,
	})
	if err != nil {
		bLog.Fatalf("Ошибка при инициализации логгера")
	}
	defer log.Sync()

	// Initialize Echo
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	//service
	temperatureService := service.NewTemperatureService()

	//handler
	temperatureHandler := temperature.NewHandler(log, temperatureService)

	// Start server
	go func() {
		log.Info("Starting server",
			zap.String("service", cfg.App.Name),
			zap.String("version", cfg.App.Version),
			zap.String("environment", cfg.App.Env),
			zap.String("address", cfg.App.AppAddress()),
		)

		handlers := server.NewAPIHandlers(temperatureHandler)
		middlewares := server.CreateMiddlewares()
		// Create server instance
		apiServer := server.New(cfg, log, handlers, middlewares)

		log.Info("Temperature API")
		if err := apiServer.Start(); err != nil && err != http.ErrServerClosed {
			log.Error("Ошибка в работе сервиса", zap.Error(err))
		}

	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Shutdown gracefully
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	log.Info("Shutting down HTTP SERVER")
	if err := e.Shutdown(ctx); err != nil {
		log.Error("Error during server shutdown", zap.Error(err))
	}

}

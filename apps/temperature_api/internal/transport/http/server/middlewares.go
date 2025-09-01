package server

import (
	"compress/gzip"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateMiddlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.Recover(),
		middleware.Logger(),
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: gzip.BestSpeed,
		}),
	}
}

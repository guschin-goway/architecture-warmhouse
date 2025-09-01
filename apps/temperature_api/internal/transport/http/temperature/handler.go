package temperature

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"temperature/internal/logger"
	"temperature/internal/service"
)

type Handler struct {
	log                *logger.Logger
	temperatureService *service.TemperatureService
}

func NewHandler(log *logger.Logger, temperatureService *service.TemperatureService) *Handler {
	return &Handler{
		log:                log,
		temperatureService: temperatureService,
	}
}

// GetTemperatureBySensorID (GET /temperature/{sensorID})
func (h *Handler) GetTemperatureBySensorID(c echo.Context, sensorID string) error {
	tempResp, err := h.temperatureService.GetTemp(sensorID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tempResp)
}

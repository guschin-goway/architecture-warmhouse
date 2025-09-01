package server

import (
	"temperature/internal/transport/http/temperature"
)

type TemperatureHandler = temperature.Handler

type APIHandlers struct {
	TemperatureHandler
}

func NewAPIHandlers(temperatureHandler *temperature.Handler) *APIHandlers {
	return &APIHandlers{
		TemperatureHandler: TemperatureHandler(*temperatureHandler),
	}
}

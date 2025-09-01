package service

import (
	"math/rand"
	"time"

	"temperature/internal/transport/http/api"
)

type TemperatureService struct{}

func NewTemperatureService() *TemperatureService {
	return &TemperatureService{}
}

func (o *TemperatureService) GetTemp(sensorID string) (api.TemperatureResponse, error) {

	minTemp := -30
	maxTemp := 40

	// Генерация случайной температуры
	temp := rand.Intn(maxTemp-minTemp+1) + minTemp
	location, sensorID := GetLocationAndSensorIDByLocation(sensorID)
	return api.TemperatureResponse{
		Location:  location,
		SensorId:  sensorID,
		Timestamp: time.Now(),
		Value:     float32(temp),
	}, nil
}

func GetLocationAndSensorIDByLocation(sensorID string) (string, string) {
	var location string
	switch sensorID {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}
	return location, sensorID
}

package models

import (
	"gorm.io/gorm"
)

type Temperature struct {
	gorm.Model
	SensorId           int64 `json:"sensor_id"`
	CurrentTemperature int   `json:"current_temperature"`
	Timestamp          int64 `json:"timestamp"`
}

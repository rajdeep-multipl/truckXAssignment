package models

import (
	"gorm.io/gorm"
)

type AggregatedTemperature struct {
	gorm.Model
	MaxTemperature int     `json:"max_temperature"`
	MinTemperature int     `json:"min_temperature"`
	AvgTemperature float64 `json:"avg_temperature"`
	SensorId       int64   `json:"sensor_id"`
}

type AggregatedTemperatureReq struct {
	SensorId  *int64
	StartDate string
	EndDate   string
	StartTime string
	EndTime   string
}

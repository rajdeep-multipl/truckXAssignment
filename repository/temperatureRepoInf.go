package repository

import (
	"scaleX/assignment/models"
	"time"
)

type TemperatureRepoInf interface {
	AddTemperatureRepo(temperature *models.Temperature) error
	GetAggregateValuesForEachSensor() ([]models.AggregatedTemperature, error)
	GetAggregatedDataOfSensorForTimeRange(sensorId int64, startTime time.Time, endTime time.Time) ([]models.AggregatedTemperature, error)
	AddAggregatedTemperatures(aggregatedTemperatures []models.AggregatedTemperature) error
	GetAggregatedDataOfSensor(sensorId int64) ([]models.AggregatedTemperature, error)
}

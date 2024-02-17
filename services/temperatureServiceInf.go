package services

import "scaleX/assignment/models"

type TemperatureServiceInf interface {
	AddTemperature(temperature *models.Temperature) error
	GetAggregatedTemperatureValue(aggregatedReqObj models.AggregatedTemperatureReq) ([]models.AggregatedTemperature, error)
}

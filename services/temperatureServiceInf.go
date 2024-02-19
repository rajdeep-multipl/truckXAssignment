package services

import "scaleX/assignment/models"

// in this interface I have declared method signatures of the necessary temperature methods
type TemperatureServiceInf interface {
	AddTemperature(temperature *models.Temperature) error
	GetAggregatedTemperatureValue(aggregatedReqObj models.AggregatedTemperatureReq) ([]models.AggregatedTemperature, error)
}

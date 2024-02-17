package services

import (
	"errors"
	"fmt"
	"log"
	"scaleX/assignment/models"
	"scaleX/assignment/repository"
	"time"

	"gorm.io/gorm"
)

type TemperatureService struct {
	DB *gorm.DB
}

func getTemperatureRepoInstance(db *gorm.DB) repository.TemperatureRepoInf {
	return &repository.TemperatureRepo{DB: db}
}

func getSensorRepoInstance(db *gorm.DB) repository.SensorRepoInf {
	return &repository.SensorRepo{DB: db}
}

func (handler *TemperatureService) AddTemperature(temperature *models.Temperature) error {
	temperatureRepo := getTemperatureRepoInstance(handler.DB)
	sensorRepo := getSensorRepoInstance(handler.DB)

	// check if there is any data at the sersor table
	sensorDataExists, err := sensorRepo.CheckIfSensorAlreadyExists(temperature.SensorId)
	if err != nil {
		return err
	}

	if !sensorDataExists {
		sensorRepo.AddSensorRepo(&models.Sensor{ID: temperature.SensorId})
	}

	err = temperatureRepo.AddTemperatureRepo(temperature)
	if err != nil {
		return err
	}
	return nil
}

func (handler TemperatureService) GetAggregatedTemperatureValue(aggregatedReqObj models.AggregatedTemperatureReq) ([]models.AggregatedTemperature, error) {
	temperatureRepo := getTemperatureRepoInstance(handler.DB)

	if aggregatedReqObj.SensorId == nil {
		log.Printf("error at services/GetAggregatedTemperatureValue: %s", "senssor id can't be empty")
		return nil, errors.New("sensor id can't be empty")
	}

	if len(aggregatedReqObj.StartDate) == 0 && len(aggregatedReqObj.EndDate) == 0 && len(aggregatedReqObj.StartDate) == 0 && len(aggregatedReqObj.EndTime) == 0 {
		aggregatedTemps, err := temperatureRepo.GetAggregatedDataOfSensor(*aggregatedReqObj.SensorId)
		if err != nil {
			log.Printf("error at services/GetAggregatedTemperatureValue: %s", err.Error())
			return nil, err
		}

		return aggregatedTemps, nil
	}

	startDate, err := time.Parse("2006-01-02", aggregatedReqObj.StartDate)
	if err != nil {
		if len(aggregatedReqObj.StartTime) == 0 {
			log.Printf("error at services/GetAggregatedTemperatureValue: %s", "start time or start data is not given")
			return nil, errors.New("start time or start date is not given")
		}
		currentDate := time.Now().Format("2006-01-02")
		combinedTime := fmt.Sprintf("%sT%s", currentDate, aggregatedReqObj.StartTime)
		startDate, _ = time.Parse("2006-01-02T15:04:05", combinedTime)
	}

	endDate, err := time.Parse("2006-01-02", aggregatedReqObj.EndDate)
	if err != nil {
		if len(aggregatedReqObj.EndTime) == 0 {
			log.Printf("error at services/GetAggregatedTemperatureValue: %s", "end time or end data is not given")
			return nil, errors.New("end time or end date is not given")
		}
		currentDate := time.Now().Format("2006-01-02")
		combinedTime := fmt.Sprintf("%sT%s", currentDate, aggregatedReqObj.EndTime)
		endDate, _ = time.Parse("2006-01-02T15:04:05", combinedTime)
	}

	aggregatedTemps, err := temperatureRepo.GetAggregatedDataOfSensorForTimeRange(*aggregatedReqObj.SensorId, startDate, endDate)
	if err != nil {
		log.Printf("error at services/GetAggregatedTemperatureValue: %s", err.Error())
		return nil, err
	}

	return aggregatedTemps, nil
}

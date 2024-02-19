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

// getting the temperature repository instance
func getTemperatureRepoInstance(db *gorm.DB) repository.TemperatureRepoInf {
	return &repository.TemperatureRepo{DB: db}
}

// getting the sensor repo instance
func getSensorRepoInstance(db *gorm.DB) repository.SensorRepoInf {
	return &repository.SensorRepo{DB: db}
}

// this is the service function for adding temperature. This function holds the necessary business logic for adding temeperature
func (handler *TemperatureService) AddTemperature(temperature *models.Temperature) error {
	temperatureRepo := getTemperatureRepoInstance(handler.DB)
	sensorRepo := getSensorRepoInstance(handler.DB)

	// check if there is any data at the sersor table
	sensorDataExists, err := sensorRepo.CheckIfSensorAlreadyExists(temperature.SensorId)
	if err != nil {
		return err
	}

	// if it's a new sensor data we are adding the data into the sensor table.
	if !sensorDataExists {
		sensorRepo.AddSensorRepo(&models.Sensor{ID: temperature.SensorId})
	}

	// calling repository function for adding temeperature
	err = temperatureRepo.AddTemperatureRepo(temperature)
	if err != nil {
		return err
	}
	return nil
}

// This function gets aggregated data from aggregated temperature table
func (handler TemperatureService) GetAggregatedTemperatureValue(aggregatedReqObj models.AggregatedTemperatureReq) ([]models.AggregatedTemperature, error) {
	temperatureRepo := getTemperatureRepoInstance(handler.DB)

	// checking for sensor id
	if aggregatedReqObj.SensorId == nil {
		log.Printf("error at services/GetAggregatedTemperatureValue: %s", "senssor id can't be empty")
		return nil, errors.New("sensor id can't be empty")
	}

	// if no time range is given we are only returing data that matches the sensor id from the aggregated_temperature table
	if len(aggregatedReqObj.StartDate) == 0 && len(aggregatedReqObj.EndDate) == 0 && len(aggregatedReqObj.StartDate) == 0 && len(aggregatedReqObj.EndTime) == 0 {
		// getting aggregated data only for sensor id
		aggregatedTemps, err := temperatureRepo.GetAggregatedDataOfSensor(*aggregatedReqObj.SensorId)
		if err != nil {
			log.Printf("error at services/GetAggregatedTemperatureValue: %s", err.Error())
			return nil, err
		}

		return aggregatedTemps, nil
	}

	// if start data is not given then we are taking todays date with start time for query purposes
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

	// if end date is not given we are taking end time with the current day's date for query purposes
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

	// getting aggregated data with in the time range
	aggregatedTemps, err := temperatureRepo.GetAggregatedDataOfSensorForTimeRange(*aggregatedReqObj.SensorId, startDate, endDate)
	if err != nil {
		log.Printf("error at services/GetAggregatedTemperatureValue: %s", err.Error())
		return nil, err
	}

	return aggregatedTemps, nil
}

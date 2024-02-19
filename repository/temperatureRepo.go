package repository

import (
	"log"
	"scaleX/assignment/models"
	"time"

	"gorm.io/gorm"
)

type TemperatureRepo struct {
	DB *gorm.DB
}

// This function adds new temperature record in the temeperature table
func (handler *TemperatureRepo) AddTemperatureRepo(temperature *models.Temperature) error {
	err := handler.DB.Create(&temperature).Error
	if err != nil {
		log.Printf("error at AddTemperatureRepo, error: %s", err.Error())
		return err
	}
	return nil
}

// This function returns aggregated temperature value based on the sensor id from temperature table
func (hanlder *TemperatureRepo) GetAggregateValuesForEachSensor() ([]models.AggregatedTemperature, error) {
	var aggregatedSensorValues []models.AggregatedTemperature

	err := hanlder.DB.Model(&models.Temperature{}).
		Select(`sensor_id, min(current_temperature) as min_temperature, max(current_temperature) as max_temperature, avg(current_temperature) as avg_temperature`).
		Where("timestamp >= ?", time.Now().Add(-time.Hour).UnixNano()).
		Group("sensor_id").
		Find(&aggregatedSensorValues).Error
	if err != nil {
		log.Printf("error at GetAggregateValuesForEachSensor, %s", err.Error())
		return nil, err
	}

	return aggregatedSensorValues, nil
}

// This function adds new record to the aggregated_temperature table
func (handler *TemperatureRepo) AddAggregatedTemperatures(aggregatedTemperatures []models.AggregatedTemperature) error {
	err := handler.DB.Create(&aggregatedTemperatures).Error
	if err != nil {
		log.Printf("error at AddAggregatedTemperatures, %s", err.Error())
		return err
	}
	return nil
}

// This function get the aggregated tempereture value of a sensor within a time range
func (handler *TemperatureRepo) GetAggregatedDataOfSensorForTimeRange(sensorId int64, startTime time.Time, endTime time.Time) ([]models.AggregatedTemperature, error) {
	var aggregatedTemperatures []models.AggregatedTemperature

	err := handler.DB.Model(&models.AggregatedTemperature{}).
		Where("sensor_id = ? AND created_at BETWEEN ? AND ?", sensorId, startTime.Format("2006-01-02T15:04:05"), endTime.Format("2006-01-02T15:04:05")).
		Find(&aggregatedTemperatures).Error
	if err != nil {
		log.Printf("error at GetAggregatedDataOfSensorForTimeRange, %s", err.Error())
		return nil, err
	}

	return aggregatedTemperatures, nil
}

// This function returns the aggregated temperature value of a sensor
func (handler *TemperatureRepo) GetAggregatedDataOfSensor(sensorId int64) ([]models.AggregatedTemperature, error) {
	var aggregatedTemperatures []models.AggregatedTemperature

	err := handler.DB.Model(&models.AggregatedTemperature{}).
		Where("sensor_id = ?", sensorId).
		Find(&aggregatedTemperatures).Error
	if err != nil {
		log.Printf("error at GetAggregatedDataOfSensor, %s", err.Error())
		return nil, err
	}

	return aggregatedTemperatures, nil
}

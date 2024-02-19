package repository

import (
	"log"
	"scaleX/assignment/models"

	"gorm.io/gorm"
)

type SensorRepo struct {
	DB *gorm.DB
}

// Adds new record in the sensor table
func (handler *SensorRepo) AddSensorRepo(sensor *models.Sensor) error {
	// using gorm Create function which maps structs and tables
	err := handler.DB.Create(&sensor).Error
	if err != nil {
		log.Printf("error at repository/AddSensorRepo: %s\n", err.Error())
		return err
	}
	return nil
}

// Checks if there is already a record for the sensor id
func (handler *SensorRepo) CheckIfSensorAlreadyExists(sensorId int64) (bool, error) {
	var count int64
	// Using gorm Where and Count function to get the count of the sensor id
	err := handler.DB.Model(&models.Sensor{}).Where("id = ?", sensorId).Count(&count).Error
	if err != nil {
		log.Printf("error at CheckIfSensorAlreadyExists: %s\n", err.Error())
		return false, err
	}

	// If the count is more than one we are returning a bool response
	if count > 0 {
		return true, nil
	}

	return false, nil

}

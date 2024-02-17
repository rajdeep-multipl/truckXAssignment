package repository

import (
	"log"
	"scaleX/assignment/models"

	"gorm.io/gorm"
)

type SensorRepo struct {
	DB *gorm.DB
}

func (handler *SensorRepo) AddSensorRepo(sensor *models.Sensor) error {
	err := handler.DB.Create(&sensor).Error
	if err != nil {
		log.Printf("error at repository/AddSensorRepo: %s\n", err.Error())
		return err
	}
	return nil
}

func (handler *SensorRepo) CheckIfSensorAlreadyExists(sensorId int64) (bool, error) {
	var count int64
	err := handler.DB.Model(&models.Sensor{}).Where("id = ?", sensorId).Count(&count).Error
	if err != nil {
		log.Printf("error at CheckIfSensorAlreadyExists: %s\n", err.Error())
		return false, err
	}

	if count > 0 {
		log.Printf("error at CheckIfSensorAlreadyExists: %s\n", err.Error())
		return true, nil
	}

	return false, nil

}

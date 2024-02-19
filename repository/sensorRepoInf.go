package repository

import "scaleX/assignment/models"

// This interface holds the method signatures of the repository functions of Sensor
type SensorRepoInf interface {
	AddSensorRepo(sensor *models.Sensor) error
	CheckIfSensorAlreadyExists(sensorId int64) (bool, error)
}

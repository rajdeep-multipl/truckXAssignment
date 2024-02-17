package repository

import "scaleX/assignment/models"

type SensorRepoInf interface {
	AddSensorRepo(sensor *models.Sensor) error
	CheckIfSensorAlreadyExists(sensorId int64) (bool, error)
}

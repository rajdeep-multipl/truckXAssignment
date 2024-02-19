package controllers

import "github.com/gin-gonic/gin"

// this interfaces declares necessary controller functions for temperatures
type TemperatureControllerInf interface {
	PostTemperature(c *gin.Context)
	GetAggregatedValueOfSensor(c *gin.Context)
}

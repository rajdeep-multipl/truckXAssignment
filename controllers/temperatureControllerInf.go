package controllers

import "github.com/gin-gonic/gin"

type TemperatureControllerInf interface {
	PostTemperature(c *gin.Context)
	GetAggregatedValueOfSensor(c *gin.Context)
}

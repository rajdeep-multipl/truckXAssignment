package routers

import (
	"scaleX/assignment/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// This function returns the temperature controller instance
func getTemperatureControllerInstance(db *gorm.DB) controllers.TemperatureControllerInf {
	return &controllers.TemperatureController{DB: db}
}

// This functions declares the necessary routes
func SetupRouter(DB *gorm.DB) *gin.Engine {
	tempController := getTemperatureControllerInstance(DB)

	r := gin.Default()
	api := r.Group("/api")
	{

		api.POST("/temperature", tempController.PostTemperature)
		api.GET("/aggregate", tempController.GetAggregatedValueOfSensor)

	}

	return r
}

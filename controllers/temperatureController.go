package controllers

import (
	"log"
	"net/http"
	"scaleX/assignment/dto"
	"scaleX/assignment/models"
	"scaleX/assignment/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TemperatureController struct {
	DB *gorm.DB
}

func getTemperatureServiceInstance(db *gorm.DB) services.TemperatureServiceInf {
	return &services.TemperatureService{DB: db}
}

func (hanlder TemperatureController) PostTemperature(c *gin.Context) {
	var temperatureReqDto dto.TemperatureReqDTO
	var temperatureResDto dto.TemperatureResDTO
	var temperature models.Temperature
	temparatureService := getTemperatureServiceInstance(hanlder.DB)

	err := c.BindJSON(&temperatureReqDto)
	if err != nil {
		log.Printf("error at controller/PostTemperature(): %s\n", err.Error())
		temperatureResDto.ErrorMsg = "failed to bind user details sent via request body"
		temperatureResDto.Status = http.StatusBadRequest
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}

	if temperatureReqDto.SensorId == nil {
		log.Printf("error at controller/PostTemperature(): sensor id is not given\n")
		temperatureResDto.ErrorMsg = "sensor id is not given"
		temperatureResDto.Status = http.StatusBadRequest
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}
	temperature.SensorId = *temperatureReqDto.SensorId

	if temperatureReqDto.TemperatureValue == nil {
		log.Printf("error at controller/PostTemperature(): temperature value is not given\n")
		temperatureResDto.ErrorMsg = "temperature value is not given"
		temperatureResDto.Status = http.StatusBadRequest
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}
	temperature.CurrentTemperature = *temperatureReqDto.TemperatureValue

	if temperatureReqDto.Timestamp == nil {
		timeNow := time.Now().UnixNano()
		temperatureReqDto.Timestamp = &timeNow
	}
	temperature.Timestamp = *temperatureReqDto.Timestamp

	err = temparatureService.AddTemperature(&temperature)
	if err != nil {
		log.Printf("error at controller/PostTemperature(): %s", err.Error())
		temperatureResDto.ErrorMsg = "some error occured while adding data"
		temperatureResDto.Status = http.StatusInternalServerError
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}

	temperatureResDto.Message = "Successfully inserted the data into database"
	temperatureResDto.Status = http.StatusAccepted
	c.JSON(temperatureResDto.Status, temperatureResDto)
}

func (handler TemperatureController) GetAggregatedValueOfSensor(c *gin.Context) {
	var aggregatedTempReqDto dto.AggregatedTemperatureReqDto
	var aggregatedTempResDto dto.AggregatedTemperatureResDto
	var aggregatedTempReq models.AggregatedTemperatureReq
	temparatureService := getTemperatureServiceInstance(handler.DB)

	err := c.ShouldBindJSON(&aggregatedTempReqDto)
	if err != nil {
		log.Printf("error at temperature controller/GetAggregatedValueOfSensor: %s", err.Error())
		aggregatedTempResDto.ErrorMsg = "failed to bind request body"
		aggregatedTempResDto.Status = http.StatusBadGateway
		c.JSON(aggregatedTempResDto.Status, aggregatedTempResDto)
		return
	}

	aggregatedTempReq.SensorId = aggregatedTempReqDto.SensorId
	aggregatedTempReq.StartDate = aggregatedTempReqDto.StartDate
	aggregatedTempReq.EndDate = aggregatedTempReqDto.EndDate
	aggregatedTempReq.StartTime = aggregatedTempReqDto.StartTime
	aggregatedTempReq.EndTime = aggregatedTempReqDto.EndTime

	aggregatedTemp, err := temparatureService.GetAggregatedTemperatureValue(aggregatedTempReq)
	if err != nil {
		log.Printf("error at controller/GetAggregatedValueOfSensor: %s", err.Error())
		aggregatedTempResDto.ErrorMsg = "could not retrive the temperature value from database"
		aggregatedTempResDto.Status = http.StatusInternalServerError
		c.JSON(aggregatedTempResDto.Status, aggregatedTempResDto)
		return
	}

	aggregatedTempResDto.Status = http.StatusAccepted
	var aggregatedTemperaturesRes []dto.AggregatedTemperature
	for _, v := range aggregatedTemp {
		aggregatedTemperaturesRes = append(aggregatedTemperaturesRes, dto.AggregatedTemperature{
			SensorId:       v.SensorId,
			MaxTemperature: v.MaxTemperature,
			MinTemperature: v.MinTemperature,
			AvgTemperature: v.AvgTemperature,
			Time:           v.CreatedAt.Format("15:04:05"),
			Date:           v.CreatedAt.Format("2006-01-02"),
		})
	}

	aggregatedTempResDto.AggregatedTemperatures = aggregatedTemperaturesRes

	c.JSON(aggregatedTempResDto.Status, aggregatedTempResDto)
}

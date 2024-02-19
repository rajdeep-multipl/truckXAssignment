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

// This function returns the temperature service instance
func getTemperatureServiceInstance(db *gorm.DB) services.TemperatureServiceInf {
	return &services.TemperatureService{DB: db}
}

// This is the controller for adding new temerature values. We retrive the request body and do some basic checks
func (hanlder TemperatureController) PostTemperature(c *gin.Context) {
	// I am using request dto to bind the incoming request payload in this structure.
	var temperatureReqDto dto.TemperatureReqDTO

	// I am using response dto to send the response.
	var temperatureResDto dto.TemperatureResDTO
	var temperature models.Temperature
	temparatureService := getTemperatureServiceInstance(hanlder.DB)

	// binding the requst payload into the request dto
	err := c.BindJSON(&temperatureReqDto)
	if err != nil {
		log.Printf("error at controller/PostTemperature() while binding: %s\n", err.Error())
		temperatureResDto.ErrorMsg = "failed to bind user details sent via request body"
		temperatureResDto.Status = http.StatusBadRequest
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}

	// checking for sensor info, because it is mandatory
	if temperatureReqDto.SensorId == nil {
		log.Printf("error at controller/PostTemperature() while checking the sensor Id: sensor id is not given\n")
		temperatureResDto.ErrorMsg = "sensor id is not given"
		temperatureResDto.Status = http.StatusBadRequest
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}
	temperature.SensorId = *temperatureReqDto.SensorId

	// checking for temeperature value, because it is mandatory
	if temperatureReqDto.TemperatureValue == nil {
		log.Printf("error at controller/PostTemperature() while checking the temeperature value: temperature value is not given\n")
		temperatureResDto.ErrorMsg = "temperature value is not given"
		temperatureResDto.Status = http.StatusBadRequest
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}
	temperature.CurrentTemperature = *temperatureReqDto.TemperatureValue

	// checking if the timestamp value is given. if not given I am adding the current timestamp.
	if temperatureReqDto.Timestamp == nil {
		timeNow := time.Now().UnixNano()
		temperatureReqDto.Timestamp = &timeNow
	}
	temperature.Timestamp = *temperatureReqDto.Timestamp

	// Adding the temperature value to the temperature table
	err = temparatureService.AddTemperature(&temperature)
	if err != nil {
		log.Printf("error at controller/PostTemperature() while adding temeperature: %s", err.Error())
		temperatureResDto.ErrorMsg = "some error occured while adding data"
		temperatureResDto.Status = http.StatusInternalServerError
		c.JSON(temperatureResDto.Status, temperatureResDto)
		return
	}

	// returning the json response
	temperatureResDto.Message = "Successfully inserted the data into database"
	temperatureResDto.Status = http.StatusAccepted
	c.JSON(temperatureResDto.Status, temperatureResDto)
}

// This method gets aggregated value from aggregated_temeperature value with in a given time range and sensor id or
// just by sensor_id
func (handler TemperatureController) GetAggregatedValueOfSensor(c *gin.Context) {
	var aggregatedTempReqDto dto.AggregatedTemperatureReqDto
	var aggregatedTempResDto dto.AggregatedTemperatureResDto
	var aggregatedTempReq models.AggregatedTemperatureReq
	temparatureService := getTemperatureServiceInstance(handler.DB)

	//Binding request payload
	err := c.ShouldBindJSON(&aggregatedTempReqDto)
	if err != nil {
		log.Printf("error at temperature controller/GetAggregatedValueOfSensor while binding request payload: %s", err.Error())
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

	// fetching the aggregated values
	aggregatedTemp, err := temparatureService.GetAggregatedTemperatureValue(aggregatedTempReq)
	if err != nil {
		log.Printf("error at controller/GetAggregatedValueOfSensor while getting the aggregated value: %s", err.Error())
		aggregatedTempResDto.ErrorMsg = "could not retrive the temperature value from database"
		aggregatedTempResDto.Status = http.StatusInternalServerError
		c.JSON(aggregatedTempResDto.Status, aggregatedTempResDto)
		return
	}

	// building the response. putting nessesary data from returned object to response dto
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

	// sending json response
	c.JSON(aggregatedTempResDto.Status, aggregatedTempResDto)
}

package dto

type TemperatureReqDTO struct {
	SensorId         *int64 `json:"sensor_id" binding:"required"`
	TemperatureValue *int   `json:"temperature_value" binding:"required"`
	Timestamp        *int64 `json:"time_stamp"`
}

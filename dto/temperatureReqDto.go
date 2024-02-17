package dto

type TemperatureReqDTO struct {
	SensorId         *int64 `json:"sensor_id"`
	TemperatureValue *int   `json:"temperature_value"`
	Timestamp        *int64 `json:"time_stamp"`
}

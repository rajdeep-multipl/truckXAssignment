package dto

type AggregatedTemperatureReqDto struct {
	SensorId  *int64 `json:"sensor_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

package dto

type AggregatedTemperatureResDto struct {
	ErrorDTO
	Status                 int                     `json:"status"`
	AggregatedTemperatures []AggregatedTemperature `json:"aggregated_temperatures"`
}

type AggregatedTemperature struct {
	SensorId       int64   `json:"sensor_id"`
	MaxTemperature int     `json:"max_temperature"`
	MinTemperature int     `json:"min_temperature"`
	AvgTemperature float64 `json:"avg_temperature"`
	Time           string  `json:"time"`
	Date           string  `json:"date"`
}

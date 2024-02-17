package dto

type TemperatureResDTO struct {
	ErrorDTO
	Status  int    `json:"status"`
	Message string `json:"message"`
}

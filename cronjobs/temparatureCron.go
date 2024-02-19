package cronjobs

import (
	"log"
	"scaleX/assignment/repository"

	"github.com/robfig/cron"
	"gorm.io/gorm"
)

type TemperatureCron struct {
	DB *gorm.DB
}

// This functions returns an object of Temperature Repo. Which is used for calling temperature related db calls.
func getTemperatureRepoInstance(db *gorm.DB) repository.TemperatureRepoInf {
	return &repository.TemperatureRepo{DB: db}
}

// This function deployes cron related to temperature
func (handler TemperatureCron) RunTemperatureCronJobs() {
	temperatureRepo := getTemperatureRepoInstance(handler.DB)

	c := cron.New()

	// This cron runs every hour
	c.AddFunc("@every 1h", func() { GatherNewTemperatureValuesForSensor(temperatureRepo) })

	c.Start()
}

/*
This function runs on hourly basis. It fetches aggregated records from the temperature table
for the last 1 hour and if there is one or more records, I am adding those records in the aggregated_temperatures table
*/
func GatherNewTemperatureValuesForSensor(temperatureRepo repository.TemperatureRepoInf) {
	// We are getting aggreated value for each sensor from the temperature table
	aggregatedTemperatures, err := temperatureRepo.GetAggregateValuesForEachSensor()
	if err != nil {
		log.Printf("error at GatherNewTemperatureValuesForSensor, %s\n", err.Error())
		return
	}

	// inserting aggregated records if there is any new data available
	if len(aggregatedTemperatures) > 0 {
		err = temperatureRepo.AddAggregatedTemperatures(aggregatedTemperatures)
		if err != nil {
			log.Printf("error at GatherNewTemperatureValuesForSensor, %s\n", err.Error())
			return
		}
	}

}

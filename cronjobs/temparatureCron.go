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

func getTemperatureRepoInstance(db *gorm.DB) repository.TemperatureRepoInf {
	return &repository.TemperatureRepo{DB: db}
}

func (handler TemperatureCron) RunTemperatureCronJobs() {
	temperatureRepo := getTemperatureRepoInstance(handler.DB)

	c := cron.New()

	c.AddFunc("@every 1h", func() { GatherNewTemperatureValuesForSensor(temperatureRepo) })

	c.Start()
}

func GatherNewTemperatureValuesForSensor(temperatureRepo repository.TemperatureRepoInf) {
	aggregatedTemperatures, err := temperatureRepo.GetAggregateValuesForEachSensor()
	if err != nil {
		log.Printf("error at GatherNewTemperatureValuesForSensor, %s", err.Error())
	}

	if len(aggregatedTemperatures) > 0 {
		temperatureRepo.AddAggregatedTemperatures(aggregatedTemperatures)
	}

}

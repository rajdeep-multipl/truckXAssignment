package cronjobs

import "gorm.io/gorm"

func getTemperatureCronInstance(db *gorm.DB) TemperatureCronInf {
	return TemperatureCron{DB: db}
}

func RunCronJobs(DB *gorm.DB) {
	temperatureCron := getTemperatureCronInstance(DB)
	temperatureCron.RunTemperatureCronJobs()
}

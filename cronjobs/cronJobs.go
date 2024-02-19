package cronjobs

import "gorm.io/gorm"

// This function returns the object of temperature cron interface.
func getTemperatureCronInstance(db *gorm.DB) TemperatureCronInf {
	return TemperatureCron{DB: db}
}

// I am using this function to call the all the cron object instance methods. for now I am only calling the temperature instance.
func RunCronJobs(DB *gorm.DB) {
	temperatureCron := getTemperatureCronInstance(DB)
	temperatureCron.RunTemperatureCronJobs()
}

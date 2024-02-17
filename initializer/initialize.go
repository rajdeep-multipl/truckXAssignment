package initializer

import (
	"log"
	"scaleX/assignment/database"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type InitialEntities struct {
	DB *gorm.DB
}

func getDBInstance() database.DB {
	return new(database.SqlDB)
}

func conncetToDB() *gorm.DB {
	return getDBInstance().ConnectDB()
}

func loadConfigs() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func Initialize() InitialEntities {
	loadConfigs()

	initialSetup := InitialEntities{
		DB: conncetToDB(),
	}

	return initialSetup
}

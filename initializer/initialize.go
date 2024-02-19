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
	/*
		getDBInstance method returns object of the database.DB interface. We are calling ConnectDB() method on it to connect to
		the local sql database
	*/
	return getDBInstance().ConnectDB()
}

func loadConfigs() {
	// we are loading config from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func Initialize() InitialEntities {
	// By calling this function we are loading the configs from th .env file, to use them in our application.
	loadConfigs()

	// Initial setup object holds the objects which are initiated at the beginning of the program.
	// For now it is just holding the DB instance.
	initialSetup := InitialEntities{
		DB: conncetToDB(),
	}

	return initialSetup
}

package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlDB struct {
	DB *gorm.DB
}

var once sync.Once

/*
This method returns a sql db driver instance
*/
func (handler *SqlDB) ConnectDB() *gorm.DB {
	/*
		I want exactly one DB instance across the whole application. Using once.Do we can achive that. which is basically
		singleton design pattern.
	*/
	once.Do(
		func() {
			// I wanted to print the query that gorm runs behind the scenes. This logger instance of gorm package help achive that.
			newLogger := logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      true,
				},
			)

			// I am getting the values which are previously loaded by the loadConfigs() function.
			dbUser := os.Getenv("DB_USER")
			dbPassword := os.Getenv("DB_PASSWORD")
			dbHost := os.Getenv("DB_HOST")
			dbPort := os.Getenv("DB_PORT")
			dbName := os.Getenv("DB_NAME")

			// Making a connection to the local database
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
			if err != nil {
				log.Fatal("failed to connect to db, error:", err.Error())
			}

			// This starts the debug mode by printing the query that runs behind the scene.
			db = db.Debug()

			// This db instance is being used everywhere in the application
			handler.DB = db
		},
	)

	return handler.DB
}

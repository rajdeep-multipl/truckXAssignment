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

func (handler *SqlDB) ConnectDB() *gorm.DB {
	once.Do(
		func() {
			newLogger := logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second, // Slow SQL threshold
					LogLevel:      logger.Info, // Log level
					Colorful:      true,        // Enable color
				},
			)

			dbUser := os.Getenv("DB_USER")
			dbPassword := os.Getenv("DB_PASSWORD")
			dbHost := os.Getenv("DB_HOST")
			dbPort := os.Getenv("DB_PORT")
			dbName := os.Getenv("DB_NAME")

			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
			if err != nil {
				log.Fatal("failed to connect to db, error:", err.Error())
			}

			db = db.Debug()

			handler.DB = db
		},
	)

	return handler.DB
}

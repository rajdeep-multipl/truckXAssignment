package database

import "gorm.io/gorm"

type DB interface {
	ConnectDB() *gorm.DB
}

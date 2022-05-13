package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitiateConnection() {
	dsn := "host=localhost user=douglas password=docker dbname=go_grpc sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Error at database connection:", err)
	}
}

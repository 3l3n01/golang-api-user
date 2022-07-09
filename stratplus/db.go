package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDbConnection() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "usuario"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "db"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

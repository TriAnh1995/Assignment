package middleware

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		"test", "", "test")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, ServerError
	}
	return db, nil
}

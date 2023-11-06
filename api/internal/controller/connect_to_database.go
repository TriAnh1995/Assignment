package controller

import (
	"database/sql"
	"fmt"
)

func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=disable",
		"test", "",
		"@localhost:5432", "test"))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

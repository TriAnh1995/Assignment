package main

import (
	"assignment/cmd/routes"
	"assignment/internal/controller"
	"assignment/internal/handler"
	"assignment/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func main() {
	println("Hello world")

	os.Setenv("PORT", "3000")
	pgConn, err := ConnectToDatabase()
	if err != nil {
		log.Fatal("Error when connect DB", err)
	}
	// init repo, controller, handler
	repo := repository.New(pgConn)
	ctrl := controller.New(repo)
	rtr := routes.New(
		handler.New(ctrl),
	)

	r := gin.Default()
	rtr.Routes(r)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

var db *sql.DB
var gormDB *gorm.DB

func ConnectToDatabase() (*sql.DB, error) {
	if db == nil {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf(
			"postgresql://%s:%s@%s/%s?sslmode=disable",
			"test", "",
			"@localhost:5432", "test"))
		if err != nil {
			return nil, err
		}
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

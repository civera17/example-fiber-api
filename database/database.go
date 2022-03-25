package database

import (
	"log"
	"os"

	"github.com/civera17/fintech-test/models"
	"github.com/civera17/fintech-test/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// connectDb
func ConnectDb() {
	dsn := "host=localhost user=postgres password=postgres dbname=gp-db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.Book{}, &models.QueryInfo{})
	db.Use(&middleware.SqlInfoPlugin{})

	DB = Dbinstance{
		Db: db,
	}
}


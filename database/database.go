package database

import (
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/civera17/fintech-test/middleware"
	"github.com/civera17/fintech-test/models"
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

func ConnectMockDb() sqlmock.Sqlmock {
	sql, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Failed to connect to mock database. \n", err)
		os.Exit(2)
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sql,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector)

	DB = Dbinstance{
		Db: db,
	}

	return mock
}


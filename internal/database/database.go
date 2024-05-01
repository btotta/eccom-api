package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service interface {
	GetDB() *gorm.DB
}

type service struct {
	db *gorm.DB
}

var (
	database    = os.Getenv("DB_DATABASE")
	password    = os.Getenv("DB_PASSWORD")
	username    = os.Getenv("DB_USERNAME")
	port        = os.Getenv("DB_PORT")
	host        = os.Getenv("DB_HOST")
	queryLogger = os.Getenv("QUERY_LOGGER") == "true"
	dbInstance  *service
)

func New() Service {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, username, password, database, port)

	var logMode logger.LogLevel

	if queryLogger {
		logMode = logger.Info
	} else {
		logMode = logger.Silent
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get database: ", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	Migration(db)

	dbInstance = &service{db: db}

	return dbInstance

}

func (s *service) GetDB() *gorm.DB {
	if s.db == nil {
		New()
		return s.db
	}

	return s.db
}

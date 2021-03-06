package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/roryjarrard/fiber-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func getDSN() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	name := os.Getenv("DBNAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Denver", host, port, user, pass, name)
	return dsn, nil
}

func ConnectDB() error {
	dsn, err := getDSN()
	if err != nil {
		return err
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")

	err = runMigrations(db)
	if err != nil {
		return err
	}

	Database = DbInstance{
		Db: db,
	}
	return nil
}

func runMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)
	if err != nil {
		return err
	}
	return nil
}

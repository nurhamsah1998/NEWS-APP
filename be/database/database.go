package database

import (
	"log"
	"os"

	"github.com/nurhamsah1998/news/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func DBConnection() {
	dsn := "host=localhost user=postgres password=root dbname=news-app port=5000 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to the database", err.Error())
		os.Exit(2)
	}
	log.Println("Coneection to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migration")
	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.NewsCategory{}, &models.NewsPost{})
	Database = DbInstance{Db: db}
}

package database

import (
	"WebAPI1/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	dsn := "host=localhost port=5432 user=postgres password=2828 dbname=assignment sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	//log.Println("Running Migrations")
	db.AutoMigrate(&models.ReportOutput{})

	Database = DbInstance{Db: db}

}

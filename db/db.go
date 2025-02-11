package db

import (
	"fmt"

	"github.com/Tomoki108/burny/burny/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "host=localhost user=burny_user password=pass dbname=burny_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not open DB: %w", err)
	}

	if err := db.AutoMigrate(
		&models.Project{},
		&models.Sprint{},
		&models.SprintStat{},
	); err != nil {
		return fmt.Errorf("could not migrate DB: %w", err)
	}

	DB = db

	return nil
}

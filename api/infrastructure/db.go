package infrastructure

import (
	"fmt"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := "host=localhost user=burny_user password=pass dbname=burny_db port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not open DB: %w", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Sprint{},
	); err != nil {
		return fmt.Errorf("could not migrate DB: %w", err)
	}

	DB = db

	return nil
}

func NewTransactioner() domain.Transactioner {
	return Transactioner{
		DB: DB,
	}
}

type Transactioner struct {
	DB *gorm.DB
}

func (t Transactioner) Transaction(fn func(tx domain.Transaction) error) error {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

func (t Transactioner) Default() domain.Transaction {
	return t.DB
}

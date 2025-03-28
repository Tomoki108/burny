package infrastructure

import (
	"fmt"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		config.Conf.DB_HOST,
		config.Conf.DB_USER,
		config.Conf.DB_PASS,
		config.Conf.DB_NAME,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not open DB: %w", err)
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Project{},
		&model.Sprint{},
		&model.APIKey{},
	); err != nil {
		return fmt.Errorf("could not migrate DB: %w", err)
	}

	DB = db

	return nil
}

type Transactioner struct {
	DB *gorm.DB
}

func NewTransactioner() domain.Transactioner {
	return Transactioner{
		DB: DB,
	}
}

func (t Transactioner) Transaction(fn func(tx domain.Transaction) error) error {
	return t.DB.Transaction(func(tx *gorm.DB) error {
		return fn(tx)
	})
}

func (t Transactioner) Default() domain.Transaction {
	return t.DB
}

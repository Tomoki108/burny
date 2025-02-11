package models

import (
	"time"

	"gorm.io/gorm"
)

type Sprint struct {
	gorm.Model
	ProjectID int       `json:"project_id" gorm:"index"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	ActualSP  int       `json:"actual_sp"`
	IdealSP   int       `json:"ideal_sp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	TotalSP        int       `json:"total_sp"`
	SprintDuration int       `json:"sprint_duration"`
	Sprints        []*Sprint `json:"sprints"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

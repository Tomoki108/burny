package models

import (
	"time"

	"gorm.io/gorm"
)

type Sprint struct {
	gorm.Model
	ProjectID  uint        `json:"project_id" gorm:"index"`
	Name       string      `json:"name"`
	StartDate  time.Time   `json:"start_date"`
	EndDate    time.Time   `json:"end_date"`
	SprintStat *SprintStat `json:"sprint_stat" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

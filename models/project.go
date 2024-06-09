package models

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name               string    `json:"name"`
	StartDate          time.Time `json:"start_date"`
	EndDate            time.Time `json:"end_date"`
	InitialStoryPoints int       `json:"initial_story_points"`
	WeeksInSprint      int       `json:"weeks_in_sprint"`
	Sprints            []*Sprint `json:"sprints"`
}

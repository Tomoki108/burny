package models

import (
	"gorm.io/gorm"
)

type SprintStat struct {
	gorm.Model
	SprintID            uint `json:"sprint_id" gorm:"index"`
	Velocity            int  `json:"velocity"`
	RemaintgStoryPoints int  `json:"remaining_story_points"`
}

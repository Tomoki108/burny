package model

import (
	"time"

	"github.com/Tomoki108/burny/domain"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID         uint      `json:"user_id" gorm:"index"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TotalSP        int       `json:"total_sp"`
	StartDate      time.Time `json:"start_date"`
	SprintDuration int       `json:"sprint_duration"`
	SprintCount    int       `json:"sprint_count"`
}

func (p *Project) ToDomain() *domain.Project {
	return &domain.Project{
		ID:             p.ID,
		UserID:         p.UserID,
		Title:          p.Title,
		Description:    p.Description,
		TotalSP:        p.TotalSP,
		StartDate:      p.StartDate,
		SprintDuration: p.SprintDuration,
		SprintCount:    p.SprintCount,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func FromDomainProject(project *domain.Project) *Project {
	return &Project{
		UserID:         project.UserID,
		Title:          project.Title,
		Description:    project.Description,
		TotalSP:        project.TotalSP,
		StartDate:      project.StartDate,
		SprintCount:    project.SprintCount,
		SprintDuration: project.SprintDuration,
		Model: gorm.Model{
			CreatedAt: project.CreatedAt,
			UpdatedAt: project.UpdatedAt,
		},
	}
}

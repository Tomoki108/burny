package model

import (
	"time"

	"github.com/Tomoki108/burny/domain"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TotalSP        int       `json:"total_sp"`
	StartDate      time.Time `json:"start_date"`
	SprintDuration int       `json:"sprint_duration"`
	SprintCount    int       `json:"sprint_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *Project) ToDomain() *domain.Project {
	return &domain.Project{
		ID:             p.ID,
		Title:          p.Title,
		Description:    p.Description,
		TotalSP:        p.TotalSP,
		StartDate:      p.StartDate,
		SprintDuration: p.SprintDuration,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func FromDomainProject(project *domain.Project) *Project {
	return &Project{
		Title:          project.Title,
		Description:    project.Description,
		TotalSP:        project.TotalSP,
		StartDate:      project.StartDate,
		SprintCount:    project.SprintCount,
		SprintDuration: project.SprintDuration,
		CreatedAt:      project.CreatedAt,
		UpdatedAt:      project.UpdatedAt,
	}
}

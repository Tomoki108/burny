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
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	TotalSP        int       `json:"total_sp"`
	SprintDuration int       `json:"sprint_duration"`
	Sprints        []*Sprint `json:"sprints"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (p *Project) ToDomain() *domain.Project {
	return &domain.Project{
		ID:             p.ID,
		Title:          p.Title,
		Description:    p.Description,
		StartDate:      p.StartDate,
		EndDate:        p.EndDate,
		TotalSP:        p.TotalSP,
		SprintDuration: p.SprintDuration,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
	}
}

func FromDomainProject(project *domain.Project) *Project {
	return &Project{
		Title:          project.Title,
		Description:    project.Description,
		StartDate:      project.StartDate,
		EndDate:        project.EndDate,
		TotalSP:        project.TotalSP,
		SprintDuration: project.SprintDuration,
	}
}

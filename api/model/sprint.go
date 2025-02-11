package model

import (
	"time"

	"github.com/Tomoki108/burny/domain"
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

func (s *Sprint) ToDomain() *domain.Sprint {
	return &domain.Sprint{
		ID:        s.ID,
		ProjectID: s.ProjectID,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		ActualSP:  s.ActualSP,
		IdealSP:   s.IdealSP,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func FromDomainSprint(sprint *domain.Sprint) *Sprint {
	return &Sprint{
		ProjectID: sprint.ProjectID,
		StartDate: sprint.StartDate,
		EndDate:   sprint.EndDate,
		ActualSP:  sprint.ActualSP,
		IdealSP:   sprint.IdealSP,
	}
}

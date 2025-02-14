package model

import (
	"time"

	"github.com/Tomoki108/burny/domain"
	"gorm.io/gorm"
)

type Sprint struct {
	gorm.Model
	UserID    uint      `json:"user_id" gorm:"index"`
	ProjectID uint      `json:"project_id" gorm:"index"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	ActualSP  int       `json:"actual_sp"`
	IdealSP   int       `json:"ideal_sp"`
}

func (s *Sprint) ToDomain() *domain.Sprint {
	return &domain.Sprint{
		ID:        s.ID,
		UserID:    s.UserID,
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
		Model: gorm.Model{
			ID:        sprint.ID,
			CreatedAt: sprint.CreatedAt,
			UpdatedAt: sprint.UpdatedAt,
		},
		UserID:    sprint.UserID,
		ProjectID: sprint.ProjectID,
		StartDate: sprint.StartDate,
		EndDate:   sprint.EndDate,
		ActualSP:  sprint.ActualSP,
		IdealSP:   sprint.IdealSP,
	}
}

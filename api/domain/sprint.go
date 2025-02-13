package domain

import "time"

type Sprint struct {
	ID        uint      `json:"id"`
	ProjectID uint      `json:"project_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	ActualSP  int       `json:"actual_sp"`
	IdealSP   int       `json:"ideal_sp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SprintRepository interface {
	Create(tx Transaction, sprint *Sprint) (*Sprint, error)
	List() ([]*Sprint, error)
	Update(sprint *Sprint) (*Sprint, error)
}

package domain

import "time"

type Sprint struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ProjectID uint      `json:"project_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	ActualSP  int       `json:"actual_sp"`
	IdealSP   int       `json:"ideal_sp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SprintRepository interface {
	Create(tx Transaction, sprint *Sprint) error
	Get(tx Transaction, projectID, sprintID uint) (*Sprint, error)
	List(tx Transaction, pojectID uint) ([]*Sprint, error)
	Update(tx Transaction, projectID, sprintID uint, actualSP int) (*Sprint, error)
	Delete(tx Transaction, projectID, sprintID uint) error
}

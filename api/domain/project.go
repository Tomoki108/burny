package domain

import "time"

type Project struct {
	ID             uint      `json:"project_id"`
	UserID         uint      `json:"user_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TotalSP        int       `json:"total_sp"`
	StartDate      time.Time `json:"start_date"`
	SprintDuration int       `json:"sprint_duration"`
	SprintCount    int       `json:"sprint_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProjectRepository interface {
	List(tx Transaction, userID uint) ([]*Project, error)
	Create(tx Transaction, project *Project) (*Project, error)
	Get(tx Transaction, userID, id uint) (*Project, error)
	Update(tx Transaction, project *Project) (*Project, error)
	Delete(tx Transaction, userID, id uint) error
}

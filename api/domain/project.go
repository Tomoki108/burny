package domain

import "time"

type Project struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	TotalSP        int       `json:"total_sp"`
	SprintDuration int       `json:"sprint_duration"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProjectRepository interface {
	List() ([]*Project, error)
	Create(project *Project) (*Project, error)
	Get(id uint) (*Project, error)
	Update(project *Project) (*Project, error)
	Delete(id uint) error
}

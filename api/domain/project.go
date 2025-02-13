package domain

import "time"

type Project struct {
	ID             uint
	Title          string
	Description    string
	TotalSP        int
	StartDate      time.Time
	SprintDuration int
	SprintCount    int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type ProjectRepository interface {
	List() ([]*Project, error)
	Create(tx Transaction, project *Project) (*Project, error)
	Get(tx Transaction, id uint) (*Project, error)
	Update(tx Transaction, project *Project) (*Project, error)
	Delete(tx Transaction, id uint) error
}

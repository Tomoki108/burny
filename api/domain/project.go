package domain

import (
	"encoding/json"
	"time"
)

type Project struct {
	ID             uint      `json:"id"`
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

func (p Project) MarshalJSON() ([]byte, error) {
	type Alias Project // 再帰呼び出しを避けるためのエイリアス
	aux := &struct {
		StartDate string `json:"start_date"`
		*Alias
	}{
		Alias:     (*Alias)(&p),
		StartDate: p.StartDate.Format("2006-01-02"),
	}
	return json.Marshal(aux)
}

func (p *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	aux := &struct {
		StartDate string `json:"start_date"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", aux.StartDate)
	if err != nil {
		return err
	}
	p.StartDate = t
	return nil
}

type ProjectRepository interface {
	List(tx Transaction, userID uint) ([]*Project, error)
	Create(tx Transaction, project *Project) (*Project, error)
	Get(tx Transaction, userID, id uint) (*Project, error)
	Update(tx Transaction, project *Project) (*Project, error)
	Delete(tx Transaction, userID, id uint) error
}

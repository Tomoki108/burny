package domain

import (
	"encoding/json"
	"time"
)

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

func (s Sprint) MarshalJSON() ([]byte, error) {
	type Alias Sprint // 再帰呼び出しを避けるためのエイリアス
	aux := &struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		*Alias
	}{
		Alias:     (*Alias)(&s),
		StartDate: s.StartDate.Format("2006-01-02"),
		EndDate:   s.EndDate.Format("2006-01-02"),
	}
	return json.Marshal(aux)
}

func (s *Sprint) UnmarshalJSON(data []byte) error {
	type Alias Sprint
	aux := &struct {
		StartDate string `json:"start_date"`
		EndDate   string `json:"end_date"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02", aux.StartDate)
	if err != nil {
		return err
	}
	s.StartDate = t
	return nil
}

type SprintRepository interface {
	Create(tx Transaction, sprint *Sprint) error
	Get(tx Transaction, projectID, sprintID uint) (*Sprint, error)
	List(tx Transaction, pojectID uint) ([]*Sprint, error)
	Update(tx Transaction, projectID, sprintID uint, actualSP int) (*Sprint, error)
	Delete(tx Transaction, projectID, sprintID uint) error
	UpsertList(tx Transaction, sprints []*Sprint) ([]*Sprint, error)
}

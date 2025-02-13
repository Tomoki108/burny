package io

import "time"

type GetProjectRequest struct {
	ID uint `json:"-" param:"project_id"`
}

type CreateProjectRequest struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TotalSP        int       `json:"total_sp"`
	StartDate      time.Time `json:"start_date"`
	SprintDuration int       `json:"sprint_duration"`
	SprintCount    int       `json:"sprint_count"`
}

type UpdateProjectRequest struct {
	ID          uint   `json:"-" param:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TotalSP     int    `json:"total_sp"`
	SprintCount int    `json:"sprint_count"`
}

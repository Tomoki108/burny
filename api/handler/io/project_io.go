package io

import "time"

type ProjectRequestBase struct {
	ProjectID uint `json:"-" param:"project_id"`
}

type GetProjectRequest ProjectRequestBase

type CreateProjectRequest struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	TotalSP        int       `json:"total_sp"`
	StartDate      time.Time `json:"start_date"`
	SprintDuration int       `json:"sprint_duration"`
	SprintCount    int       `json:"sprint_count"`
}

type UpdateProjectRequest struct {
	ProjectRequestBase
	Title       string `json:"title"`
	Description string `json:"description"`
	TotalSP     int    `json:"total_sp"`
	SprintCount int    `json:"sprint_count"`
}

type DeleteProjectRequest ProjectRequestBase

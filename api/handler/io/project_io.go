package io

import "time"

type ProjectRequestBase struct {
	ProjectID uint `json:"-" param:"project_id"`
}

type GetProjectRequest ProjectRequestBase

type CreateProjectRequest struct {
	Title          string    `json:"title" validate:"required,max=100"`
	Description    string    `json:"description" validate:"max=500"`
	TotalSP        int       `json:"total_sp" validate:"required,max=1000"`
	StartDate      time.Time `json:"start_date" validate:"required"`
	SprintDuration int       `json:"sprint_duration" validate:"required,oneof=1 2 3"`
	SprintCount    int       `json:"sprint_count" validate:"required,min=1,max=100"`
}

type UpdateProjectRequest struct {
	ProjectRequestBase
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description" validate:"max=500"`
	TotalSP     int    `json:"total_sp" validate:"required,max=1000"`
	SprintCount int    `json:"sprint_count" validate:"required,min=1,max=100"`
}

type DeleteProjectRequest ProjectRequestBase

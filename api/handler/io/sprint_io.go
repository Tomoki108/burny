package io

type ListSprintRequest struct {
	ProjectID uint `json:"-" param:"project_id"`
}

type UpdateSprintRequest struct {
	ID        uint `json:"-" param:"sprint_id"`
	ProjectID uint `json:"-" param:"project_id"`
	ActualSP  int  `json:"actual_sp"`
}

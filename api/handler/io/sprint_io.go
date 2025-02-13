package io

type ProjectSprintRequestBase struct {
	ProjectID uint `json:"-" param:"project_id"`
}

type ListSprintRequest ProjectSprintRequestBase

type UpdateSprintRequest struct {
	ProjectSprintRequestBase
	ID       uint `json:"-" param:"sprint_id"`
	ActualSP int  `json:"actual_sp"`
}

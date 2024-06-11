package handlers

import (
	"burny-api/db"
	"burny-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListProjectsHandler(c echo.Context) error {
	var projects []models.Project
	if err := db.DB.Preload("Sprints").Find(&projects).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, projects)
}

func CreateProjectHandler(c echo.Context) error {
	project := new(models.Project)
	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Create(project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, project)
}

func GetProjectHandler(c echo.Context) error {
	id := c.Param("id")
	var project models.Project
	if err := db.DB.Preload("Sprints").First(&project, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, project)
}

func UpdateProjectHandler(c echo.Context) error {
	id := c.Param("id")
	var project models.Project
	if err := db.DB.First(&project, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Save(&project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, project)
}

func DeleteProjectHandler(c echo.Context) error {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Project{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

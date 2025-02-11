package handlers

import (
	"net/http"

	"github.com/Tomoki108/burny/db"
	"github.com/Tomoki108/burny/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateSprintHandler(c echo.Context) error {
	sprint := new(models.Sprint)
	if err := c.Bind(sprint); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := db.DB.Create(sprint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, sprint)
}

func GetSprintHandler(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	id := c.Param("id")
	var sprint models.Sprint
	if err := db.Preload("SprintStat").First(&sprint, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, sprint)
}

func UpdateSprintHandler(c echo.Context) error {
	id := c.Param("id")
	var sprint models.Sprint
	if err := db.DB.First(&sprint, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(&sprint); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := db.DB.Save(&sprint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sprint)
}

func DeleteSprintHandler(c echo.Context) error {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Sprint{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func ListSprintsHandler(c echo.Context) error {
	var sprints []models.Sprint
	if err := db.DB.Preload("SprintStat").Find(&sprints).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sprints)
}

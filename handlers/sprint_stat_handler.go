package handlers

import (
	"burny-api/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SprintStat struct {
	gorm.Model
	SprintID             uint `json:"sprint_id" gorm:"index"`
	Velocity             int  `json:"velocity"`
	RemainingStoryPoints int  `json:"remaining_story_points"`
}

func CreateSprintStatHandler(c echo.Context) error {
	sprintStat := new(SprintStat)
	if err := c.Bind(sprintStat); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := db.DB.Create(sprintStat).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, sprintStat)
}

func GetSprintStatHandler(c echo.Context) error {
	id := c.Param("id")
	var sprintStat SprintStat
	if err := db.DB.First(&sprintStat, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, sprintStat)
}

func UpdateSprintStatHandler(c echo.Context) error {
	id := c.Param("id")
	var sprintStat SprintStat
	if err := db.DB.First(&sprintStat, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	if err := c.Bind(&sprintStat); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := db.DB.Save(&sprintStat).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sprintStat)
}

func DeleteSprintStatHandler(c echo.Context) error {
	id := c.Param("id")
	if err := db.DB.Delete(&SprintStat{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func ListSprintStatsHandler(c echo.Context) error {
	var sprintStats []SprintStat
	if err := db.DB.Find(&sprintStats).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sprintStats)
}

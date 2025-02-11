package handler

import (
	"net/http"

	"github.com/Tomoki108/burny/db"
	"github.com/Tomoki108/burny/model"

	"github.com/labstack/echo/v4"
)

// @Summary      List projects
// @Description  List projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Project
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /projects [get]
func ListProjectsHandler(c echo.Context) error {
	var projects []model.Project
	if err := db.DB.Preload("Sprints").Find(&projects).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, projects)
}

// @Summary      Create projects
// @Description  Create projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}   model.Project
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /projects [post]
func CreateProjectHandler(c echo.Context) error {
	project := new(model.Project)
	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Create(project).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, project)
}

// @Summary      Get projects
// @Description  Get projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Project
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /projects/{id} [get]
func GetProjectHandler(c echo.Context) error {
	id := c.Param("id")
	var project model.Project
	if err := db.DB.Preload("Sprints").First(&project, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, project)
}

// @Summary      Update projects
// @Description  Update projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Project
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /projects/{id} [put]
func UpdateProjectHandler(c echo.Context) error {
	id := c.Param("id")
	var project model.Project
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

// @Summary      Delete projects
// @Description  Delete projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Project
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /projects/{id} [delete]
func DeleteProjectHandler(c echo.Context) error {
	id := c.Param("id")
	if err := db.DB.Delete(&model.Project{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	Usecase usecase.ProjectUseCase
}

// @Summary      List projects
// @Description  List projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.Project
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /projects [get]
func (h ProjectHandler) List(c echo.Context) error {
	projects, err := h.Usecase.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, projects)
}

// @Summary      Create a project
// @Description  Create a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param 	 	 request body io.CreateProjectRequest true "request"
// @Success      200  {object}  domain.Project
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /projects [post]
func (h ProjectHandler) Create(c echo.Context) error {
	req := new(io.CreateProjectRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	project := &domain.Project{
		Title:          req.Title,
		Description:    req.Description,
		TotalSP:        req.TotalSP,
		StartDate:      req.StartDate,
		SprintDuration: req.SprintDuration,
		SprintCount:    req.SprintCount,
	}
	created, err := h.Usecase.Create(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, created)
}

// @Summary      Get a project
// @Description  Get a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Project
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /projects/{id} [get]
func (h ProjectHandler) Get(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	project, err := h.Usecase.Get(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, project)
}

// @Summary      Update a project
// @Description  Update a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Project
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /projects/{id} [put]
func (h ProjectHandler) Update(c echo.Context) error {
	project := new(domain.Project)
	if err := c.Bind(project); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updated, err := h.Usecase.Update(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, updated)
}

// @Summary      Delete a projects
// @Description  Delete a projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Success      200  {object}  domain.Project
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /projects/{id} [delete]
func (h ProjectHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = h.Usecase.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

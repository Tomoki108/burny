package handler

import (
	"errors"
	"net/http"

	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	UseCase usecase.ProjectUseCase
}

func NewProjectHandler(usecase usecase.ProjectUseCase) ProjectHandler {
	return ProjectHandler{
		UseCase: usecase,
	}
}

// @Summary      List projects
// @Description  List projects
// @Tags         projects
// @Security	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Success      200 {array} domain.Project
// @Router       /projects [get]
func (h ProjectHandler) List(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	projects, err := h.UseCase.List(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, projects)
}

// @Summary      Create a project
// @Description  Create a project
// @Tags         projects
// @Security	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param 	 	 request body io.CreateProjectRequest true "request"
// @Success      201 {object} domain.Project
// @Failure      400 {object} io.ErrorResponse
// @Router       /projects [post]
func (h ProjectHandler) Create(c echo.Context) error {
	req := new(io.CreateProjectRequest)
	if err := handleReq(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := c.Get("user_id").(uint)
	created, err := h.UseCase.Create(userID, *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusCreated, created)
}

// @Summary      Get a project
// @Description  Get a project
// @Tags         projects
// @Security	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Success      200 {object} domain.Project
// @Failure      404 {object} io.ErrorResponse
// @Router       /projects/{project_id} [get]
func (h ProjectHandler) Get(c echo.Context) error {
	req := new(io.GetProjectRequest)
	if err := handleReq(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := c.Get("user_id").(uint)
	project, err := h.UseCase.Get(userID, req.ProjectID)
	if errors.Is(err, usecase.ErrProjectNotFound) {
		return c.JSON(http.StatusNotFound, io.NewErrResp(err.Error()))
	} else if err != nil {
		return c.JSON(http.StatusNotFound, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, project)
}

// @Summary      Update a project
// @Description  Update a project
// @Tags         projects
// @Security	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Param 	 	 request body io.UpdateProjectRequest true "request"
// @Success      200 {object} domain.Project
// @Failure      400 {object} io.ErrorResponse
// @Failure      404 {object} io.ErrorResponse
// @Router       /projects/{project_id} [put]
func (h ProjectHandler) Update(c echo.Context) error {
	req := new(io.UpdateProjectRequest)
	if err := handleReq(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := c.Get("user_id").(uint)
	updated, err := h.UseCase.Update(userID, *req)
	if errors.Is(err, usecase.ErrProjectNotFound) {
		return c.JSON(http.StatusNotFound, io.NewErrResp(err.Error()))
	} else if errors.Is(err, usecase.ErrSprintHasAlreadyStarted) {
		return c.JSON(http.StatusBadRequest, io.NewErrResp(err.Error()))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, updated)
}

// @Summary      Delete a projects
// @Description  Delete a projects
// @Tags         projects
// @Security	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Success      204
// @Failure      404 {object} io.ErrorResponse
// @Router       /projects/{project_id} [delete]
func (h ProjectHandler) Delete(c echo.Context) error {
	req := new(io.DeleteProjectRequest)
	if err := handleReq(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := c.Get("user_id").(uint)
	err := h.UseCase.Delete(userID, *req)
	if errors.Is(err, usecase.ErrProjectNotFound) {
		return c.JSON(http.StatusNotFound, io.NewErrResp(err.Error()))
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}

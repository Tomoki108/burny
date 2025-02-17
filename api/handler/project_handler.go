package handler

import (
	"net/http"

	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	UseCase usecase.ProjectUseCase
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
	if err := handleReq(c, req); err != nil {
		return err
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
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Success      200 {object}  domain.Project
// @Failure      400 {object} io.ErrorResponse
// @Failure      404 {object} io.ErrorResponse
// @Failure      500 {object} io.ErrorResponse
// @Router       /projects/{project_id} [get]
func (h ProjectHandler) Get(c echo.Context) error {
	req := new(io.GetProjectRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	userID := c.Get("user_id").(uint)
	project, err := h.UseCase.Get(userID, req.ProjectID)
	if err != nil {
		return c.JSON(http.StatusNotFound, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, project)
}

// @Summary      Update a project
// @Description  Update a project
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Param 	 	 request body io.UpdateProjectRequest true "request"
// @Success      200 {object}  domain.Project
// @Failure      400 {object} io.ErrorResponse
// @Failure      404 {object} io.ErrorResponse
// @Failure      500 {object} io.ErrorResponse
// @Router       /projects/{project_id} [put]
func (h ProjectHandler) Update(c echo.Context) error {
	req := new(io.UpdateProjectRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	userID := c.Get("user_id").(uint)
	updated, err := h.UseCase.Update(userID, *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, updated)
}

// @Summary      Delete a projects
// @Description  Delete a projects
// @Tags         projects
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Success      200 {object}  domain.Project
// @Failure      400 {object} io.ErrorResponse
// @Failure      404 {object} io.ErrorResponse
// @Failure      500 {object} io.ErrorResponse
// @Router       /projects/{project_id} [delete]
func (h ProjectHandler) Delete(c echo.Context) error {
	req := new(io.DeleteProjectRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	userID := c.Get("user_id").(uint)
	err := h.UseCase.Delete(userID, *req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}

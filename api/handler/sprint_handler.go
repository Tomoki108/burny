package handler

import (
	"net/http"

	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"

	"github.com/labstack/echo/v4"
)

type SprintHandler struct {
	UseCase usecase.SprintUseCase
}

// @Summary      List sprints
// @Description  List sprints
// @Tags         sprints
// @Accept       json
// @Produce      json
// @Param 	 	 project_id path int true "project_id"
// @Success      200  {array}  domain.Sprint
// @Failure      404
// @Failure      500
// @Router       /projects/{project_id}/sprints [get]
func (h SprintHandler) List(c echo.Context) error {
	req := new(io.ListSprintRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	sprints, err := h.UseCase.List(req.ProjectID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, sprints)
}

// @Summary      Update a sprint
// @Description  Update a sprint
// @Tags         sprints
// @Accept       json
// @Produce      json
// @Success      200  {array}  domain.Sprint
// @Param 	 	 project_id path int true "project_id"
// @Param 	 	 sprint_id path int true "sprint_id"
// @Param 	 	 request body io.UpdateSprintRequest true "request"
// @Failure      404
// @Failure      500
// @Router       /projects/{project_id}/sprints/{sprint_id} [patch]
func (h SprintHandler) Update(c echo.Context) error {
	req := new(io.UpdateSprintRequest)
	if err := handleReq(c, req); err != nil {
		return err
	}

	updated, err := h.UseCase.Update(req.ProjectID, req.ID, req.ActualSP)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, updated)
}

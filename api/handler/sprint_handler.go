package handler

import (
	"net/http"

	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/model"

	"github.com/labstack/echo/v4"
)

type SprintHandler struct {
	Repo domain.SprintRepository
}

func (h SprintHandler) List(c echo.Context) error {
	sprints, err := h.Repo.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, sprints)
}

func (h SprintHandler) Update(c echo.Context) error {
	sprint := new(model.Sprint)
	if err := c.Bind(sprint); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	updated, err := h.Repo.Update(sprint.ToDomain())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, updated)
}

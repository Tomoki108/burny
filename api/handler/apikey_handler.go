package handler

import (
	"net/http"

	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/usecase"
	"github.com/labstack/echo/v4"
)

type APIKeyHandler struct {
	Usecase usecase.APIKeyUseCase
}

func NewAPIKeyHandler(usecase usecase.APIKeyUseCase) APIKeyHandler {
	return APIKeyHandler{
		Usecase: usecase,
	}
}

// @Summary      Create API Key
// @Description  Create API Key
// @Tags         API Key
// @Produce      json
// @Success      200 {object} io.CreateAPIKeyResponse
// @Router       /apikey [post]
func (h APIKeyHandler) Create(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	apiKey, err := h.Usecase.Create(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusCreated, apiKey)
}

func (h APIKeyHandler) Get(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	apiKeys, err := h.Usecase.Get(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, apiKeys)
}

// @Summary      Delete API Key
// @Description  Delete API Key
// @Tags         API Key
// @Produce      json
// @Success      204
// @Failure      404 {object} io.ErrorResponse
// @Router       /apikey [delete]
func (h APIKeyHandler) Delete(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	err := h.Usecase.Delete(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}

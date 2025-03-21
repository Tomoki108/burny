package handler

import (
	"errors"
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
// @Security	 ApiKeyAuth
// @Produce      json
// @Success      201 {object} io.CreateAPIKeyResponse
// @Failure      409 {object} io.ErrorResponse
// @Router       /apikeys [post]
func (h APIKeyHandler) Create(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	apiKeyResp, err := h.Usecase.Create(userID)
	if err != nil {
		if errors.Is(err, usecase.ErrAPIKeyAlreadyExists) {
			return c.JSON(http.StatusConflict, io.NewErrResp(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusCreated, apiKeyResp)
}

// @Summary      Check API Key Status
// @Description  Check if the user has an API key
// @Tags         API Key
// @Security	 ApiKeyAuth
// @Produce      json
// @Success      200 {object} io.APIKeyStatusResponse
// @Router       /apikeys/status [get]
func (h APIKeyHandler) CheckStatus(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	exists, err := h.Usecase.CheckStatus(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.JSON(http.StatusOK, io.APIKeyStatusResponse{Exists: exists})
}

// @Summary      Delete API Key
// @Description  Delete API Key
// @Tags         API Key
// @Security	 ApiKeyAuth
// @Produce      json
// @Success      204
// @Failure      404 {object} io.ErrorResponse
// @Router       /apikeys [delete]
func (h APIKeyHandler) Delete(c echo.Context) error {
	userID := c.Get("user_id").(uint)
	err := h.Usecase.Delete(userID)
	if err != nil {
		if errors.Is(err, usecase.ErrAPIKeyNotFound) {
			return c.JSON(http.StatusNotFound, io.NewErrResp(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, io.NewErrResp(err.Error()))
	}

	return c.NoContent(http.StatusNoContent)
}

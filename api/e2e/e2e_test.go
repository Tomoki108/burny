package e2e

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/di"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/sign_up", strings.NewReader(`{"email":"test@test.com","password":"passwd12345"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	c := e.NewContext(req, recorder)

	if err := config.Init(); err != nil {
		log.Fatal(err.Error())
	}

	// DB接続
	if err := infrastructure.ConnectDB(); err != nil {
		log.Fatal(err.Error())
	}

	di.ProvideDependencies()
	var authH handler.AuthHandler
	di.Container.Invoke(func(h handler.AuthHandler) {
		authH = h
	})

	if assert.NoError(t, authH.SignUp(c)) {
		// レスポンスの検証
		assert.Equal(t, http.StatusOK, recorder.Code)
		// その他のアサーション（例: レスポンスボディの検証）を実施
	}
}

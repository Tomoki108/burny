package e2e

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/di"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var e *echo.Echo
var testTx *gorm.DB

func init() {
	// 環境変数の読み込み
	if err := config.Init(); err != nil {
		log.Fatal(err.Error())
	}
	// DB接続
	if err := infrastructure.ConnectDB(); err != nil {
		log.Fatal(err.Error())
	}
	// DIコンテナの初期化
	di.ProvideDependencies()
	// テスト用トランザクションの初期化
	testTx = infrastructure.DB.Begin()
	di.Container.Decorate(func(transactioner domain.Transactioner) domain.Transactioner {
		return infrastructure.NewTestTransactioner(testTx)
	})
	// Echoインスタンス生成
	e = echo.New()
}

func TestE2E(t *testing.T) {
	defer testTx.Rollback()

	UserCanSignUp(t)
}

func UserCanSignUp(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/sign_up", strings.NewReader(`{"email":"test@test.com","password":"passwd12345"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(req, recorder)

	var authH handler.AuthHandler
	di.Container.Invoke(func(h handler.AuthHandler) {
		authH = h
	})

	if assert.NoError(t, authH.SignUp(c)) {
		assert.Equal(t, http.StatusCreated, recorder.Code)
		body, err := removeDynamicFields(recorder.Body.Bytes(), "password")
		if err != nil {
			t.Fatal(err)
		}
		g := goldie.New(t)
		g.Assert(t, "signup_response", body)
	}
}

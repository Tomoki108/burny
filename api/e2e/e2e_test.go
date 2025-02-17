package e2e

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/di"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler"
	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/sebdah/goldie/v2"
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
	token := UserCanSignIn(t)
}

func UserCanSignUp(t *testing.T) {
	reqBody := strings.NewReader(`{"email":"test@test.com","password":"passwd12345"}`)
	req := httptest.NewRequest(http.MethodPost, "/sign_up", reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(req, recorder)

	var authH handler.AuthHandler
	di.Container.Invoke(func(h handler.AuthHandler) {
		authH = h
	})

	if err := authH.SignUp(c); err != nil {
		t.Fatal(err)
	}
	if http.StatusCreated != recorder.Code {
		t.Fatal("status code is not 201")
	}
	body, err := removeDynamicFields(recorder.Body.Bytes(), "password")
	if err != nil {
		t.Fatal(err)
	}
	g := goldie.New(t)
	g.Assert(t, "signup_response", body)
}

func UserCanSignIn(t *testing.T) (token string) {
	reqBody := strings.NewReader(`{"email":"test@test.com","password":"passwd12345"}`)
	req := httptest.NewRequest(http.MethodPost, "/sign_int", reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	recorder := httptest.NewRecorder()
	c := e.NewContext(req, recorder)

	var authH handler.AuthHandler
	di.Container.Invoke(func(h handler.AuthHandler) {
		authH = h
	})

	if err := authH.SignIn(c); err != nil {
		t.Fatal(err)
	}
	if http.StatusOK != recorder.Code {
		t.Fatal("status code is not 200")
	}

	bodyBytes := recorder.Body.Bytes()
	res := io.SignInResponse{}
	err := json.Unmarshal(bodyBytes, &res)
	if err != nil {
		t.Fatal(err)
	}
	return res.JwtToken
}

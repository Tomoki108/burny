package scenario

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Tomoki108/burny/config"
	"github.com/Tomoki108/burny/domain"
	"github.com/Tomoki108/burny/handler/io"
	"github.com/Tomoki108/burny/infrastructure"
	"github.com/Tomoki108/burny/server"
	"github.com/labstack/echo/v4"
	"github.com/sebdah/goldie/v2"
	"gorm.io/gorm"
)

var e *echo.Echo
var testTx *gorm.DB
var mailerMock *infrastructure.AWSSESMailerMock

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
	server.InitDIContainer()
	// テスト用トランザクションの初期化
	testTx = infrastructure.DB.Begin()
	server.Container.Decorate(func(transactioner domain.Transactioner) domain.Transactioner {
		return infrastructure.Transactioner{DB: testTx}
	})
	// テスト用メール送信モックの初期化
	mailerMock = infrastructure.NewAWSSESMailerMock()
	server.Container.Decorate(func(m domain.Mailer) domain.Mailer {
		return mailerMock
	})

	// サーバーの取得
	e = server.NewEchoServer()
}

func TestScenario(t *testing.T) {
	defer testTx.Rollback()

	// Authentication
	UserCanSignUp(t)
	jwtToken := UserCanSignIn(t)

	// Project Operations
	projectID := UserCanCreateProject(t, jwtToken)
	UserCanListProjects(t, jwtToken)
	UserCanGetProject(t, jwtToken, projectID)
	UserCanUpdateProject(t, jwtToken, projectID)
	defer UserCanDeleteProject(t, jwtToken, projectID)

	// API Key Operations
	UserCanCheckAPIKeyStatus(t, jwtToken, false)
	apiKey := UserCanCreateAPIKey(t, jwtToken)
	UserCanCheckAPIKeyStatus(t, jwtToken, true)
	defer UserCanDeleteAPIKey(t, jwtToken)

	// Sprint Operations
	sprintID := UserCanListSprints(t, apiKey, projectID)
	UserCanUpdateSprint(t, apiKey, projectID, sprintID)
}

func UserCanSignUp(t *testing.T) string {
	// Arrange
	reqBody := strings.NewReader(`{"email":"test@test.com","password":"passwd12345"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/sign_up", reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusCreated, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes(), "password")
	if err != nil {
		t.Fatal(err)
	}
	g := goldie.New(t)
	g.Assert(t, "signup_response", body)

	sentMail := mailerMock.SentMails[0]
	if sentMail.To != "test@test.com" {
		t.Fatalf("Expected mail to be test@test.com but got %s", sentMail.To)
	}
	if sentMail.Subject != "Burny Email Verification" {
		t.Fatalf("Expected mail subject to be Burny Email Verification but got %s", sentMail.Subject)
	}

	// sentMail.Bodyからメールアドレス検証URLを抽出して返す

}

func UserCanSignIn(t *testing.T) (token string) {
	// Arrange
	reqBody := strings.NewReader(`{"email":"test@test.com","password":"passwd12345"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/sign_in", reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}

	bodyBytes := recorder.Body.Bytes()
	res := io.SignInResponse{}
	err := json.Unmarshal(bodyBytes, &res)
	if err != nil {
		t.Fatal(err)
	}
	return res.JwtToken
}

func UserCanCreateProject(t *testing.T, token string) (projectID uint) {
	// Arrange
	reqBody := strings.NewReader(`{
	  	"title": "test project",
  		"description": "this is test project",
		"start_date": "2100-04-01T09:00:00+09:00",
  		"sprint_count": 100,
  		"sprint_duration": 1,
  		"total_sp": 1000
	}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusCreated, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	g := goldie.New(t)
	g.Assert(t, "create_project_response", body)

	res := domain.Project{}
	err = json.Unmarshal(recorder.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}

	return res.ID
}

func UserCanListProjects(t *testing.T, token string) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	goldie.New(t).Assert(t, "list_projects_response", body)
}

func UserCanGetProject(t *testing.T, token string, projectID uint) {
	// Arrange
	url := "/api/v1/projects/" + uintToStr(projectID)
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	goldie.New(t).Assert(t, "get_project_response", body)
}

func UserCanUpdateProject(t *testing.T, token string, projectID uint) {
	// Arrange
	updateJSON := `{
		"title": "updated project",
		"description": "updated description",
		"start_date": "2100-04-02T09:00:00+09:00",
		"sprint_count": 40,
		"sprint_duration": 2,
		"total_sp": 800
	}`
	url := "/api/v1/projects/" + uintToStr(projectID)
	reqBody := strings.NewReader(updateJSON)
	req := httptest.NewRequest(http.MethodPut, url, reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	goldie.New(t).Assert(t, "update_project_response", body)
}

func UserCanDeleteProject(t *testing.T, token string, projectID uint) {
	// Arrange
	url := "/api/v1/projects/" + uintToStr(projectID)
	req := httptest.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusNoContent, recorder); err != nil {
		t.Fatal(err)
	}
}

func UserCanCheckAPIKeyStatus(t *testing.T, token string, expected bool) {
	// Arrange
	req := httptest.NewRequest(http.MethodGet, "/api/v1/apikeys/status", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}
	var response struct {
		Exists bool `json:"exists"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.Exists != expected {
		t.Fatalf("Expected API key status to be %v, but got %v", expected, response.Exists)
	}
}

func UserCanCreateAPIKey(t *testing.T, token string) (rawKey string) {
	// Arrange
	req := httptest.NewRequest(http.MethodPost, "/api/v1/apikeys", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusCreated, recorder); err != nil {
		t.Fatal(err)
	}

	var response struct {
		RawKey string `json:"raw_key"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.RawKey == "" {
		t.Fatal("API key should not be empty")
	}
	t.Logf("API Key created successfully: %s", response.RawKey[:5]+"...")

	return response.RawKey
}

func UserCanDeleteAPIKey(t *testing.T, token string) {
	// Arrange
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/apikeys", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusNoContent, recorder); err != nil {
		t.Fatal(err)
	}
	t.Log("API Key deleted successfully")
}

func UserCanListSprints(t *testing.T, apiKey string, projectID uint) (sprintID uint) {
	// Arrange
	url := "/api/v1/projects/" + uintToStr(projectID) + "/sprints"
	req := httptest.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("X-API-Key", apiKey)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	goldie.New(t).Assert(t, "list_sprints_response", body)

	res := []domain.Sprint{}
	err = json.Unmarshal(recorder.Body.Bytes(), &res)
	if err != nil {
		t.Fatal(err)
	}
	return res[0].ID
}

func UserCanUpdateSprint(t *testing.T, apiKey string, projectID, sprintID uint) {
	// Arrange
	url := "/api/v1/projects/" + uintToStr(projectID) + "/sprints/" + uintToStr(sprintID)
	updateJSON := `{
		"actual_sp": 100
	}`
	reqBody := strings.NewReader(updateJSON)
	req := httptest.NewRequest(http.MethodPatch, url, reqBody)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("X-API-Key", apiKey)
	recorder := httptest.NewRecorder()

	// Act
	e.ServeHTTP(recorder, req)

	// Assert
	if err := assertSatusCode(http.StatusOK, recorder); err != nil {
		t.Fatal(err)
	}
	body, err := removeDynamicFields(recorder.Body.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	goldie.New(t).Assert(t, "update_sprint_response", body)
}

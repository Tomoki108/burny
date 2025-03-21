package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Tomoki108/burny/domain"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type APIKeyAuthMiddleware struct {
	repo          domain.APIKeyRepository
	transactioner domain.Transactioner
}

func NewAPIKeyAuthMiddleware(repo domain.APIKeyRepository, transactioner domain.Transactioner) *APIKeyAuthMiddleware {
	return &APIKeyAuthMiddleware{
		repo:          repo,
		transactioner: transactioner,
	}
}

// NOTE: make sure APIKeyAuthMiddleware is called after JWTAuthMiddleware
func (m *APIKeyAuthMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get("auth_method") == "jwt" {
				return next(c)
			}

			log.Default().Print("hello")

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "ApiKey" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header")
			}

			apiKey := parts[1]

			// すべてのAPIKeyを取得して検証
			// 注: パフォーマンスを考慮すると、キーのハッシュをインデックス化するなど
			// より効率的な方法を検討する必要がある
			allKeys, err := m.repo.GetAll(m.transactioner.Default())
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to validate ApiKey")
			}

			for _, key := range allKeys {
				if err := bcrypt.CompareHashAndPassword([]byte(key.Key), []byte(apiKey)); err == nil {
					c.Set("user_id", key.UserID)
					c.Set("auth_method", "apikey")
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid ApiKey")
		}
	}
}

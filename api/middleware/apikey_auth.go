package middleware

import (
	"net/http"
	"strings"

	"github.com/Tomoki108/burny/domain"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type APIKeyAuthMiddleware struct {
	repo domain.APIKeyRepository
}

func NewAPIKeyAuthMiddleware(repo domain.APIKeyRepository) *APIKeyAuthMiddleware {
	return &APIKeyAuthMiddleware{
		repo: repo,
	}
}

// NOTE: APIKeyAuthMiddleware must be used before JWTAuthMiddleware
func (m *APIKeyAuthMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "ApiKey" {
				return next(c)
			}

			apiKey := parts[1]

			// すべてのAPIKeyを取得して検証
			// 注: パフォーマンスを考慮すると、キーのハッシュをインデックス化するなど
			// より効率的な方法を検討する必要がある
			allKeys, err := m.repo.GetAll(nil)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to validate API key")
			}

			for _, key := range allKeys {
				// ハッシュと平文のキーを比較
				if err := bcrypt.CompareHashAndPassword([]byte(key.Key), []byte(apiKey)); err == nil {
					// 認証成功
					c.Set("user_id", key.UserID)
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API key")
		}
	}
}

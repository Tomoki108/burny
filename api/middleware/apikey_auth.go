package middleware

import (
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

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization Header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "ApiKey" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization Header")
			}

			apiKeyRaw := parts[1]

			// Get all API keys and find matching one
			// We need to check all because bcrypt hash can't be used for direct DB lookup
			apiKeys, err := m.repo.GetAll(m.transactioner.Default())
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "Failed to validate ApiKey")
			}

			var authenticated *domain.APIKey
			for _, key := range apiKeys {
				if err := bcrypt.CompareHashAndPassword([]byte(key.Key), []byte(apiKeyRaw)); err == nil {
					authenticated = key
					break
				}
			}

			if authenticated == nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid ApiKey")
			}

			c.Set("user_id", authenticated.UserID)
			c.Set("auth_method", "apikey")
			return next(c)
		}
	}
}

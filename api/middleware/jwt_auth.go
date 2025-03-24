package middleware

import (
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JWTAuthMiddleware struct {
	secret []byte
}

func NewJWTAuthMiddleware(secret []byte) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		secret: secret,
	}
}

// NOTE: make sure JWTAuthMiddleware is called before APIKeyAuthMiddleware
func (m *JWTAuthMiddleware) Middleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: m.secret,
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			// whether user_id is string or float64 is not guaranteed
			if userID, ok := claims["user_id"].(string); ok {
				userIDInt, _ := strconv.Atoi(userID)
				c.Set("user_id", uint(userIDInt))
			} else if userIDFloat, ok := claims["user_id"].(float64); ok {
				c.Set("user_id", uint(userIDFloat))
			}
			c.Set("auth_method", "jwt")
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return nil // ignore error and let the api key middleware handle request
		},
		ContinueOnIgnoredError: true,
	})
}

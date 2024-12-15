package controller

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"smart_school_for_mirea/internal/jwt"
)

const (
	authorizationHeader = "Authorization"
	emailCtx            = "email"
	userIDCtx           = "userId"
)

func (h *handlers) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if len(c.Path()) > 5 && c.Path()[:5] == "/auth" {
			return next(c)
		}

		log := h.logger.With(slog.String("op", "AuthMiddleware"))

		jwtToken := c.Request().Header.Get(authorizationHeader)
		if jwtToken == "" {
			log.Debug("no authorization header")

			return echo.NewHTTPError(http.StatusUnauthorized, "No authorization header")
		}

		claims, err := jwt.ParseToken(jwtToken, []byte(h.authSecretKey))
		if err != nil {
			log.Debug("failed to parse token: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect token")
		}

		id, ok := claims["id"].(string)
		if !ok {
			log.Debug("no id in claims: %+v", claims)
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect token")
		}
		userId, err := uuid.Parse(id)
		if err != nil {
			log.Debug("failed to parse user id: %v", err)
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect token")
		}

		email, ok := claims["email"].(string)
		if !ok {
			log.Warn("no email in claims: %+v", claims)
			return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect token")
		}

		log.Info("email", email)
		log.Info("id", userId)

		c.Set(emailCtx, email)
		c.Set(userIDCtx, userId)

		return next(c)
	}
}

package middlewares

import (
	"errors"
	"strings"

	"github.com/RhoNit/budgetapi/common"
	"github.com/RhoNit/budgetapi/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppMiddleware struct {
	Logger echo.Logger
	DB     *gorm.DB
}

func (appMiddleware *AppMiddleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Add("Vary", "Authorization")

		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return common.SendUnauthorizedResponse(c, "Please provide a Bearer token")
		}

		authHeaderSplit := strings.Split(authHeader, " ")
		if len(authHeaderSplit) != 2 {
			return common.SendUnauthorizedResponse(c, "Provide token in following format <Bearer 1w56wdgh7>")
		}

		accessToken := authHeaderSplit[1]

		claims, err := common.ParseJWTSignedAccessToken(accessToken)
		if err != nil {
			return common.SendUnauthorizedResponse(c, err.Error())
		}

		if common.IsClaimExpired(claims) {
			return common.SendUnauthorizedResponse(c, "Token is expired")
		}

		var user models.User
		result := appMiddleware.DB.First(&user, claims.ID)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return common.SendUnauthorizedResponse(c, "user not found")
		}
		if result.Error != nil {
			return common.SendUnauthorizedResponse(c, "Invalid access token")
		}

		c.Set("user", user)

		return next(c)
	}
}

// supply jwt token to the AuthMiddleware
// middleware intercepts and validates the user out
// if the token is invalid, we would bounce the user out
// middleware attaches current user with the current context

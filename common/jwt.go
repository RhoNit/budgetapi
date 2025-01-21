package common

import (
	"errors"
	"os"
	"time"

	"github.com/RhoNit/budgetapi/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type CustomJWTClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWTTokens(user *models.User) (*string, *string, error) {
	// define the claims to generate access token
	claims := CustomJWTClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	}

	// create access token with the claims
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the accessToken
	signedAccessToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, nil, err
	}

	// create refresh token with the claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomJWTClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
		},
	})

	// sign the accessToken
	signedRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, nil, err
	}

	return &signedAccessToken, &signedRefreshToken, nil
}

func ParseJWTSignedAccessToken(signedAccessToken string) (*CustomJWTClaims, error) {
	parsedJWTAccessToken, err := jwt.ParseWithClaims(signedAccessToken, &CustomJWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	} else if claims, ok := parsedJWTAccessToken.Claims.(CustomJWTClaims); ok {
		return &claims, nil
	} else {
		return nil, errors.New("unknown claim type, can't proceed")
	}
}

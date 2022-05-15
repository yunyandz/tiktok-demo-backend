package jwtx

import (
	"github.com/yunyandz/tiktok-demo-backend/internal/errorx"

	"github.com/yunyandz/tiktok-demo-backend/internal/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	EXPIRES_TIME = 3600 // s
	ISSUER       = "x"
	SIGNED_KEY   = "666666"
)

type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateUserClaims(username string) (string, error) {
	// Create the Claims
	claims := UserClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * EXPIRES_TIME).Unix(),
			Issuer:    ISSUER,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(SIGNED_KEY))
	if err != nil {
		logger.Suger().Errorw("CreateUserClaims SignedString failed", "err", err.Error())
		return "", err
	}
	return ss, err
}

func ParseUserClaims(tokenString string) (*UserClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Suger().Error("ParseWithClaims but without SigningMethodHMAC method")
			return nil, errorx.ErrTokenMethod
		}
		return []byte(SIGNED_KEY), nil
	})

	if err != nil {
		//handle error in more detail
		// TODO:
		// It may be necessary to customize the Valid function which throw a custom error if an error occurs.
		// Allow for more detailed error handling.

		//if verr , ok := err.(*jwt.ValidationError); ok && errors.Is(verr.Inner, errorx.ExpiredTokenErr);

		return nil, err
		//return nil, errors.New(errorx.ParseJWTFailed)
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errorx.ErrInvalidToken
}

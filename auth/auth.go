package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type customClaims struct {
	UserID string `json:"userid"`
	jwt.StandardClaims
}

const expiration = 30

func GenerateToken(userID string) (string, error) {
	claims := customClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * expiration).Unix(),
			IssuedAt:  jwt.TimeFunc().Unix(),
			Issuer:    "dazwallet",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	authKey := os.Getenv("AUTH_KEY")
	signedToken, err := token.SignedString([]byte(authKey))

	return signedToken, err
}

func GetUserIDFromToken(t string) (string, error) {

	authKey := os.Getenv("AUTH_KEY")
	fmt.Println(t)

	token, err := jwt.ParseWithClaims(
		t,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(authKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return "", errors.New("jwt is expired")
	}

	ui := claims.UserID
	return ui, nil
}

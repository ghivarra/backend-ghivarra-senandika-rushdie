package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTData jwt.MapClaims

func SignJWT(username string, roleID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  username,
		"role": roleID,
		"iat":  time.Now(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	key := []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(bearerToken string) (bool, error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return false, fmt.Errorf("failed to sign the JWT")
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return false, err
	}

	// check claim
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// store claims in admin
		// to be used later
		JWTData = claims
	} else {
		return false, fmt.Errorf("failed to get JWT payload")
	}

	return true, nil
}

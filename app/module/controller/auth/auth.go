package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/ghivarra/app/module/database"
	"github.com/ghivarra/app/module/library"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func signJWT(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud": data.Username,
		"sub": data.UserRoleID,
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	key := []byte(os.Getenv("JWT_KEY"))
	tokenString, err := token.SignedString(key)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Register(c *gin.Context) {
	var user UserRegister
	c.ShouldBindBodyWithJSON(&user)

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	database.Connect()
	query, param, err := sq.Insert(`"user"`).Columns("username", "password", "email", "name", "user_role_id").Values(user.Username, user.Password, user.Email, user.Name, user.UserRoleID).ToSql()

	fmt.Println(query, param)

	if err != nil {
		library.ErrorServer("Failed to connect to database", err, c)
	}

	_, err = database.DB.Query(context.Background(), query, param...)
	if err != nil {
		library.ErrorServer("Failed to query database", err, c)
	}
}

func Login(c *gin.Context) {
	var user UserLogin
	c.ShouldBindBodyWithJSON(&user)

	database.Connect()

	condition := squirrel.Eq{"username": user.Username}

	query, param, err := sq.Select("id", "password", "user_role_id").From(`"user"`).Where(condition).Limit(1).ToSql()

	if err != nil {
		library.ErrorServer("Failed to connect to database", err, c)
	}

	result, err := database.DB.Query(context.Background(), query, param...)
	if err != nil {
		library.ErrorServer("Failed to query database", err, c)
	}

	id := ""
	passwordHash := ""
	userRoleID := ""

	for result.Next() {
		result.Scan(&id, &passwordHash, &userRoleID)
	}

	check := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))

	if check != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Akun dan password tidak cocok",
		})
	}

	// sign with JWT
	var jwtData JWTData
	jwtData.Username = user.Username
	jwtData.UserRoleID = userRoleID
	token, errSign := signJWT(jwtData)

	if errSign != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Gagal login, ada masalah pada sistem",
		})
		return
	}

	fmt.Println(token)

	// return berhasil
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Anda berhasil login!",
		"data": gin.H{
			"token": token,
		},
	})
}

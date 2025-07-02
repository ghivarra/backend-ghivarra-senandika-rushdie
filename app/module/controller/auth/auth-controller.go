package auth

import (
	"fmt"
	"strings"

	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library/jwt"
	userModel "github.com/ghivarra/app/module/model/user-model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var form UserRegister
	c.ShouldBindBodyWithJSON(&form)

	// hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	form.Password = string(hash)

	// connect DB
	database.Connect()

	// seed new user
	newUser := userModel.User{Username: form.Username, Password: form.Password, Name: form.Name, Email: form.Email, UserRoleID: uint(form.UserRoleID)}

	// fail if there is same username or email
	var total int64
	database.CONN.Model(&userModel.User{}).Where("username = ?", form.Username).Or("email = ?", form.Email).Count(&total)

	if total > 0 {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": "Username atau Email sudah digunakan oleh akun lain",
		})
		return
	}

	err := database.CONN.Create(&newUser).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": "Gagal menambah user baru",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("User %s berhasil ditambahkan", form.Name),
	})
}

func Login(c *gin.Context) {
	var form UserLogin
	c.ShouldBindBodyWithJSON(&form)

	// connect db and load model
	database.Connect()

	// get result
	type PartialUser struct {
		Password   string
		UserRoleID uint
	}
	var user PartialUser
	database.CONN.Model(userModel.User{}).Select("password", "user_role_id").Where("username = ?", form.Username).First(&user)

	if user.Password == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Kombinasi akun dan password tidak cocok",
		})
		return
	}

	// validate password
	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password))

	if check != nil {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Kombinasi akun dan password tidak cocok",
		})
	}

	// sign with JWT
	token, errSign := jwt.SignJWT(form.Username, int(user.UserRoleID))
	if errSign != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Gagal login, ada masalah pada sistem",
		})
		return
	}

	// return berhasil
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Anda berhasil login!",
		"data": gin.H{
			"token": token,
		},
	})
}

func Check(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	bearerToken := strings.Replace(authHeader, "Bearer ", "", 1)
	if bearerToken == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda belum terotentikasi",
		})
		return
	}

	// validate
	valid, err := jwt.ValidateJWT(bearerToken)
	if err != nil || !valid {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": "Anda belum terotentikasi",
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Anda sudah terotentikasi",
	})
}

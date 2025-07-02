package service

import (
	"context"
	"fmt"

	"github.com/ghivarra/app/module/database"
	"github.com/ghivarra/app/module/library"
	"github.com/gin-gonic/gin"
)

// Get all product
func UserGetAll(c *gin.Context) {
	query, param, err := builder.Select(`"user".id`, "username", "password", `"user".name`, "email", "user_role_id", `"user_role".name AS user_role_name`).From(`"user"`).Join(`"user_role" ON user_role_id = "user_role".id`).ToSql()

	if err != nil {
		library.ErrorServer("Failed to connect to database", err, c)
	}

	database.Connect()

	result, err := database.DB.Query(context.Background(), query, param...)
	if err != nil {
		library.ErrorServer("Failed to query database", err, c)
	}

	var users []User

	for result.Next() {

		var user User

		err = result.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Email, &user.UserRoleID, &user.UserRoleName)

		if err != nil {
			fmt.Println(err)
		}

		users = append(users, user)
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    users,
	})
}

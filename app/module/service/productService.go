package service

import (
	"context"
	"fmt"

	"github.com/ghivarra/app/module/database"
	"github.com/ghivarra/app/module/library"
	"github.com/gin-gonic/gin"
)

// Get all product
func ProductGet(c *gin.Context) {
	query, param, err := builder.Select("id, username, password, name, email, user_role_id").From("user").ToSql()

	fmt.Println(query, param)

	if err != nil {
		library.ErrorServer("Failed to connect to database", err, c)
	}

	database.Connect()

	result, err := database.DB.Query(context.Background(), query, param...)
	if err != nil {
		library.ErrorServer("Failed to query database", err, c)
	}

	var users []string

	fmt.Println(result, users)
}

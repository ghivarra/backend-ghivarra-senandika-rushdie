package userRole

import (
	"github.com/ghivarra/app/database"
	userRoleModel "github.com/ghivarra/app/module/model/user-role"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	// connect db
	database.Connect()

	// get all roles
	type PartialUserRole struct {
		ID   uint
		Name string
	}
	var roles []PartialUserRole
	database.CONN.Model(&userRoleModel.UserRole{}).Find(&roles)

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Berhasil menarik data",
		"data":    roles,
	})
}

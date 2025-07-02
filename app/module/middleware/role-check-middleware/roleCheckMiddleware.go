package roleCheckMiddleware

import (
	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library/jwt"
	"github.com/ghivarra/app/module/model"
	"github.com/gin-gonic/gin"
)

func Run(c *gin.Context) {
	// connect database and check roles
	database.Connect()

	// get module
	type PartialModule struct {
		ID   uint
		Name string
	}
	var module PartialModule
	database.CONN.Model(&model.UserModule{}).Select("id", "name").Where("name = ?", c.Request.URL).First(&module)

	// module not found
	// then it is for public user
	if module.Name == "" {
		return
	}

	// check modules, if empty then abort with 403 forbidden
	var total int64
	database.CONN.Model(&model.UserRoleModuleList{}).Where("user_role_id = ?", jwt.JWTData["role"]).Where("user_module_id = ?", module.ID).Count(&total)

	if total < 1 {
		c.AbortWithStatusJSON(403, gin.H{
			"status":  "error",
			"message": "Anda tidak memiliki izin untuk mengakses halaman ini",
		})
		return
	}
}

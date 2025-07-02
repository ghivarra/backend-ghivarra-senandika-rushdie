package product

import (
	"github.com/ghivarra/app/database"
	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {

}

func Get(c *gin.Context) {

}

func Create(c *gin.Context) {
	var form ProductCreate
	c.ShouldBind(&form)

	// connect DB
	database.Connect()

	// get user

}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

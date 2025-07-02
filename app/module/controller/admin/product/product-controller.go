package product

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library"
	"github.com/ghivarra/app/module/library/jwt"
	"github.com/ghivarra/app/module/model"
	"github.com/gin-gonic/gin"
)

func All(c *gin.Context) {

}

func Get(c *gin.Context) {

}

func Create(c *gin.Context) {
	var form ProductCreate
	errForm := c.ShouldBind(&form)
	if errForm != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": errForm.Error(),
		})
		return
	}

	// validate access
	username, err := jwt.JWTData.GetSubject()
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"status":  "error",
			"message": errForm.Error(),
		})
		return
	}

	fmt.Println(username)

	// connect DB
	database.Connect()

	// parse form
	var Product model.Product
	price, _ := strconv.Atoi(form.Price)
	stock, _ := strconv.Atoi(form.Stock)

	nameAsSlug := strings.ReplaceAll(form.Name, " ", "")
	nameAsSlug = nameAsSlug[:120]
	nameAsSlug = fmt.Sprintf("%s-%s", nameAsSlug, library.RandomString(20))

	Product.Name = form.Name
	Product.Description = form.Description
	Product.Price = uint(price)
	Product.Stock = uint(stock)
	Product.Slug = strings.ToLower(nameAsSlug)

	// convert all slug into -
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

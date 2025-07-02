package product

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ghivarra/app/common"
	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library"
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

	// connect DB
	database.Connect()

	// parse form
	var Product model.Product
	price, _ := strconv.Atoi(form.Price)
	stock, _ := strconv.Atoi(form.Stock)

	// check name length
	limit := min(len(form.Name), 120)

	nameAsSlug := strings.ReplaceAll(form.Name, " ", "-")
	nameAsSlug = nameAsSlug[:limit]
	nameAsSlug = fmt.Sprintf("%s-%s", nameAsSlug, library.RandomString(20))

	Product.Name = form.Name
	Product.Description = form.Description
	Product.Price = uint(price)
	Product.Stock = uint(stock)
	Product.Slug = strings.ToLower(nameAsSlug)

	// add photo
	fmt.Println(common.ROOT)
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

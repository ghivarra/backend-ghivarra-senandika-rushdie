package cart

import (
	"strconv"

	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library/jwt"
	"github.com/ghivarra/app/module/model"
	"github.com/gin-gonic/gin"
)

type ProductAdd struct {
	ProductID int `json:"product_id" binding:"required,numeric"`
}

func AddProduct(c *gin.Context) {
	var form ProductAdd
	errForm := c.ShouldBindBodyWithJSON(&form)
	if errForm != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": errForm.Error(),
		})
		return
	}

	// database
	database.Connect()

	// parse data
	userID, _ := jwt.JWTData.GetSubject()
	intUserID, _ := strconv.Atoi(userID)

	var cart model.Cart
	cart.UserID = uint(intUserID)
	cart.ProductID = uint(form.ProductID)

	err := database.CONN.Create(&cart).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Produk berhasil ditambahkan ke keranjang",
	})
}

func Get(c *gin.Context) {

}

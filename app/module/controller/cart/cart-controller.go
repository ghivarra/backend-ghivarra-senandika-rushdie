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
	// parse data
	userID, _ := jwt.JWTData.GetSubject()

	// connect db
	database.Connect()

	type CartData struct {
		ID           int
		UserID       int
		ProductID    int
		ProductName  string
		Price        int
		Stock        int
		Photo        string
		Slug         string
		MerchantID   int
		MerchantName string
	}

	var products []CartData
	database.CONN.Model(&model.Cart{}).
		Select(`"cart".id`, `"cart".user_id`, `"cart".product_id`, `"product".name as product_name`, "price", "stock", "Photo", "slug", `"product".user_id as merchant_id`, `"user".name as merchant_name`).
		Joins(`JOIN "product" ON "cart".product_id = "product".id`).
		Joins(`JOIN "user" ON "product".user_id = "user".id`).
		Where(`"cart".user_id = ?`, userID).
		Scan(&products).Debug()

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Produk di keranjang berhasil ditarik",
		"data":    products,
	})
}

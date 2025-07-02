package cart

import (
	"encoding/json"
	"strconv"

	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library/jwt"
	"github.com/ghivarra/app/module/model"
	"github.com/gin-gonic/gin"
)

type ProductAdd struct {
	ProductID int `json:"product_id" binding:"required,numeric"`
	Quantity  int `json:"quantity" binding:"required,numeric"`
}

type ProductBuy struct {
	CartIDS []int `json:"cart_ids" binding:"required"`
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
	cart.Quantity = uint(form.Quantity)

	// check if already exist
	type OldData struct {
		ID       int
		Quantity int
	}
	var oldCartData OldData
	database.CONN.Model(&model.Cart{}).Select("id", "quantity").Where("user_id = ?", cart.UserID).Where("product_id = ?", cart.ProductID).First(&oldCartData)

	if oldCartData.ID != 0 {
		cart.ID = uint(oldCartData.ID)
		cart.Quantity = uint(oldCartData.Quantity + form.Quantity)
	}

	err := database.CONN.Save(&cart).Error
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
		Quantity     int
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
		Select(`"cart".id`, `"cart".user_id`, `"cart".product_id`, "quantity", `"product".name as product_name`, "price", "stock", "Photo", "slug", `"product".user_id as merchant_id`, `"user".name as merchant_name`).
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

func Buy(c *gin.Context) {
	var form ProductBuy
	errForm := c.ShouldBindBodyWithJSON(&form)
	if errForm != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": errForm.Error(),
		})
		return
	}

	database.Connect()
	// parse data
	userID, _ := jwt.JWTData.GetSubject()
	intUserID, _ := strconv.Atoi(userID)

	// get all data
	type CartItem struct {
		model.Cart
		Price int
	}

	var items []CartItem
	database.CONN.Model(&model.Cart{}).
		Select(`"cart".*`, "price").
		Joins(`JOIN "product" ON "cart".product_id = "product".id`).
		Find(&items, form.CartIDS)

	// calculate price
	price := 0
	orderIDS := []int{}

	for _, product := range items {
		var orderItem model.Order
		orderItem.BuyerID = uint(intUserID)
		orderItem.SellerID = product.UserID
		orderItem.ProductID = product.ID
		orderItem.Price = uint(product.Price)
		orderItem.Quantity = product.Quantity

		// insert into database
		// and delete from cart
		database.CONN.Create(&orderItem)
		orderIDS = append(orderIDS, int(orderItem.ID))

		// increase price for invoice
		price += product.Price * int(product.Quantity)
	}

	// delete from cart
	var deletedData []model.Order
	database.CONN.Model(&model.Cart{}).Where("id IN ?", form.CartIDS).Delete(&deletedData)

	// bonus
	bonuses := []string{}
	realPrice := price

	// check price for gratis ongkir
	if price > 15000 {
		bonuses = append(bonuses, "gratis-ongkir")
	}

	// check price for diskon 10%
	if price > 50000 {
		bonuses = append(bonuses, "diskon-10%")
		price = price * 90 / 100
	}

	// convert bonus to json
	bonusJSON, _ := json.Marshal(bonuses)

	// add details to invoice
	var invoiceData model.Invoice
	invoiceData.Price = uint(price)
	invoiceData.Details = string(bonusJSON)
	invoiceData.UserID = uint(intUserID)

	// insert invoice
	database.CONN.Create(&invoiceData)

	// foreach product
	// and insert into invoice order list
	for _, orderID := range orderIDS {
		var invoiceOrderList model.InvoiceOrderList
		invoiceOrderList.InvoiceID = invoiceData.ID
		invoiceOrderList.OrderID = uint(orderID)

		database.CONN.Create(&invoiceOrderList)
	}

	// return yang harus dibayar
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Order berhasil!",
		"data": gin.H{
			"price_before": realPrice,
			"price_after":  price,
			"discount":     realPrice - price,
			"bonus":        bonuses,
		},
	})
}

package order

import (
	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library"
	"github.com/ghivarra/app/module/model"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	// connect DB
	database.Connect()

	type OrderDetail struct {
		model.Order
		ProductName string
		Photo       string
		Stock       int
	}

	intUserID := library.GetCurrentUserID()

	var orders []OrderDetail
	database.CONN.Model(&model.Order{}).
		Select(`"order".*`, `"product".name as product_name`, "Photo", "Stock").
		Where("user_id = ?", intUserID).
		Joins(`"product" ON "order".product_id = "product".id`).
		Find(&orders)

	// send
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Order berhasil ditarik",
		"data":    orders,
	})
}

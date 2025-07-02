package product

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ghivarra/app/common"
	"github.com/ghivarra/app/database"
	"github.com/ghivarra/app/module/library"
	"github.com/ghivarra/app/module/library/jwt"
	"github.com/ghivarra/app/module/model"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	// connect db
	database.Connect()

	// get all products
	type PartialProduct struct {
		ID           uint
		Slug         string
		Name         string
		Description  string
		Photo        string
		Price        uint
		Stock        uint
		MerchantID   uint
		MerchantName string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	// check if singular or plural
	queryID := c.Query("id")
	var result any
	if queryID == "" {
		var products []PartialProduct
		database.CONN.Model(&model.Product{}).Select(`"product".id`, `"product".name`, "description", "photo", "price", "stock", "slug", "user_id as merchant_id", `"user".name as merchant_name`, `"product".created_at`, `"product".updated_at`).Joins(`JOIN "user" ON user_id = "user".id`).Find(&products)
		result = products
	} else {
		var products PartialProduct
		database.CONN.Model(&model.Product{}).Select(`"product".id`, `"product".name`, "description", "photo", "price", "stock", "slug", "user_id as merchant_id", `"user".name as merchant_name`, `"product".created_at`, `"product".updated_at`).Joins(`JOIN "user" ON user_id = "user".id`).Where(`"product".id = ?`, queryID).Limit(1).First(&products)
		result = products
	}

	// get all
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil ditarik",
		"data":    result,
	})
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

	// parse form
	var data model.Product
	price, _ := strconv.Atoi(form.Price)
	stock, _ := strconv.Atoi(form.Stock)

	// check name length
	limit := min(len(form.Name), 120)

	nameAsSlug := strings.ReplaceAll(form.Name, " ", "-")
	nameAsSlug = nameAsSlug[:limit]
	nameAsSlug = fmt.Sprintf("%s-%s", nameAsSlug, library.RandomString(20))

	data.Name = form.Name
	data.Description = form.Description
	data.Price = uint(price)
	data.Stock = uint(stock)
	data.Slug = strings.ToLower(nameAsSlug)

	userID := library.GetCurrentUserID()
	data.UserID = uint(userID)

	fmt.Println(jwt.JWTData.GetSubject())

	// upload foto
	dotIndex := strings.LastIndex(form.Photo.Filename, ".")
	if dotIndex == -1 {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": "Tidak ada ekstensi pada foto yang diupload",
		})
	}
	ext := form.Photo.Filename[dotIndex:]

	fileName := fmt.Sprintf("%s%s", data.Slug, ext)
	c.SaveUploadedFile(form.Photo, fmt.Sprintf("%s/upload/%s", common.ROOT, fileName))

	// tambah foto
	data.Photo = fileName

	// simpan foto di DB
	database.Connect()
	err := database.CONN.Create(&data).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil diinput",
	})
}

func Update(c *gin.Context) {
	var form ProductUpdate
	errForm := c.ShouldBind(&form)
	if errForm != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": errForm.Error(),
		})
		return
	}

	// parse form
	var data model.Product
	price, _ := strconv.Atoi(form.Price)
	stock, _ := strconv.Atoi(form.Stock)

	// check name length
	limit := min(len(form.Name), 120)

	nameAsSlug := strings.ReplaceAll(form.Name, " ", "-")
	nameAsSlug = nameAsSlug[:limit]
	nameAsSlug = fmt.Sprintf("%s-%s", nameAsSlug, library.RandomString(20))

	data.ID = uint(form.ID)
	data.Name = form.Name
	data.Description = form.Description
	data.Price = uint(price)
	data.Stock = uint(stock)

	userID := library.GetCurrentUserID()
	data.UserID = uint(userID)

	fmt.Println(jwt.JWTData.GetSubject())

	// upload foto
	dotIndex := strings.LastIndex(form.Photo.Filename, ".")
	if dotIndex == -1 {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": "Tidak ada ekstensi pada foto yang diupload",
		})
	}
	ext := form.Photo.Filename[dotIndex:]

	fileName := fmt.Sprintf("%s%s", nameAsSlug, ext)
	errUpload := c.SaveUploadedFile(form.Photo, fmt.Sprintf("%s/upload/%s", common.ROOT, fileName))

	if errUpload != nil {
		// update foto
		data.Photo = fileName
	}

	// simpan foto di DB
	database.Connect()
	err := database.CONN.Model(&model.Product{}).Where("id = ?", form.ID).Save(&data).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// return
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Data berhasil diinput",
	})
}

func Delete(c *gin.Context) {
	var form ProductDelete
	errForm := c.ShouldBindBodyWithJSON(&form)
	if errForm != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"status":  "error",
			"message": errForm.Error(),
		})
		return
	}

	// delete
	var deletedData model.Product
	database.Connect()
	database.CONN.Model(&model.Product{}).
		Where("id = ?", form.ID).
		Delete(&deletedData)

	// return
	c.JSON(200, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Produk %s berhasil dihapus", deletedData.Name),
	})
}

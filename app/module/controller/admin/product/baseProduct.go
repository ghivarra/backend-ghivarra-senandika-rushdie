package product

import (
	"mime/multipart"
)

type ProductCreate struct {
	Name        string                `form:"name" binding:"required,max=150"`
	Description string                `form:"description"`
	Price       string                `form:"price" binding:"required,numeric"`
	Stock       string                `form:"stock" binding:"required,numeric"`
	Photo       *multipart.FileHeader `form:"photo"`
}

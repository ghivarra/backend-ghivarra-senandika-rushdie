package product

import (
	"mime/multipart"
)

type ProductCreate struct {
	Name        string                `form:"name" binding:"required,max=150"`
	Description string                `form:"description"`
	Price       string                `form:"price" binding:"required,numeric"`
	Stock       string                `form:"stock" binding:"required,numeric"`
	Photo       *multipart.FileHeader `form:"photo" binding:"required"`
}

type ProductUpdate struct {
	ID          int                   `form:"id" binding:"required,numeric"`
	Name        string                `form:"name" binding:"max=150"`
	Description string                `form:"description"`
	Price       string                `form:"price" binding:"numeric"`
	Stock       string                `form:"stock" binding:"numeric"`
	Photo       *multipart.FileHeader `form:"photo"`
}

type ProductDelete struct {
	ID int `form:"id" binding:"required,numeric"`
}

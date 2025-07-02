package product

import (
	"mime/multipart"

	"github.com/Masterminds/squirrel"
)

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
var roleName string

type ProductCreate struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Photo       *multipart.FileHeader `form:"photo"`
	Price       string                `form:"price"`
	Stock       string                `form:"stock"`
}

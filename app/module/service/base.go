package service

import "github.com/Masterminds/squirrel"

// connect db and builder
var builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type ProductType struct {
	id           string
	name         string
	description  string
	photo        string
	price        string
	stock        string
	slug         string
	user_role_id string
}

type User struct {
	id           string
	username     string
	password     string
	name         string
	email        string
	user_role_id string
}

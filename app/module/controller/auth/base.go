package auth

import "github.com/Masterminds/squirrel"

var sq = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type UserRegister struct {
	Username   string `json:"username" bind:"alpha"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	UserRoleID int    `json:"user_role_id"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTData struct {
	Username   string
	UserRoleID string
}

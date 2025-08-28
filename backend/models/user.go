package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	User_id   int    `json:"user_id"`
	User_name int    `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Point     int    `json:"point"`
	Role      string `json:"role"`
}

type Claims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

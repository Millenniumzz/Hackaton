package handlers

import (
	"database/sql"
	"time"

	"hackaton-backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("my_secret_key")

func HandleRegister(c *fiber.Ctx, db *sql.DB) error {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	role := "user"
	if req.Role == "admin" {
		role = "admin"
	}

	_, err := db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)",
		req.Username, req.Email, string(hashedPassword), role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.JSON(fiber.Map{"message": "Register success"})
}

func HandleLogin(c *fiber.Ctx, db *sql.DB) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	err := db.QueryRow("SELECT id, username, email, password, role FROM users WHERE username = ?", req.Username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Username not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	claims := &models.Claims{
		ID:   user.ID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JwtKey)

	return c.JSON(fiber.Map{"token": tokenString, "role": user.Role})
}

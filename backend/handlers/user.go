package handlers

import (
	"database/sql"
	"hackaton-backend/models"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT id, username, email, role FROM users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role)
		users = append(users, u)
	}

	return c.JSON(users)
}

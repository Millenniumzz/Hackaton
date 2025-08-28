package handlers

import (
	"database/sql"
	"hackaton-backend/models"

	//"hackaton-backend/models"

	"github.com/gofiber/fiber/v2"
)

func HandleProduct(c *fiber.Ctx, db *sql.DB) error {

	/*var res struct {
		Product_id          int    `json:"product_id"`
		Product_name        string `json:"product_name"`
		Product_cost        int    `json:"product_cost"`
		Product_type        int    `json:"product_type"`
		Product_description string `json:"product_description"`
		Product_picture     []byte `json:"product_picture"`
	}

	var req struct{
		User_id int `json:"user_id"`
	}*/

	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Product already exists"})
	}
	defer rows.Close()

	var product []models.Product
	for rows.Next() {
		var u models.Product
		err := rows.Scan(&u.Product_id, &u.Product_name, &u.Product_cost, &u.Product_type, &u.Product_description, &u.Product_picture)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		product = append(product, u)
	}

	return c.JSON(product)
}

func HandleAddProduct(c *fiber.Ctx, db *sql.DB) error {

	var req struct {
		Product_name        string `json:"product_name"`
		Product_cost        int    `json:"product_cost"`
		Product_type        int    `json:"product_type"`
		Product_description string `json:"product_description"`
		Product_picture     []byte `json:"product_picture"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := db.Exec("INSERT INTO product (product_name, product_cost, product_type, product_description,product_picture) VALUES (?, ?, ?, ?,?)",
		req.Product_name, req.Product_cost, req.Product_type, req.Product_description, req.Product_picture)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "already exists"})
	}

	return c.JSON(fiber.Map{"message": "add success"})
}

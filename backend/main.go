package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "modernc.org/sqlite"

	"hackaton-backend/handlers"
	//"hackaton-backend/middleware"
)

func main() {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		user_id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		password TEXT,
		point INTEGER,
		role TEXT DEFAULT 'user'
	)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS product (
		product_id INTEGER PRIMARY KEY AUTOINCREMENT,
		product_name TEXT,
		product_cost INTEGER,
		product_type TEXT,
		product_description TEXT,
		product_picture BLOB
	)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS trans (
			trans_id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			product_id INTEGER,
			date TEXT,
			FOREIGN KEY(user_id) REFERENCES users(user_id),
			FOREIGN KEY(product_id) REFERENCES product(product_id)
		)`)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Routes
	app.Post("/register", func(c *fiber.Ctx) error { return handlers.HandleRegister(c, db) })
	app.Post("/login", func(c *fiber.Ctx) error { return handlers.HandleLogin(c, db) })
	app.Get("/list", func(c *fiber.Ctx) error { return handlers.HandleProduct(c, db) })
	app.Post("/add", func(c *fiber.Ctx) error { return handlers.HandleAddProduct(c, db) })
	/*
		app.Get("/users", middleware.AuthMiddleware("admin"), func(c *fiber.Ctx) error {
			return handlers.HandleGetUsers(c, db)
		})
	*/
	app.Get("/users", func(c *fiber.Ctx) error { return handlers.HandleGetUsers(c, db) })

	log.Println("Server running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))

}

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
	db, err := sql.Open("sqlite", "./db/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT UNIQUE,
		password TEXT,
		role TEXT DEFAULT 'user'
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
	app.Get("/users", func(c *fiber.Ctx) error {
		return handleGetUsers(c, db)
	})

	log.Println("Backend running at http://localhost:8080")

	/*
		app.Get("/users", middleware.AuthMiddleware("admin"), func(c *fiber.Ctx) error {
			return handlers.HandleGetUsers(c, db)
		})
	*/

	log.Fatal(app.Listen(":8080"))
}

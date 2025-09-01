package main

import (
	"summarizer-backend/handlers"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("env not found")
	}
	app := fiber.New()

	app.Use(logger.New())

	app.Post("/summarize", handlers.SummarizeHandler)

	app.Listen(":8080")
}
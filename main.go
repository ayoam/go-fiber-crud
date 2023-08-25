package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	router := fiber.New()
	app := fiber.New()

	app.Mount("/api/v1", router)

	//test
	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "welcome to golang!",
		})
	})

	// Start server
	log.Fatal(app.Listen(":8000"))
}

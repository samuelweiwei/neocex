package main

import (
	"neocex/v2/logging"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	logging.Logger.Info("Starting server......")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}

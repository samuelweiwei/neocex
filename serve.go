package main

import (
	"neocex/v2/logging"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// app := fiber.New()
	app := fiber.New(fiber.Config{
		Prefork:               true,
		ServerHeader:          "Fiber",
		StrictRouting:         true,
		CaseSensitive:         true,
		DisableStartupMessage: false,
	})

	logging.Logger.Info("Starting server......")

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })
	app.Get("/", handleRoot)

	if err := app.Listen(":3000"); err != nil {
		logging.Logger.Fatal(err.Error())
	}
}

func handleRoot(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

package router

import "github.com/gofiber/fiber/v2"

func serve() error {
	app := fiber.New()

	app.Get("/liveness", func(c *fiber.Ctx) error {
		return c.SendString("I am alive......")
	}).Name("liveness")

	app.Get("/readiness", func(c *fiber.Ctx) error {
		return c.SendString("I am ready to serve......")
	})

	return nil
}

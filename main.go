package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(
		fiber.Config{
			AppName:           "Go-API",
			Prefork:           true,
			GETOnly:           true,
			EnablePrintRoutes: true,
		},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "alive"})
	})

	app.Listen(":3000")
}

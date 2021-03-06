package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New(
		fiber.Config{
			AppName:           "Go-API",
			Prefork:           true,
			GETOnly:           true,
			EnablePrintRoutes: true,
		},
	)

	app.Use(logger.New())
	app.Use(
		basicauth.New(
			basicauth.Config{
				Users: map[string]string{
					"admin": "admin",
				},
				Realm: "Forbidden",
			},
		),
	)
	app.Get("/dashboard", monitor.New())

	app.Get(
		"/", func(c *fiber.Ctx) error {
			return c.SendString("Hello, World!")
		},
	)

	app.Get(
		"/status", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{"status": "alive"})
		},
	)

	app.Listen(":80")
}

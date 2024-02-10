package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("hello world, this my web sack path /x for value and path /name for yourname")
	// })
	app.Static("/", "./public")

	app.Get("/x/:x", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("x"))
	})

	app.Get("/name/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("ສະບາຍດີ " + c.Params("name"))
		}
		return c.SendString("ເຈົ້າຈົ່ງບອກຊື່ມາດຽວນີ້ ")
	})
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("api is : " + c.Params("*"))

	})

	app.Listen(":8000")
}

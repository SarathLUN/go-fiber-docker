package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln("Can not start server on port 3000: ", err)
	}
}

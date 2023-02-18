package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// declare configuration
	//appBaseURL := "localhost"
	appPort := "3000"
	// initialize fiber app
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})
	err := app.Listen(":" + appPort)
	if err != nil {
		log.Fatalln("Can not start server on port 3000: ", err)
	}
}

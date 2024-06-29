package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
)

func main(){
	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Railway!")
	})

	app.Get("/env", func(c *fiber.Ctx) error{
		return c.SendString("Hello, ENV! " + os.Getenv("TEST_ENV"))
	})

	app.Get("/tjson", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
		 "message": "Hello",
		})
	   })

	port := os.Getenv("PORT")

	if port == ""{
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
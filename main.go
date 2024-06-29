package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Response struct {
    Message string `json:"message"`
    Status  int    `json:"status"`
}


func main(){
	
	app := fiber.New()

	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // Allow all origins
        AllowHeaders: "Origin, Content-Type, Accept",
    }))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Railway!")
	})

	app.Get("/env", func(c *fiber.Ctx) error{
		return c.SendString("Hello, ENV! " + os.Getenv("TEST_ENV"))
	})

	app.Get("/json", func(c *fiber.Ctx) error {
        response := Response{
            Message: "Hello, World!",
            Status:  200,
        }
        return c.JSON(response)
    })

	port := os.Getenv("PORT")

	if port == ""{
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}
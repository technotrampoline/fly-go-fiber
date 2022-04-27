package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"math/rand"
	"strconv"
)

const port = 8080

func main() {
	application := fiber.New()

	application.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "Hello, World!",
		})
	})

	application.Get("/random", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"number": rand.Intn(10),
		})
	})

	log.Fatalln(application.Listen(":" + strconv.Itoa(port)))
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id	int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
}

var todos = []Todo {
	{
		Id: 1,
		Name: "Arjuna",
		Completed: false,
	},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Lagi Belajar, Go!!")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
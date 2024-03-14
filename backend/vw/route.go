package vw

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	app.Post("/vws", func(c *fiber.Ctx) error {
		var request *CreateVWRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(409).SendString("Cannot parse body")
		}
		virtualWindows, err := CreateVW(request.Name, request.Os)
		if err != nil {
			fmt.Println("There was an error while trying to create a vw: ", err)
		}
		fmt.Println("Successfully created new vw: ", virtualWindows.Name)
		return c.Status(200).JSON(virtualWindows)
	})
	app.Get("/vws", func(c *fiber.Ctx) error {
		all, err := Repository.GetAll()
		if err != nil {
			return c.Status(200).JSON([]VirtualWindows{})
		}
		return c.Status(200).JSON(all)
	})
}

type CreateVWRequest struct {
	Name string `json:"name"`
	Os   string `json:"os"`
}

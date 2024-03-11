package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"virtual-windows-manager/auth"
	"virtual-windows-manager/database"
)

func main() {
	err := database.Connect("mongodb://localhost:27017", "vwm")
	if err != nil {
		panic("Cannot connect to the database")
	}
	auth.Repository = auth.NewMongoUserRepository()

	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("css/", "public/css")
	app.Static("images/", "public/images")

	users := []User{{Name: "Maga"}, {Name: "Justin"}}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "This is a title",
			"Users": users,
		}, "layouts/default")
	})
	fmt.Println(app.Listen(":3000"))
}

type User struct {
	Name string
}

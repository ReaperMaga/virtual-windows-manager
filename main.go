package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
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

	app.Get("/clicked", func(c *fiber.Ctx) error {
		return c.Render("partials/test", fiber.Map{
			"Text": "Works",
		})
	})
	fmt.Println(app.Listen(":3000"))
}

type User struct {
	Name string
}

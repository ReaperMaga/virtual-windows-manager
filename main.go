package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"virtual-windows-manager/auth"
	"virtual-windows-manager/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Please provide a env file")
	}

	err = database.Connect(os.Getenv("MONGODB_CONNECTION_URI"), "vwm")
	if err != nil {
		panic("Cannot connect to the database")
	}
	auth.Repository = auth.NewMongoUserRepository()

	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("css/", "public/css")
	app.Static("js/", "public/js")
	app.Static("images/", "public/images")

	users := []User{{Name: "Maga"}, {Name: "Justin"}}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "This is a title",
			"Users": users,
		}, "layouts/default")
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		user, err := auth.Repository.FindByNameOrErr("maga")
		if err != nil {
			return c.SendString("User does not exists")
		}
		success := auth.Repository.Delete(user)
		return c.SendString(strconv.FormatBool(success))
	})
	fmt.Println(app.Listen(":3000"))
}

type User struct {
	Name string
}

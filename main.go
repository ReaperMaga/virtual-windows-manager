package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"os"
	"virtual-windows-manager/auth"
	"virtual-windows-manager/database"
	"virtual-windows-manager/middleware"
	"virtual-windows-manager/vw"
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
	auth.Initialize()

	vw.Initialize()

	vw.CreateVW("test")

	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("css/", "public/css")
	app.Static("js/", "public/js")
	app.Static("images/", "public/images")

	app.Use(middleware.AuthMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "This is a title",
		}, "layouts/default", "layouts/panel")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		c.Accepts("application/x-www-form-urlencoded")
		username := c.FormValue("username")
		password := c.FormValue("password")

		success, user := auth.IsAuthAndGetUser(username, password)
		if !success {
			return c.Redirect("/login?error=wrong_credentials")
		}
		session, err := auth.CreateSession(user)
		if err != nil {
			return c.Redirect("/login?error=session_creation_failed")
		}
		c.Cookie(&fiber.Cookie{
			Name:    "session_token",
			Value:   session.Id,
			Expires: session.ExpireAt,
		})
		return c.Redirect("/")
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		foundErr := c.Query("error")
		return c.Render("login", fiber.Map{
			"Error": foundErr,
		}, "layouts/default")
	})

	fmt.Println(app.Listen(":3000"))
}

type User struct {
	Name string
}

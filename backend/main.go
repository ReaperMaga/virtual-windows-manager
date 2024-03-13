package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app := fiber.New()
	app.Use(cors.New())
	app.Use(middleware.AuthMiddleware)

	auth.InitRoutes(app)

	app.Post("/create", func(c *fiber.Ctx) error {
		c.Accepts("application/x-www-form-urlencoded")

		name := c.FormValue("name")
		osString := c.FormValue("os")

		virtualWindows, err := vw.CreateVW(name, osString)
		if err != nil {
			fmt.Println("There was an error while trying to create a vw: ", err)
		}
		fmt.Println("Successfully created new vw: ", virtualWindows.Name)
		return c.Redirect("/")
	})

	fmt.Println(app.Listen(":8080"))
}

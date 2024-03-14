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
	godotenv.Load()
	err := database.Connect(os.Getenv("MONGODB_CONNECTION_URI"), "vwm")
	if err != nil {
		panic("Cannot connect to the database")
	}
	auth.Initialize()

	vw.Initialize()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(middleware.AuthMiddleware)

	auth.InitRoutes(app)
	vw.InitRoutes(app)

	fmt.Println(app.Listen(":8080"))
}

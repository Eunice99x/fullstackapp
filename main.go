package main

import (
	"github.com/eunice99x/fullstackapp/src/database"
	"github.com/eunice99x/fullstackapp/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	//app.Use(cors.New(cors.Config{
	//	AllowCredentials: true,
	//}))

	routes.Setup(app)

	app.Listen(":3001")
}

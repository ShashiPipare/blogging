package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"main.go/connection"
	"main.go/routes"
)

func main() {
	log.Println("Main called!")
	connection.Init()
	connection.ConnectDB()
	app := fiber.New()
	routes.Configure(app)
	app.Listen(":8000")
}

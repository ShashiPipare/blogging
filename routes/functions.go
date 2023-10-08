package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/blog"
)

func Configure(app *fiber.App) {
	api := app.Group("/api")
	blog.Route(api)
}

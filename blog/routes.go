package blog

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func Route(router fiber.Router) {
	group := router.Group("/blog")
	group.Post("/create", timeout.NewWithContext(create, time.Second*10))
	group.Post("/update", timeout.NewWithContext(update, time.Second*10))
	group.Post("/delete", timeout.NewWithContext(delete, time.Second*10))
	group.Post("/getOne", timeout.NewWithContext(getOne, time.Second*10))
	group.Post("/getAll", timeout.NewWithContext(getAll, time.Second*10))

}

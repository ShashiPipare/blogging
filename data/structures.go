package data

import "github.com/gofiber/fiber/v2"

type model struct {
	*fiber.Ctx
}

type Object map[string]any

type Response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty""`
}

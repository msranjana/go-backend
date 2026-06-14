package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-dob-api/internal/handler"
)

func Register(app *fiber.App, h *handler.UserHandler) {
	u := app.Group("/users")
	u.Post("/", h.Create)
	u.Get("/", h.List)
	u.Get("/:id", h.GetByID)
	u.Put("/:id", h.Update)
	u.Delete("/:id", h.Delete)
}

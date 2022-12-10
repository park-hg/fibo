package user

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router, handler Handler) {
	router.Post("/", handler.Create)
	router.Get("/:id", handler.Get)
	router.Put("/:id/login", handler.Login)
	router.Put("/:id/logout", handler.Logout)
}

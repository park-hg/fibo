package comment

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router, handler Handler) {
	router.Post("/", handler.Create)
	router.Get("/:id", handler.Get)
	router.Get("/:post_id", handler.GetByPost)
	router.Put("/:id", handler.Save)
	router.Delete("/:id", handler.Delete)
}

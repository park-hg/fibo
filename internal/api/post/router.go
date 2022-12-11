package post

import "github.com/gofiber/fiber/v2"

func Rounter(router fiber.Router, handler Handler) {
	router.Post("/", handler.Create)
	router.Get("/:id", handler.Get)
	router.Get("/users/:author_id", handler.GetByAuthor)
	router.Get("/", handler.GetAll)
	router.Put("/:id", handler.Save)
	router.Delete("/:id", handler.Delete)
}

package user

import "github.com/gofiber/fiber/v2"

func Router(router fiber.Router, handler Handler) {
	router.Post("/", handler.Create)
	router.Get("/:name", handler.GetByName)
	router.Get("/:name/posts", handler.GetPostsByName)
	router.Get("/:name/comments", handler.GetCommentsByName)
	router.Put("/:id/login", handler.Login)
	router.Put("/:id/logout", handler.Logout)
	router.Get("/is-online", handler.IsLoggedIn)
}

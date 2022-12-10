package user

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	GetByName(c *fiber.Ctx) error
	GetPostsByName(c *fiber.Ctx) error
	GetCommentsByName(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	IsLoggedIn(c *fiber.Ctx) error
}

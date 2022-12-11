package post

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetByAuthor(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

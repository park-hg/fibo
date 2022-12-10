package comment

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Get(c *fiber.Ctx) error
	GetByPostID(c *fiber.Ctx) error
	GetByUserID(c *fiber.Ctx) error
	Save(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

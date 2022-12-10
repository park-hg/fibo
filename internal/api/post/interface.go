package post

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(ctx *fiber.Ctx) error
	Get(ctx *fiber.Ctx) error
	GetByAuthor(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Save(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

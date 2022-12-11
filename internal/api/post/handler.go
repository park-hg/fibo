package post

import (
	"fibo/internal/post/usecases"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	puc usecases.PostUseCases
}

func (h handler) Create(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) Get(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetByAuthor(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetAll(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) Save(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) Delete(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewHandler(puc usecases.PostUseCases) Handler {
	return &handler{puc: puc}
}

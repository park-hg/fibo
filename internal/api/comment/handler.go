package comment

import (
	"fibo/internal/comment/usecases"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	cuc usecases.CommentUseCases
}

func (h handler) Create(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) Get(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetByPostID(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetByUserID(c *fiber.Ctx) error {
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

func NewHandler(cuc usecases.CommentUseCases) Handler {
	return &handler{cuc: cuc}
}

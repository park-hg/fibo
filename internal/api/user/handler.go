package user

import (
	"fibo/internal/user/usecases"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	uuc usecases.UserUseCases
}

func (h handler) Create(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetByName(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetPostsByName(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) GetCommentsByName(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) Login(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) Logout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h handler) IsLoggedIn(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewHandler(uuc usecases.UserUseCases) Handler {
	return &handler{uuc: uuc}
}

package user

import (
	"context"
	"errors"
	"fibo/internal/user/domain"
	"fibo/internal/user/usecases"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type handler struct {
	uuc usecases.UserUseCases
}

func (h handler) Create(c *fiber.Ctx) error {
	var user domain.User

	c.Set("Content-Type", "application/json")
	err := c.BodyParser(&user)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.uuc.Create(context.Background(), user)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(&user)
}

func (h handler) GetByName(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	name := c.Params("name")
	if name == "" {
		return errors.New("name is empty")
	}

	users, err := h.uuc.GetByName(context.Background(), name)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(&users)
}

func (h handler) GetPostsByName(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	//TODO implement me
	panic("implement me")
}

func (h handler) GetCommentsByName(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	//TODO implement me
	panic("implement me")
}

func (h handler) Login(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	userID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.uuc.Login(context.Background(), uint(userID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return nil
}

func (h handler) Logout(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	userID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.uuc.Logout(context.Background(), uint(userID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return nil
}

func (h handler) IsLoggedIn(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	name := c.Params("name")
	isOnline, err := h.uuc.IsLoggedIn(context.Background(), name)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"online": isOnline,
	})
}

func NewHandler(uuc usecases.UserUseCases) Handler {
	return &handler{uuc: uuc}
}

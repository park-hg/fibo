package post

import (
	"context"
	"fibo/internal/post/domain"
	"fibo/internal/post/usecases"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type handler struct {
	puc usecases.PostUseCases
}

func (h handler) Create(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	post := domain.Post{}
	err := c.BodyParser(&post)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.puc.Create(context.Background(), post)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusCreated).JSON(&post)
}

func (h handler) Get(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	postID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	post, err := h.puc.Get(context.Background(), uint(postID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(&post)
}

func (h handler) GetByAuthor(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	authorID, err := strconv.ParseUint(c.Params("author_id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	posts, err := h.puc.GetByAuthor(context.Background(), uint(authorID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(&posts)
}

func (h handler) GetAll(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	posts, err := h.puc.GetAll(context.Background())
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(&posts)
}

func (h handler) Save(c *fiber.Ctx) error {

	var post domain.Post

	c.Set("Content-Type", "application/json")
	postID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = c.BodyParser(&post)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.puc.Save(context.Background(), uint(postID), post)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (h handler) Delete(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	postID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.puc.Delete(context.Background(), uint(postID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func NewHandler(puc usecases.PostUseCases) Handler {
	return &handler{puc: puc}
}

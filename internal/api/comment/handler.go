package comment

import (
	"context"
	"fibo/internal/comment/domain"
	"fibo/internal/comment/usecases"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type handler struct {
	cuc usecases.CommentUseCases
}

func (h handler) Create(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	comment := domain.Comment{}
	err := c.BodyParser(&comment)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.cuc.Create(context.Background(), comment)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusCreated).JSON(&comment)
}

func (h handler) Get(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	commentID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	comment, err := h.cuc.Get(context.Background(), uint(commentID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.Status(http.StatusOK).JSON(&comment)
}

func (h handler) GetByPostID(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	postID, err := strconv.ParseUint(c.Params("post_id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	comments, err := h.cuc.GetByPostID(context.Background(), uint(postID))
	return c.Status(http.StatusOK).JSON(&comments)
}

func (h handler) GetByAuthorID(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	authorID, err := strconv.ParseUint(c.Params("author_id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	comments, err := h.cuc.GetByPostID(context.Background(), uint(authorID))
	return c.Status(http.StatusOK).JSON(&comments)
}

func (h handler) Save(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	commentID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	comment := domain.Comment{}
	err = c.BodyParser(&comment)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.cuc.Save(context.Background(), uint(commentID), comment)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}

func (h handler) Delete(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	commentID, err := strconv.ParseUint(c.Params("id"), 0, 32)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	err = h.cuc.Delete(context.Background(), uint(commentID))
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}

func NewHandler(cuc usecases.CommentUseCases) Handler {
	return &handler{cuc: cuc}
}

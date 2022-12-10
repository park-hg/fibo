package usecases

import (
	"context"
	"fibo/internal/post/domain"
)

type PostUseCases interface {
	Create(ctx context.Context, p domain.Post) error
	Get(ctx context.Context, id uint) (domain.Post, error)
	GetByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error)
	GetAll() ([]domain.Post, error)
	Save(ctx context.Context, p domain.Post) error
	Delete(ctx context.Context, id uint) error
}

type usecases struct {
	postRepo domain.PostRepository
}

func (u usecases) Create(ctx context.Context, p domain.Post) error {
	//TODO implement me
	panic("implement me")
}

func (u usecases) Get(ctx context.Context, id uint) (domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecases) GetByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecases) GetAll() ([]domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (u usecases) Save(ctx context.Context, p domain.Post) error {
	//TODO implement me
	panic("implement me")
}

func (u usecases) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewPostUseCases(pr domain.PostRepository) PostUseCases {

	return &usecases{postRepo: pr}
}

package usecases

import (
	"context"
	"fibo/internal/post/domain"
)

type PostUseCases interface {
	Create(ctx context.Context, p domain.Post) error
	Get(ctx context.Context, postID uint) (domain.Post, error)
	GetByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error)
	GetAll(ctx context.Context) ([]domain.Post, error)
	Save(ctx context.Context, postID uint, p domain.Post) error
	Delete(ctx context.Context, postID uint) error
}

type usecases struct {
	postRepo domain.PostRepository
}

func (u usecases) Create(ctx context.Context, p domain.Post) error {
	return u.postRepo.Create(ctx, p)
}

func (u usecases) Get(ctx context.Context, postID uint) (domain.Post, error) {
	return u.postRepo.Get(ctx, postID)
}

func (u usecases) GetByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error) {
	return u.postRepo.GetByAuthor(ctx, authorID)
}

func (u usecases) GetAll(ctx context.Context) ([]domain.Post, error) {
	return u.postRepo.GetAll(ctx)
}

func (u usecases) Save(ctx context.Context, postID uint, p domain.Post) error {
	return u.postRepo.Save(ctx, postID, p)
}

func (u usecases) Delete(ctx context.Context, postID uint) error {
	return u.postRepo.Delete(ctx, postID)
}

func NewPostUseCases(pr domain.PostRepository) PostUseCases {
	return &usecases{postRepo: pr}
}

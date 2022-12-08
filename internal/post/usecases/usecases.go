package usecases

import (
	"context"
	"fibo/internal/post/domain"
)

type PostUseCases interface {
	Create(ctx context.Context, p *domain.Post) error
	Get(ctx context.Context, id int) (*domain.Post, error)
	GetByAuthor(ctx context.Context, author string) ([]domain.Post, error)
	GetAll() ([]domain.Post, error)
	Save(ctx context.Context, p *domain.Post) error
	Delete(ctx context.Context, id int) error
}

type usecases struct {
	postRepo domain.PostRepository
}

func NewPostUseCases(pr domain.PostRepository) PostUseCases {

	return &usecases{postRepo: pr}
}

func (uc usecases) Create(ctx context.Context, p *domain.Post) error {

}

func (uc usecases) Get(ctx context.Context, id int) (*domain.Post, error) {

}

func (uc usecases) GetByAuthor(ctx context.Context, author string) ([]domain.Post, error) {

}

func (uc usecases) GetAll() ([]domain.Post, error) {

}

func (uc usecases) Save(ctx context.Context, p *domain.Post) error {

}

func (uc usecases) Delete(ctx context.Context, id int) error {

}

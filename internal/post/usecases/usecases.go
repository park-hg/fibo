package usecases

import (
	"context"
	"fibo/internal/post/domain"
)

type PostUseCases struct {
	postRepo domain.PostRepository
}

func NewPostRepository(pr domain.PostRepository) *PostUseCases {
	return &PostUseCases{postRepo: pr}
}

func (puc PostUseCases) Create(ctx context.Context, p *domain.Post) error {

	return puc.postRepo.Create(ctx, p)
}

func (puc PostUseCases) Get(ctx context.Context, id int) (*domain.Post, error) {

	return puc.postRepo.Get(ctx, id)
}

func (puc PostUseCases) GetByAuthor(ctx context.Context, author string) ([]domain.Post, error) {

	return puc.postRepo.GetByAuthor(ctx, author)
}

func (puc PostUseCases) GetAll() ([]domain.Post, error) {

	return puc.postRepo.GetAll()
}

func (puc PostUseCases) Save(ctx context.Context, p *domain.Post) error {

	return puc.postRepo.Save(ctx, p)
}

func (puc PostUseCases) Delete(ctx context.Context, id int) error {

	return puc.postRepo.Delete(ctx, id)
}

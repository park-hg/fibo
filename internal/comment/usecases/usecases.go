package usecases

import (
	"context"
	"fibo/internal/comment/domain"
)

type CommentUseCases interface {
	Create(ctx context.Context, c *domain.Comment) error
	Get(ctx context.Context, id int) (*domain.Comment, error)
	GetByPost(ctx context.Context, post string) ([]domain.Comment, error)
	Save(ctx context.Context, c *domain.Comment) error
	Delete(ctx context.Context, id int) error
}

type usecases struct {
	commentRepo domain.CommentRepository
}

func (uc usecases) Create(ctx context.Context, c *domain.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) Get(ctx context.Context, id int) (*domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) GetByPost(ctx context.Context, post string) ([]domain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) Save(ctx context.Context, c *domain.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewCommentRepository(cr domain.CommentRepository) CommentUseCases {
	return &usecases{commentRepo: cr}
}

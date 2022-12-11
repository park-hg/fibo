package usecases

import (
	"context"
	"fibo/internal/comment/domain"
)

type CommentUseCases interface {
	Create(ctx context.Context, c domain.Comment) error
	Get(ctx context.Context, commentID uint) (domain.Comment, error)
	GetByPostID(ctx context.Context, postID uint) ([]domain.Comment, error)
	GetByAuthorID(ctx context.Context, AuthorID uint) ([]domain.Comment, error)
	Save(ctx context.Context, commentID uint, c domain.Comment) error
	Delete(ctx context.Context, commentID uint) error
}

type usecases struct {
	commentRepo domain.CommentRepository
}

func (uc usecases) Create(ctx context.Context, c domain.Comment) error {
	return uc.commentRepo.Create(ctx, c)
}

func (uc usecases) Get(ctx context.Context, commentID uint) (domain.Comment, error) {
	return uc.commentRepo.Get(ctx, commentID)
}

func (uc usecases) GetByPostID(ctx context.Context, postID uint) ([]domain.Comment, error) {
	return uc.commentRepo.GetByPostID(ctx, postID)
}

func (uc usecases) GetByAuthorID(ctx context.Context, AuthorID uint) ([]domain.Comment, error) {
	return uc.commentRepo.GetByUserID(ctx, AuthorID)
}

func (uc usecases) Save(ctx context.Context, commentID uint, c domain.Comment) error {
	return uc.commentRepo.Save(ctx, commentID, c)
}

func (uc usecases) Delete(ctx context.Context, commentID uint) error {
	return uc.commentRepo.Delete(ctx, commentID)
}

func NewCommentRepository(cr domain.CommentRepository) CommentUseCases {
	return usecases{commentRepo: cr}
}

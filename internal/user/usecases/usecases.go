package usecases

import (
	"context"
	commentDomain "fibo/internal/comment/domain"
	postDomain "fibo/internal/post/domain"
	"fibo/internal/user/domain"
)

type UserUseCases interface {
	Create(ctx context.Context, u domain.User) error
	GetByName(ctx context.Context, name string) (domain.User, error)
	GetPostsByName(ctx context.Context, name string) ([]postDomain.Post, error)
	GetCommentsByName(ctx context.Context, name string) ([]commentDomain.Comment, error)
	Login(ctx context.Context, id uint) error
	Logout(ctx context.Context, id uint) error
	IsLoggedIn(ctx context.Context, id uint) error
}

type usecases struct {
	userRepo domain.UserRepository
}

func (uc usecases) Create(ctx context.Context, u domain.User) error {
	return uc.userRepo.Create(ctx, u)
}

func (uc usecases) GetByName(ctx context.Context, name string) (domain.User, error) {
	return uc.userRepo.GetByName(ctx, name)
}

func (uc usecases) GetPostsByName(ctx context.Context, name string) ([]postDomain.Post, error) {
	return uc.userRepo.GetPostsByName(ctx, name)
}

func (uc usecases) GetCommentsByName(ctx context.Context, name string) ([]commentDomain.Comment, error) {
	return uc.userRepo.GetCommentsByName(ctx, name)
}

func (uc usecases) Login(ctx context.Context, id uint) error {
	return uc.userRepo.Login(ctx, id)
}

func (uc usecases) Logout(ctx context.Context, id uint) error {
	return uc.userRepo.Logout(ctx, id)
}

func (uc usecases) IsLoggedIn(ctx context.Context, id uint) error {
	return uc.userRepo.IsLoggedIn(ctx, id)
}

func NewUserUseCases(ur domain.UserRepository) UserUseCases {
	return &usecases{userRepo: ur}
}

package usecases

import (
	"context"
	"fibo/internal/user/domain"
)

type UserUseCases interface {
	Create(ctx context.Context, u domain.User) error
	GetByName(ctx context.Context, name string) ([]domain.User, error)
	Login(ctx context.Context, userID uint) error
	Logout(ctx context.Context, userID uint) error
	IsLoggedIn(ctx context.Context, name string) (bool, error)
}

type usecases struct {
	userRepo domain.UserRepository
}

func (uc usecases) Create(ctx context.Context, u domain.User) error {
	return uc.userRepo.Create(ctx, u)
}

func (uc usecases) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	return uc.userRepo.GetByName(ctx, name)
}

func (uc usecases) Login(ctx context.Context, userID uint) error {
	return uc.userRepo.Login(ctx, userID)
}

func (uc usecases) Logout(ctx context.Context, userID uint) error {
	return uc.userRepo.Logout(ctx, userID)
}

func (uc usecases) IsLoggedIn(ctx context.Context, name string) (bool, error) {
	return uc.userRepo.IsLoggedIn(ctx, name)
}

func NewUserUseCases(ur domain.UserRepository) UserUseCases {
	return &usecases{userRepo: ur}
}

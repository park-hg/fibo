package usecases

import (
	"context"
	"fibo/internal/user/domain"
)

type UserUseCases interface {
	Create(ctx context.Context, u *domain.User) error
	Get(ctx context.Context, name string) error
	Login(ctx context.Context, id int) error
	Logout(ctx context.Context, id int) error
	IsLoggedIn(ctx context.Context, id int) error
}

type usecases struct {
	userRepo domain.UserRepository
}

func (uc usecases) Create(ctx context.Context, u *domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) Get(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) Login(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) Logout(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (uc usecases) IsLoggedIn(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewUserUseCases(ur domain.UserRepository) UserUseCases {
	return &usecases{userRepo: ur}
}

package usecases

import (
	"fibo/internal/user/domain"
)

type UserUseCases struct {
	userRepo domain.UserRepository
}

func NewUserRepository(ur domain.UserRepository) *UserUseCases {
	return &UserUseCases{userRepo: ur}
}

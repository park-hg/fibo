package adapter

import (
	"fibo/internal/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	LoggedIn bool
}

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:       &u.ID,
		Name:     u.Name,
		LoggedIn: u.LoggedIn,
	}
}

func ToDBModel(u domain.User) User {
	return User{
		Model:    gorm.Model{},
		Name:     u.Name,
		LoggedIn: u.LoggedIn,
	}
}

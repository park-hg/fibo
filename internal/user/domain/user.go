package domain

import (
	"context"
)

type User struct {
	ID       *uint  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	LoggedIn bool   `json:"logged_in,omitempty"`
	//Posts    []postDomain.Post       `json:"posts,omitempty"`
	//Comments []commentDomain.Comment `json:"comments,omitempty"`
}

type UserRepository interface {
	Create(ctx context.Context, u User) error
	GetByName(ctx context.Context, name string) ([]User, error)
	Login(ctx context.Context, userID uint) error
	Logout(ctx context.Context, userID uint) error
	IsLoggedIn(ctx context.Context, userID uint) (bool, error)
}

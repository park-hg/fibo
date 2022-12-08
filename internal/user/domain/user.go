package domain

import (
	"context"
	commentDomain "fibo/internal/comment/domain"
	postDomain "fibo/internal/post/domain"
)

type User struct {
	Name     string                  `json:"name"`
	LoggedIn bool                    `json:"logged_in"`
	Posts    []postDomain.Post       `json:"posts"`
	Comments []commentDomain.Comment `json:"comments"`
}

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	Get(ctx context.Context, name string) error
	Login(ctx context.Context, id int) error
	Logout(ctx context.Context, id int) error
	IsLoggedIn(ctx context.Context, id int) error
}

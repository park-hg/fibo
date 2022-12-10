package domain

import (
	"context"
	commentDomain "fibo/internal/comment/domain"
	postDomain "fibo/internal/post/domain"
	"fibo/internal/user/adapter"
	"gorm.io/gorm"
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
	GetByName(ctx context.Context, name string) (User, error)
	GetPostsByName(ctx context.Context, name string) ([]postDomain.Post, error)
	GetCommentsByName(ctx context.Context, name string) ([]commentDomain.Comment, error)
	Login(ctx context.Context, id uint) error
	Logout(ctx context.Context, id uint) error
	IsLoggedIn(ctx context.Context, id uint) error
}

func (u *User) ToDBModel() adapter.User {
	//PostDBModels := make([]postAdapter.Post, len(u.Posts))
	//for i, p := range u.Posts {
	//	PostDBModels[i] = p.ToDBModel()
	//}
	//CommentDBModels := make([]commentAdapter.Comment, len(u.Comments))
	//for j, c := range u.Comments {
	//	CommentDBModels[j] = c.ToDBModel()
	//}
	return adapter.User{
		Model:    gorm.Model{},
		Name:     u.Name,
		LoggedIn: u.LoggedIn,
		//Posts:    PostDBModels,
		//Comments: CommentDBModels,
	}
}

package domain

import (
	"context"
	"time"
)

type PostStatus int

const (
	PostTodo PostStatus = iota + 1
	PostInProgress
	PostDone
)

type Post struct {
	ID        uint       `json:"id,omitempty"`
	Title     string     `json:"title"`
	Status    PostStatus `json:"status"`
	Content   string     `json:"content"`
	AuthorID  uint       `json:"author_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type PostRepository interface {
	Create(ctx context.Context, p Post) error
	Get(ctx context.Context, postID uint) (Post, error)
	GetByAuthor(ctx context.Context, authorID uint) ([]Post, error)
	GetAll(ctx context.Context) ([]Post, error)
	Save(ctx context.Context, postID uint, p Post) error
	Delete(ctx context.Context, postID uint) error
}

package domain

import (
	"context"
	"time"
)

type Comment struct {
	ID       uint      `json:"id,omitempty"`
	AuthorID uint      `json:"author_id,omitempty"`
	PostID   uint      `json:"post_id,omitempty"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
}

type CommentRepository interface {
	Create(ctx context.Context, c Comment) error
	Get(ctx context.Context, commentID uint) (Comment, error)
	GetByPostID(ctx context.Context, postID uint) ([]Comment, error)
	GetByUserID(ctx context.Context, userID uint) ([]Comment, error)
	Save(ctx context.Context, commentID uint, c Comment) error
	Delete(ctx context.Context, commentID uint) error
}

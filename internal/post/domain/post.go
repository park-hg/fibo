package domain

import (
	"context"
	"fibo/internal/post/adapter"
	"gorm.io/gorm"
	"time"
)

type PostStatus int

const (
	PostTodo PostStatus = iota + 1
	PostInProgress
	PostDone
)

type Post struct {
	ID       *uint      `json:"id,omitempty"`
	Title    string     `json:"title"`
	Status   PostStatus `json:"status"`
	Content  string     `json:"content"`
	AuthorID uint       `json:"author_id,omitempty"`
	//Comments  []domain.Comment `json:"comments"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type PostRepository interface {
	Create(ctx context.Context, p Post) error
	Get(ctx context.Context, id uint) (Post, error)
	GetByAuthor(ctx context.Context, authorID uint) ([]Post, error)
	GetAll() ([]Post, error)
	Save(ctx context.Context, p Post) error
	Delete(ctx context.Context, id uint) error
}

func (p *Post) ToDBModel() adapter.Post {
	//CommentDBModels := make([]commentAdapter.Comment, len(p.Comments))
	//for i, c := range p.Comments {
	//	CommentDBModels[i] = c.ToDBModel()
	//}
	return adapter.Post{
		Model:   gorm.Model{},
		Title:   p.Title,
		Status:  int(p.Status),
		Content: p.Content,
		UserID:  p.AuthorID,
		//Comments: CommentDBModels,
	}
}

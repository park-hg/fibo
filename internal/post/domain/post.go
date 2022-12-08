package domain

import "context"

type PostStatus int

const (
	PostTodo PostStatus = iota + 1
	PostInProgress
	PostDone
)

type Post struct {
	Title    string     `json:"title"`
	Status   PostStatus `json:"status"`
	Content  string     `json:"content"`
	Author   string     `json:"author"`
	CreateAt string     `json:"create_at"`
}

type PostRepository interface {
	Create(ctx context.Context, p *Post) error
	Get(ctx context.Context, id int) (*Post, error)
	GetByAuthor(ctx context.Context, author string) ([]Post, error)
	GetAll() ([]Post, error)
	Save(ctx context.Context, p *Post) error
	Delete(ctx context.Context, id int) error
}

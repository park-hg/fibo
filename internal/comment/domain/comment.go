package domain

import "context"

type Comment struct {
	Author   string `json:"author"`
	Post     string `json:"post"`
	Content  string `json:"content"`
	CreateAt string `json:"create_at"`
}

type CommentRepository interface {
	Create(ctx context.Context, c *Comment) error
	Get(ctx context.Context, id int) (*Comment, error)
	GetByPost(ctx context.Context, post string) ([]Comment, error)
	Save(ctx context.Context, c *Comment) error
	Delete(ctx context.Context, id int) error
}

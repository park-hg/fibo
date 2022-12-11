package adapter

import (
	"fibo/internal/post/domain"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model

	Title   string
	Status  int
	Content string
	UserID  uint
}

func (p Post) ToDomain() domain.Post {
	return domain.Post{
		ID:        p.ID,
		Title:     p.Title,
		Status:    domain.PostStatus(p.Status),
		Content:   p.Content,
		AuthorID:  p.UserID,
		CreatedAt: &p.CreatedAt,
	}
}
func ToDBModel(p domain.Post) Post {
	return Post{
		Model:   gorm.Model{},
		Title:   p.Title,
		Status:  int(p.Status),
		Content: p.Content,
		UserID:  p.AuthorID,
	}
}

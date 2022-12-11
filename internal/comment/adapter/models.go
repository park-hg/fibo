package adapter

import (
	"fibo/internal/comment/domain"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	Content string
	UserID  uint
	PostID  uint
}

func (c Comment) ToDomain() domain.Comment {
	return domain.Comment{
		ID:       c.ID,
		AuthorID: c.UserID,
		PostID:   c.PostID,
		Content:  c.Content,
		CreateAt: c.CreatedAt,
	}
}
func ToDBModel(c domain.Comment) Comment {
	return Comment{
		Model:   gorm.Model{},
		Content: c.Content,
		UserID:  c.AuthorID,
		PostID:  c.PostID,
	}
}

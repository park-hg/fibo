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
	//Comments []adapter.Comment
}

func (p *Post) ToDomain() domain.Post {
	//CommentDomains := make([]commentDomain.Comment, len(p.Comments))
	//for i, c := range p.Comments {
	//	CommentDomains[i] = c.ToDomain()
	//}
	return domain.Post{
		ID:       &p.ID,
		Title:    p.Title,
		Status:   domain.PostStatus(p.Status),
		Content:  p.Content,
		AuthorID: p.UserID,
		//Comments:  CommentDomains,
		CreatedAt: &p.CreatedAt,
	}
}

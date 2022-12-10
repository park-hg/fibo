package adapter

import (
	"fibo/internal/user/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	LoggedIn bool
	//Posts    []postAdapter.Post `gorm:"foreignKey:AuthorID"`
	//Comments []commentAdapter.Comment
}

func (u *User) ToDomain() domain.User {
	//PostDomains := make([]postDomain.Post, len(u.Posts))
	//for i, p := range u.Posts {
	//	PostDomains[i] = p.ToDomain()
	//}
	//CommentDomains := make([]commentDomain.Comment, len(u.Comments))
	//for j, c := range u.Comments {
	//	CommentDomains[j] = c.ToDomain()
	//}
	return domain.User{
		ID:       &u.ID,
		Name:     u.Name,
		LoggedIn: u.LoggedIn,
	}
}

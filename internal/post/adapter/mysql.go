package adapter

import (
	"context"
	"fibo/internal/post/domain"
	"gorm.io/gorm"
)

type MySQLPostRepository struct {
	db *gorm.DB
}

func (m MySQLPostRepository) Create(ctx context.Context, p domain.Post) error {
	post := ToDBModel(p)
	err := m.db.WithContext(ctx).Create(&post).Error
	if err != nil {
		return err
	}

	return nil
}

func (m MySQLPostRepository) Get(ctx context.Context, postID uint) (domain.Post, error) {
	var post Post
	err := m.db.WithContext(ctx).First(&post, postID).Error
	if err != nil {
		return domain.Post{}, err
	}

	return post.ToDomain(), err
}

func (m MySQLPostRepository) GetByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error) {
	var posts []Post
	err := m.db.WithContext(ctx).Where(&Post{UserID: authorID}).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	postDomains := make([]domain.Post, len(posts))
	for i, p := range posts {
		postDomains[i] = p.ToDomain()
	}

	return postDomains, nil
}

func (m MySQLPostRepository) GetAll(ctx context.Context) ([]domain.Post, error) {
	var posts []Post
	err := m.db.WithContext(ctx).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	postDomains := make([]domain.Post, len(posts))
	for i, p := range posts {
		postDomains[i] = p.ToDomain()
	}

	return postDomains, nil
}

func (m MySQLPostRepository) Save(ctx context.Context, postID uint, p domain.Post) error {
	err := m.db.WithContext(ctx).Where("id = ?", postID).Updates(Post{
		Title:   p.Title,
		Status:  int(p.Status),
		Content: p.Content,
		UserID:  p.AuthorID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MySQLPostRepository) Delete(ctx context.Context, postID uint) error {
	err := m.db.WithContext(ctx).Delete(&Post{}, postID).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMySQLPostRepository(db *gorm.DB) domain.PostRepository {
	return MySQLPostRepository{db: db}
}

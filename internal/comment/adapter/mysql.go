package adapter

import (
	"context"
	"fibo/internal/comment/domain"
	"gorm.io/gorm"
)

type MySQLCommentRepository struct {
	db *gorm.DB
}

func (m MySQLCommentRepository) Create(ctx context.Context, c domain.Comment) error {
	comment := ToDBModel(c)
	err := m.db.WithContext(ctx).Create(&comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MySQLCommentRepository) Get(ctx context.Context, commentID uint) (domain.Comment, error) {
	var comment Comment
	err := m.db.WithContext(ctx).First(&comment, commentID).Error
	if err != nil {
		return domain.Comment{}, err
	}
	return comment.ToDomain(), nil
}

func (m MySQLCommentRepository) GetByPostID(ctx context.Context, postID uint) ([]domain.Comment, error) {
	var comments []Comment
	err := m.db.WithContext(ctx).Where(&Comment{PostID: postID}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	commentDomains := make([]domain.Comment, len(comments))
	for i, c := range comments {
		commentDomains[i] = c.ToDomain()
	}
	return commentDomains, nil
}

func (m MySQLCommentRepository) GetByUserID(ctx context.Context, userID uint) ([]domain.Comment, error) {
	var comments []Comment
	err := m.db.WithContext(ctx).Where(&Comment{UserID: userID}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	commentDomains := make([]domain.Comment, len(comments))
	for i, c := range comments {
		commentDomains[i] = c.ToDomain()
	}
	return commentDomains, nil
}

func (m MySQLCommentRepository) Save(ctx context.Context, commentID uint, c domain.Comment) error {
	err := m.db.WithContext(ctx).Where("id = ?", commentID).Updates(Comment{
		Content: c.Content,
		UserID:  c.AuthorID,
		PostID:  c.PostID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MySQLCommentRepository) Delete(ctx context.Context, commentID uint) error {
	err := m.db.WithContext(ctx).Delete(&Comment{}, commentID).Error
	if err != nil {
		return err
	}
	return nil
}

func NewMySQLCommentRepository(db *gorm.DB) domain.CommentRepository {
	return MySQLCommentRepository{db: db}
}

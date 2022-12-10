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
	//TODO implement me
	panic("implement me")
}

func (m MySQLPostRepository) Get(ctx context.Context, id uint) (domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySQLPostRepository) GetByAuthor(ctx context.Context, authorID uint) ([]domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySQLPostRepository) GetAll() ([]domain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySQLPostRepository) Save(ctx context.Context, p domain.Post) error {
	//TODO implement me
	panic("implement me")
}

func (m MySQLPostRepository) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewMySQLPostRepository(db *gorm.DB) domain.PostRepository {
	return MySQLPostRepository{db: db}
}

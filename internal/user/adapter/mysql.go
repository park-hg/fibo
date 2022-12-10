package adapter

import (
	"context"
	commentDomain "fibo/internal/comment/domain"
	postDomain "fibo/internal/post/domain"
	"fibo/internal/user/domain"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func (m MySQLUserRepository) Create(ctx context.Context, u domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (m MySQLUserRepository) GetByName(ctx context.Context, name string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySQLUserRepository) GetPostsByName(ctx context.Context, name string) ([]postDomain.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySQLUserRepository) GetCommentsByName(ctx context.Context, name string) ([]commentDomain.Comment, error) {
	//TODO implement me
	panic("implement me")
}

func (m MySQLUserRepository) Login(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (m MySQLUserRepository) Logout(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (m MySQLUserRepository) IsLoggedIn(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewMySQLUserRepository(db *gorm.DB) domain.UserRepository {
	return MySQLUserRepository{db: db}
}

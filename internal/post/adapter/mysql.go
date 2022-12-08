package adapter

import (
	"fibo/internal/post/domain"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository() domain.PostRepository {

	return
}
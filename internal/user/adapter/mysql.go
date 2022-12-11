package adapter

import (
	"context"
	"errors"
	"fibo/internal/user/domain"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func (m MySQLUserRepository) Create(ctx context.Context, u domain.User) error {
	user := ToDBModel(u)
	err := m.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MySQLUserRepository) GetByName(ctx context.Context, name string) ([]domain.User, error) {
	var users []User
	err := m.db.WithContext(ctx).Where(&User{Name: name}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	userDomains := make([]domain.User, len(users))
	for i, u := range users {
		userDomains[i] = u.ToDomain()
	}
	return userDomains, nil
}

func (m MySQLUserRepository) Login(ctx context.Context, id uint) error {
	var user User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	if user.LoggedIn {
		return errors.New("already logged in!!")
	}
	err = m.db.WithContext(ctx).Where("id = ?", id).Updates(User{
		LoggedIn: true,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MySQLUserRepository) Logout(ctx context.Context, id uint) error {
	var user User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	if !user.LoggedIn {
		return errors.New("already logged out!!")
	}

	err = m.db.WithContext(ctx).Where("id = ?", id).Updates(User{
		LoggedIn: false,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MySQLUserRepository) IsLoggedIn(ctx context.Context, name string) (bool, error) {
	var user User
	err := m.db.WithContext(ctx).Where("name = ?", name).First(&user).Error
	if err != nil {
		return false, err
	}
	return user.LoggedIn, nil
}

func NewMySQLUserRepository(db *gorm.DB) domain.UserRepository {
	return MySQLUserRepository{db: db}
}

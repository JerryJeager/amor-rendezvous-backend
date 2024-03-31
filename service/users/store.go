package users

import (
	"context"

	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *User) error
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	return &UserRepo{client: client}
}

func (o *UserRepo) CreateUser(ctx context.Context, user *User) error {
	query := o.client.Create(&user).WithContext(ctx)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

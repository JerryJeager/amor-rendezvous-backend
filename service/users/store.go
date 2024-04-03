package users

import (
	"context"
	"fmt"
	"log"

	"github.com/JerryJeager/amor-rendezvous-backend/config"
	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *User) error
	CreateToken(ctx context.Context, userEmail string, user *service.User) (string, error)
}

type UserRepo struct {
	client *gorm.DB
}

func NewUserRepo(client *gorm.DB) *UserRepo {
	fmt.Println("in the gorm client of the store")
	if client == nil {
		fmt.Println("the client is nil in store")
	}
	return &UserRepo{client: client}
}

func (o *UserRepo) CreateUser(ctx context.Context, user *User) error {
	if o.client != nil {
		fmt.Println("gorm client is available in the store...")
	}
	query := config.Session.Create(&user).WithContext(ctx)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (o *UserRepo) CreateToken(ctx context.Context, userEmail string, user *service.User) (string, error) {
	if o.client != nil {
		fmt.Println("gorm client is available in the store...")
	}
	var userModel User

	if err := config.Session.First(&userModel, "email = ?", userEmail).WithContext(ctx).Error; err != nil {
		return "", err
	}

	log.Print(userModel.Email)

	return userModel.Password, nil
}

package users

import (
	"context"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/JerryJeager/amor-rendezvous-backend/utils"
	"github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"
)

type UserSv interface {
	GetUser(ctx context.Context, userID uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, user *User) (string, error)
	CreateToken(ctx context.Context, user *service.User) (string, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (o *UserServ) GetUser(ctx context.Context, userID uuid.UUID) (*User, error) {
	return o.repo.GetUser(ctx, userID)
}

func (o *UserServ) CreateUser(ctx context.Context, user *User) (string, error) {
	id := uuid.New()
	user.ID = id
	if err := user.HashPassword(); err != nil {
		return "", err
	}
	err := o.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func (o *UserServ) CreateToken(ctx context.Context, user *service.User) (string, error) {
	pas, err := o.repo.CreateToken(ctx, user.Email, user)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(user.Password, pas)

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(*user)

	if err != nil {
		return "", err
	}

	return token, nil
}

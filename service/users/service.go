package users

import (
	"context"

	"github.com/JerryJeager/amor-rendezvous-backend/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserSv interface {
	CreateUser(ctx context.Context, user *User) (string, error)
	 CreateToken(ctx context.Context, user *User) (string, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
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

func (o *UserServ) CreateToken(ctx context.Context, user *User) (string, error) {
	pas, err := o.repo.CreateToken(ctx, user)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(pas, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(*user)

	return token, nil
}

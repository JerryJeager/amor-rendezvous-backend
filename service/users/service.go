package users

import (
	"context"

	"github.com/google/uuid"
)

type UserSv interface{
	CreateUser(ctx context.Context, user *User) (string, error)
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}

func (o *UserServ) CreateUser(ctx context.Context, user *User) (string, error){
	id := uuid.New()
	user.ID = id
	err := o.repo.CreateUser(ctx, user)
	if err != nil{
		return "", err
	}
	return id.String(), nil
}

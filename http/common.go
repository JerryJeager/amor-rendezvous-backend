package http

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(ctx *gin.Context) (any, error) {
	cUser, found := ctx.Get("user_id")
	if !found {
		return "", errors.New("get current-user context failed")
	}

	return cUser, nil
}

package http

import (
	"net/http"

	"github.com/JerryJeager/amor-rendezvous-backend/service/users"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController (serv users.UserSv) *UserController{
	return &UserController{serv: serv}
}

func (o *UserController) CreateUser(ctx *gin.Context){
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	id, err := o.serv.CreateUser(ctx, &user)

	if err != nil{
		ctx.Status(http.StatusBadRequest)
		return 
	}

	ctx.JSON(http.StatusCreated, id)
}


func (o *UserController) CreateToken(ctx *gin.Context){
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}

	token, err := o.serv.CreateToken(ctx, &user)

	if err != nil{
		ctx.Status(500)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"TOKEN": token})

}
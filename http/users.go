package http

import (
	"net/http"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/JerryJeager/amor-rendezvous-backend/service/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	serv users.UserSv
}

func NewUserController (serv users.UserSv) *UserController{
	return &UserController{serv: serv}
}

func (o *UserController) GetUser(ctx *gin.Context){
	var pp UserIDPathParm

	if err := ctx.ShouldBindUri(&pp); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "uuid is of invalid format"})
		return
	}

	user, err := o.serv.GetUser(ctx, uuid.MustParse(pp.UserID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user does not exist"})
		return
	}

	ctx.JSON(http.StatusOK, *user)
}

func (o *UserController) CreateUser(ctx *gin.Context){
	var user users.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	id, err := o.serv.CreateUser(ctx, &user)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "user with this email already exists"})
		return 
	}

	ctx.JSON(http.StatusCreated, id)
}


func (o *UserController) CreateToken(ctx *gin.Context){

	var user service.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.Status(http.StatusBadRequest)
		return
	}

	token, err := o.serv.CreateToken(ctx, &user)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "email or password is invaiid"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": token})

}

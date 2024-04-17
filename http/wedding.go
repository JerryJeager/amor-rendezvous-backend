package http

import (
	"net/http"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/JerryJeager/amor-rendezvous-backend/service/wedding"
	"github.com/gin-gonic/gin"
)

type WeddingController struct {
	serv wedding.WeddingSv
}

func NewWeddingController(serv wedding.WeddingSv) *WeddingController {
	return &WeddingController{serv: serv}
}

func (o *WeddingController) CreateWedding(ctx *gin.Context) {
	var wedding service.Wedding

	if err := ctx.ShouldBindJSON(&wedding); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid format"})
	}

	id, err := o.serv.CreateWedding(ctx, &wedding)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusCreated, id)
}

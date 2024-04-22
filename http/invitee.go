package http

import (
	"net/http"

	"github.com/JerryJeager/amor-rendezvous-backend/service/invitees"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InviteeController struct {
	serv invitees.InviteeSv
}

func NewInviteeController(serv invitees.InviteeSv) *InviteeController {
	return &InviteeController{serv: serv}
}

func (o *InviteeController) CreateInvitee(ctx *gin.Context) {
	var invitee invitees.Invitee
	if err := ctx.ShouldBindJSON(&invitee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := o.serv.CreateInvitee(ctx, &invitee)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *InviteeController) GetInvitees (ctx *gin.Context) {
	var pp WeddingIDPathParam
	if err := ctx.ShouldBindUri(&pp); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid uuid format"})
		return
	}

	invitees, err := o.serv.GetInvitees(ctx, uuid.MustParse(pp.WeddingID))

	if err != nil{
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *invitees)
}

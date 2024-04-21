package http

import (
	"net/http"

	"github.com/JerryJeager/amor-rendezvous-backend/service/invitees"
	"github.com/gin-gonic/gin"
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

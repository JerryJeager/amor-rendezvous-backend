package http

import (
	"net/http"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/JerryJeager/amor-rendezvous-backend/service/wedding"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WeddingController struct {
	serv wedding.WeddingSv
}

func NewWeddingController(serv wedding.WeddingSv) *WeddingController {
	return &WeddingController{serv: serv}
}

func (o *WeddingController) GetWedding(ctx *gin.Context) {
	var pp WeddingIDPathParam

	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid format"})
		return
	}

	wedding, err := o.serv.GetWedding(ctx, uuid.MustParse(pp.WeddingID))
	if err != nil {
		ctx.Status(http.StatusNotFound)
	}

	ctx.JSON(http.StatusOK, *wedding)
}

func (o *WeddingController) GetWeddings(ctx *gin.Context) {
	var pp UserIDPathParm
	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid uuid format"})
		return
	}

	Weddings, err := o.serv.GetWeddings(ctx, uuid.MustParse(pp.UserID))

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, *Weddings)
}

func (o *WeddingController) CreateWedding(ctx *gin.Context) {
	var wedding service.Wedding

	if err := ctx.ShouldBindJSON(&wedding); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid format"})
		return
	}

	id, err := o.serv.CreateWedding(ctx, &wedding)

	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *WeddingController) CreateEventType(ctx *gin.Context) {
	var eventType service.EventType

	if err := ctx.ShouldBindJSON(&eventType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid format"})
		return
	}

	var pp WeddingIDPathParam

	if err := ctx.ShouldBindUri(&pp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "id is of invalid format"})
		return
	}

	id, err := o.serv.CreateEventType(ctx, uuid.MustParse(pp.WeddingID), &eventType)

	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (o *WeddingController) UpdateEventType(ctx *gin.Context) {
	var weddingPP WeddingIDPathParam

	if err := ctx.ShouldBindUri(&weddingPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wedding id is of invalid format"})
		return
	}

	var eventTypeID EventTypeIDPathParam

	if err := ctx.ShouldBindUri(&eventTypeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event type id is of invalid format"})
		return
	}

	var eventType service.EventType

	if err := ctx.ShouldBindJSON(&eventType); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := o.serv.UpdateEventType(ctx, uuid.MustParse(weddingPP.WeddingID), uuid.MustParse(eventTypeID.EventTypeID), &eventType)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}

func (o *WeddingController) DeleteEventType(ctx *gin.Context) {
	var weddingPP WeddingIDPathParam

	if err := ctx.ShouldBindUri(&weddingPP); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "wedding id is of invalid format"})
		return
	}

	var eventTypeID EventTypeIDPathParam

	if err := ctx.ShouldBindUri(&eventTypeID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "event type id is of invalid format"})
		return
	}

	err := o.serv.DeleteEventType(ctx, uuid.MustParse(weddingPP.WeddingID), uuid.MustParse(eventTypeID.EventTypeID))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusNoContent)
}

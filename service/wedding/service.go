package wedding

import (
	"context"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/google/uuid"
) 

type WeddingSv interface {
	GetWedding(ctx context.Context, weddingID uuid.UUID) (*Wedding, error)
	CreateWedding(ctx context.Context, wedding *service.Wedding) (string, error)

	CreateEventType(ctx context.Context, weddingID uuid.UUID, eventType *service.EventType) (string, error)
	UpdateEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error)
	DeleteEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID) error
}

type WeddingServ struct {
	repo WeddingStore
}

func NewWeddingService(repo WeddingStore) *WeddingServ {
	return &WeddingServ{repo: repo}
}

func (o *WeddingServ) GetWedding(ctx context.Context, weddingID uuid.UUID) (*Wedding, error) {
	return o.repo.GetWedding(ctx, weddingID)
}

func (o *WeddingServ) CreateWedding(ctx context.Context, wedding *service.Wedding) (string, error) {
	id := uuid.New()

	wedding.ID = id
	if err := o.repo.CreateWedding(ctx, wedding); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *WeddingServ) CreateEventType(ctx context.Context, weddingID uuid.UUID, eventType *service.EventType) (string, error) {
	return o.repo.CreateEventType(ctx, weddingID, eventType)
}

func (o *WeddingServ) UpdateEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error) {
	return o.repo.UpdateEventType(ctx, weddingID, eventTypeID, eventType)
}

func (o *WeddingServ) DeleteEventType(ctx context.Context, weddingID, evenTypeID uuid.UUID) error {
	return o.repo.DeleteEventType(ctx, weddingID, evenTypeID)
}

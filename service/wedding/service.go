package wedding

import (
	"context"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/google/uuid"
)

type WeddingSv interface {
	CreateWedding(ctx context.Context, wedding *service.Wedding) (string, error)
}

type WeddingServ struct {
	repo WeddingStore
}

func NewWeddingService(repo WeddingStore) *WeddingServ {
	return &WeddingServ{repo: repo}
}

func (o *WeddingServ) CreateWedding(ctx context.Context, wedding *service.Wedding) (string, error) {
	id := uuid.New()

	wedding.ID = id
	if err := o.repo.CreateWedding(ctx, wedding); err != nil{
		return "", err
	}

	return id.String(), nil
}
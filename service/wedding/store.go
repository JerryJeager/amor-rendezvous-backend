package wedding

import (
	"context"

	"github.com/JerryJeager/amor-rendezvous-backend/config"
	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"gorm.io/gorm"
)

type WeddingStore interface{
	CreateWedding(ctx context.Context, wedding *service.Wedding) error
}

type WeddingRepo struct {
	client *gorm.DB
}

func NewWeddingRepo(client *gorm.DB) *WeddingRepo{
	return &WeddingRepo{client: client}
}

func (o *WeddingRepo) CreateWedding(ctx context.Context, wedding *service.Wedding) error {
	query := config.Session.Create(&wedding).WithContext(ctx)
	if query.Error != nil {
		return query.Error
	}
	return nil
}
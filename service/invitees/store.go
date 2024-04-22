package invitees

import (
	"context"
	"errors"

	"github.com/JerryJeager/amor-rendezvous-backend/config"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InviteeStore interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) error
	GetInvitees(ctx context.Context, weddingID uuid.UUID) (*Invitees, error)
}

type InviteeRepo struct {
	client *gorm.DB
}

func NewInviteeRepo(client *gorm.DB) *InviteeRepo {
	return &InviteeRepo{client: client}
}

func (o *InviteeRepo) CreateInvitee(ctx context.Context, invitee *Invitee) error {
	result := config.Session.Create(&invitee).WithContext(ctx)

	if result.Error != nil {
		return errors.New("failed to add guest")
	}

	return nil
}

func (o *InviteeRepo) GetInvitees(ctx context.Context, weddingID uuid.UUID) (*Invitees, error) {
	var invitees Invitees

	query := config.Session.WithContext(ctx).Model(Invitee{}).Where("wedding_id = ?", weddingID).Find(&invitees)

	if query.Error != nil {
		return nil, query.Error
	}

	return &invitees, nil
}

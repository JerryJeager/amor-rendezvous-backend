package invitees

import (
	"context"
	"errors"

	"github.com/JerryJeager/amor-rendezvous-backend/config"
	"gorm.io/gorm"
)

type InviteeStore interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) error
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

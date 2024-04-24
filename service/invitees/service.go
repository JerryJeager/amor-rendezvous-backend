package invitees

import (
	"context"

	"github.com/google/uuid"
)

type InviteeSv interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) (string, error)
	GetInvitees(ctx context.Context, weddingID uuid.UUID) (*Invitees, error)
	UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error
	UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error
}

type InviteeServ struct {
	repo InviteeStore
}

func NewInviteeService(repo InviteeStore) *InviteeServ {
	return &InviteeServ{repo: repo}
}

func (o *InviteeServ) CreateInvitee(ctx context.Context, invitee *Invitee) (string, error) {
	id := uuid.New()
	invitee.ID = id
	if err := IsValidStatus(invitee.Status); err != nil {
		return "", err
	}

	if err := o.repo.CreateInvitee(ctx, invitee); err != nil {
		return "", err
	}

	return id.String(), nil
}

func (o *InviteeServ) GetInvitees(ctx context.Context, weddingID uuid.UUID) (*Invitees, error) {
	return o.repo.GetInvitees(ctx, weddingID)
}

func (o *InviteeServ) UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error {
	if err := IsValidStatus(status.Status); err != nil {
		return err
	}
	return o.repo.UpdateInviteeStatus(ctx, inviteeID, status)
}

func (o *InviteeServ) UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error{
	return o.repo.UpdateInvitee(ctx, inviteeID, invitee)
}
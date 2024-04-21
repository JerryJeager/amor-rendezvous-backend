package invitees

import (
	"context"

	"github.com/google/uuid"
)

type InviteeSv interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) (string, error)
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

	if err := o.repo.CreateInvitee(ctx, invitee); err != nil{
		return "", err
	}

	return id.String(), nil
}

package invitees

import (
	"context"
	"os"

	"github.com/google/uuid"
	"gopkg.in/gomail.v2"
)

type InviteeSv interface {
	CreateInvitee(ctx context.Context, invitee *Invitee) (string, error)
	GetInvitees(ctx context.Context, weddingID uuid.UUID) (*Invitees, error)
	UpdateInviteeStatus(ctx context.Context, inviteeID uuid.UUID, status *NewStatus) error
	UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error
	DeleteInvitee(ctx context.Context, inviteeID uuid.UUID) error
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

	if err := sendEmail(invitee); err != nil {
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

func (o *InviteeServ) UpdateInvitee(ctx context.Context, inviteeID uuid.UUID, invitee *Invitee) error {
	return o.repo.UpdateInvitee(ctx, inviteeID, invitee)
}

func (o *InviteeServ) DeleteInvitee(ctx context.Context, inviteeID uuid.UUID) error {
	return o.repo.DeleteInvitee(ctx, inviteeID)
}

func sendEmail(invitee *Invitee) error {
	email := os.Getenv("EMAIL")
	emailUsername := os.Getenv("EMAILUSERNAME")
	emailPassword := os.Getenv("EMAILPASSWORD")
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", invitee.Email)
	m.SetAddressHeader("Cc", invitee.Email, email)
	m.SetHeader("Subject", "Amor Rendezvous")
	m.SetBody("text/html", "Hello <b>You've rsvp for an event</b>  <i>Testing...</i>")

	d := gomail.NewDialer("smtp.gmail.com", 587, emailUsername, emailPassword)

	// Send the email to invitee
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

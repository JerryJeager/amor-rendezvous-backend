package wedding

import (
	"context"

	"github.com/JerryJeager/amor-rendezvous-backend/config"
	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WeddingStore interface {
	GetWedding(ctx context.Context, weddingID uuid.UUID) (*Wedding, error)
	CreateWedding(ctx context.Context, wedding *service.Wedding) error
	GetWeddings(ctx context.Context, userID uuid.UUID) (*Weddings, error)

	CreateEventType(ctx context.Context, weddingID uuid.UUID, eventType *service.EventType) (string, error)
	UpdateEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error)
	DeleteEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID) error
}

type WeddingRepo struct {
	client *gorm.DB
}

func NewWeddingRepo(client *gorm.DB) *WeddingRepo {
	return &WeddingRepo{client: client}
}

func (o *WeddingRepo) GetWedding(ctx context.Context, weddingID uuid.UUID) (*Wedding, error) {
	var wedding Wedding
	query := config.Session.First(&wedding, "id = ?", weddingID).WithContext(ctx)
	if query.Error != nil {
		return &Wedding{}, query.Error
	}
	return &wedding, nil
}

func (o *WeddingRepo) GetWeddings(ctx context.Context, userID uuid.UUID) (*Weddings, error) {
	var Weddings Weddings

	query := config.Session.WithContext(ctx).Model(Wedding{}).Where("user_id = ?", userID).Find(&Weddings)

	if query.Error != nil {
		return nil, query.Error
	}

	return &Weddings, nil
}

func (o *WeddingRepo) CreateWedding(ctx context.Context, wedding *service.Wedding) error {
	query := config.Session.Create(&wedding).WithContext(ctx)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (o *WeddingRepo) CreateEventType(ctx context.Context, weddingID uuid.UUID, eventType *service.EventType) (string, error) {
	id := uuid.NewString()
	path := &[]string{id}
	updateEventType := `UPDATE weddings SET event_types = jsonb_insert(COALESCE(event_types, '{}'), ?, ?) WHERE id = ?`
	result := config.Session.Exec(updateEventType, path, *eventType, weddingID).WithContext(ctx)

	if result.Error != nil{
		return "", result.Error
	}

	return id, nil
}

func (o *WeddingRepo) UpdateEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID, eventType *service.EventType) (string, error) {
	path := &[]string{eventTypeID.String()}

	updateEventType := `UPDATE weddings SET event_types = jsonb_set(event_types, ?, ?, false) WHERE id = ?`
	result := config.Session.Exec(updateEventType, path, *eventType, weddingID).WithContext(ctx)

	if result.Error != nil{
		return "", result.Error
	}

	return eventTypeID.String(), nil
}

func (o *WeddingRepo) DeleteEventType(ctx context.Context, weddingID, eventTypeID uuid.UUID) error {
	udpateEventType := `UPDATE weddings SET event_types = jsonb_delete(event_types, ?) WHERE id = ?`
	result := config.Session.Exec(udpateEventType, eventTypeID, weddingID).WithContext(ctx)

	if result.Error != nil{
		return result.Error
	}

	return nil
}

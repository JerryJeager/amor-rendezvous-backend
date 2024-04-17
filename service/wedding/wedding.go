package wedding

import (
	"time"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/google/uuid"
)

type Wedding struct {
	service.BaseModel
	Description string    `json:"description" binding:"required"`
	EventTypes     
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	Country     string    `json:"country" binding:"required"`
	State       string    `json:"state" binding:"required"`
}

type EventType struct {
	Name        string     `json:"name" binding:"required"`
	Venue       string     `json:"venue"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date" binding:"required"`
}

type EventTypes map[string]EventType


package service

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type User struct {
	BaseModel
	Email     string `json:"email" binding:"required" gorm:"unique"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name" `
	Password  string `json:"password" binding:"required"`
	IsToWed   bool   `json:"is_to_wed"`
}

type Wedding struct {
	BaseModel
	Description string     `json:"description" binding:"required"`
	EventTypes  EventTypes `json:"event_types"`
	UserID      uuid.UUID  `json:"user_id" binding:"required"`
	Country     string     `json:"country" binding:"required"`
	State       string     `json:"state" binding:"required"`
}

type EventType struct {
	Name        string     `json:"name" binding:"required"`
	Venue       string     `json:"venue"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date" binding:"required"`
}

type EventTypes map[string]EventType

func (c EventTypes) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *EventTypes) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		// Unmarshal JSON data into the Values
		return json.Unmarshal(v, c)
	default:
		return fmt.Errorf("unsupported type for Values: %T", v)
	}
}

func (o *Wedding) MarshalJSON() ([]byte, error) {

	wedding := map[string]interface{}{
		"id":          o.ID,
		"description": o.Description,
		"user_id":     o.UserID,
		"country":     o.Country,
		"state":       o.State,
		"event_types": o.EventTypes,
	}

	return json.Marshal(wedding)
}

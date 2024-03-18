package users

import "github.com/JerryJeager/amor-rendezvous-backend/service"

type Users struct {
	service.BaseModel
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	IsToWed   bool   `json:"is_to_wed"`
}

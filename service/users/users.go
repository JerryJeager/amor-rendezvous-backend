package users

import "github.com/JerryJeager/amor-rendezvous-backend/service"

type User struct {
	service.BaseModel
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	IsToWed   bool   `json:"is_to_wed"`
}

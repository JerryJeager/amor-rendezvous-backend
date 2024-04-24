package users

import (
	"html"

	"github.com/JerryJeager/amor-rendezvous-backend/service"

	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	service.BaseModel
	Email     string `json:"email" binding:"required" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password" binding:"required"`
	IsToWed   bool   `json:"is_to_wed"`
}

type UserBasic struct {
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

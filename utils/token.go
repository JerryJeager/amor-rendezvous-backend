package utils

import (
	// "os"
	"strconv"
	"time"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user service.User) (string, error) {

    tokenLifespan, err := strconv.Atoi("24")

    if err != nil {
        return "", err
    }

    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte("A492477RY49RWIOQHGURHGUIQPHERHFUQH734H7473889FHQ89HF"))

}

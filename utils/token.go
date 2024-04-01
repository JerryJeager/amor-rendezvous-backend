package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user service.User) (string, error) {

    tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

    if err != nil {
        return "", err
    }

    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

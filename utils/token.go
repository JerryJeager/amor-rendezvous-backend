package utils

import (
	// "os"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/JerryJeager/amor-rendezvous-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateToken (c *gin.Context) error {
    token, err := GetToken(c)

    if err != nil {
        return err
    }

    _, ok := token.Claims.(jwt.MapClaims)
    if ok && token.Valid {
        return nil
    }

    return errors.New("invalid token provided")
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
    tokenString := getTokenFromRequest(c)
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        return "", nil
    })
    return token, err
}

func getTokenFromRequest(c *gin.Context) string {
    bearerToken := c.Request.Header.Get("Authorization")

    splitToken := strings.Split(bearerToken, " ")
    if len(splitToken) == 2 {
        return splitToken[1]
    }
    return ""
}



func GenerateToken(user service.User) (string, error) {

    tokenLifespan, err := strconv.Atoi("2")

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

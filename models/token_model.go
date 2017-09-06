package models

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/owenso/crypto-portfolio-api/config"
)

//Token : User token
type Token struct {
	Token string `json:"token"`
}

func (t *Token) CreateToken(u User) error {
	configFile, err := config.LoadConfiguration()
	if err != nil {
		return err
	}
	tokenSecret := []byte(configFile.Secret)
	fmt.Println(configFile.Secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["user"] = u
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString(tokenSecret)
	if err != nil {
		fmt.Println("Error while signing the token")
		return err
	}
	t.Token = tokenString

	return nil
}

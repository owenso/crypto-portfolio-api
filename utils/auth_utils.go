package utils

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/owenso/crypto-portfolio-api/config"
)

type UserCtxKeyType string

const UserCtxKey UserCtxKeyType = "user"

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			configFile, err := config.LoadConfiguration()
			if err != nil {
				RespondWithError(w, http.StatusUnauthorized, err.Error())
			}

			secret := []byte(configFile.Secret)

			return secret, nil
		})

	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			strs := claims["user"].(map[string]interface{})
			str1 := strs["id"].(string)
			ctx := r.Context()
			ctx = context.WithValue(ctx, UserCtxKey, str1)
			r = r.WithContext(ctx)
		}

		if token.Valid {
			next(w, r)
		} else {
			RespondWithError(w, http.StatusUnauthorized, "Token is not valid")
		}
	} else {
		fmt.Println(err)
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized access to this resource")
	}

}

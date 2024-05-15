package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hannder92/models"
	"strings"
)

var Email string
var IDUsuario string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	miClave := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("Token format invalid")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		//Rutina que chequea contra la BD
	}

	if !tkn.Valid {
		return &claims, false, "", errors.New("Token invalid")
	}

	return &claims, false, "", err

}

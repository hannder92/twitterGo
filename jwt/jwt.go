package jwt

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hannder92/models"
	"time"
)

func GeneroJWT(ctx context.Context, t models.Usuario) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	miClave := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(miClave)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}

package bd

import (
	"github.com/hannder92/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado {
		return usuario, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usuario.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return usuario, false
	}
	return usuario, true
}

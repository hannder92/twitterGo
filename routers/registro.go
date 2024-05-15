package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hannder92/models"
)

func Registro(ctx context.Context) models.RespAPI {
	var t models.Usuario
	var r models.RespAPI
	r.Status = 400

	fmt.Println("Entre a Registro")
	body := ctx.Value("body").(string)
	if err := json.Unmarshal([]byte(body), &t); err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "Debe especificar el Email"
		fmt.Println(r.Message)
		return r
	}
	if len(t.Password) < 6 {
		r.Message = "Debe especificar una password de al menos 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado con este Email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insetar el registro del usuario"
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	return r
}

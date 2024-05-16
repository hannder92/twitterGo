package routers

import (
	"context"
	"encoding/json"
	"github.com/hannder92/bd"
	"github.com/hannder92/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.Usuario

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Datos Incorrectos " + err.Error()
	}

	status, err := bd.ModificoRegistro(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrio un error al intentar modificar el registro. " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado modificar el registro del usuario."
		return r
	}

	r.Status = 200
	r.Message = "Modificacion de perfil exitosa!"
	return r
}

package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/hannder92/bd"
	"github.com/hannder92/models"
)

func BajaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.PathParameters["id"]
	if len(ID) < 0 {
		r.Message = "El parametro id es obligatorio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID
	status, err := bd.BorroRelacion(t)
	if err != nil {
		r.Message = "Ocurrio un error al intentar borrar la relacion " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado borrar la relacion"
		return r
	}

	r.Status = 200
	r.Message = "Relacion borrada"
	return r
}

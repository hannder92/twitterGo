package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hannder92/bd"
	"github.com/hannder92/models"
)

func ListaUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUsuario := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar el parametro page como un entero mayor que 0" + err.Error()
		return r
	}

	usuarios, status := bd.ListaUsuariosTodos(IDUsuario, int64(pageTemp), search, typeUser)
	if status == false {
		r.Message = "Error al leer los usuarios"
		return r
	}

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Status = 500
		r.Message = "Error al convertir los usuarios a json" + err.Error()
		return r
	}

	r.Message = string(respJson)
	r.Status = 200
	return r

}

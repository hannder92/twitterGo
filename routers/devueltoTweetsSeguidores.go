package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hannder92/bd"
	"github.com/hannder92/models"
)

func LeoTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	IDUsuario := claim.ID.Hex()
	pagina := request.QueryStringParameters["pagina"]

	if len(pagina) < 1 {
		r.Message = "Debe enviar el parametro pagina"
		return r
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar el parametro pagina con un valor mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweetsSeguidores(IDUsuario, pag)
	if correcto == false {
		r.Message = "Error al leer los tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al convertir los tweets de los seguidores a json"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}

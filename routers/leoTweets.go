package routers

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/hannder92/bd"
	"github.com/hannder92/models"
	"strconv"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "El parametero ID es obligatorio"
		return r
	}

	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar el parametro pagina como un valor mayor a 0"
		return r
	}

	tweets, correcto := bd.LeoTweets(ID, int64(pag))
	if !correcto {
		r.Message = "Error al leer los tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos del usuario como JSON"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}

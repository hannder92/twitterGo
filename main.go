package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hannder92/awsgo"
	"github.com/hannder92/bd"
	"github.com/hannder92/handlers"
	"github.com/hannder92/models"
	"github.com/hannder92/secretmanager"
	"net/http"
	"os"
	"strings"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Error en las variables de entorno. Deben incluir 'SecretName', 'BucketName' y 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Error en la lectura de secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	path := strings.Replace(request.PathParameters["twitter"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//Chequeo conexion a la BD o Conecto la BD
	err = bd.ConectarDB(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error conectando a la BD " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respAPI.CustomResp, nil
	}

}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}
	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}
	return traeParametro
}

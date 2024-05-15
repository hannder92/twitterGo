package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Usuario struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Nombre          string             `json:"nombre,omitempty" bson:"nombre"`
	Apellidos       string             `json:"apellidos,omitempty" bson:"apellidos"`
	FechaNacimiento time.Time          `json:"fechaNacimiento,omitempty" bson:"fechaNacimiento"`
	Email           string             `json:"email,omitempty" bson:"email"`
	Password        string             `json:"password,omitempty" bson:"password"`
	Avatar          string             `json:"avatar,omitempty" bson:"avatar"`
	Banner          string             `json:"banner,omitempty" bson:"banner"`
	Biografia       string             `json:"biografia,omitempty" bson:"biografia"`
	Ubicacion       string             `json:"ubicacion,omitempty" bson:"ubicacion"`
	SitioWeb        string             `json:"sitioweb,omitempty" bson:"sitioweb"`
}

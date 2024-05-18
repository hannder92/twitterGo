package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"usuarioId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"usuarioRelacionId,omitempty"`
	Tweet             struct {
		Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   string `bson:"fecha" json:"fecha,omitempty"`
		ID      string `bson:"_id" json:"id,omitempty"`
	}
}

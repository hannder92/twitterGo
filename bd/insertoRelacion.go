package bd

import (
	"context"

	"github.com/hannder92/models"
)

func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()

	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}

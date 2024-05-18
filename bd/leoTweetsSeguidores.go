package bd

import (
	"context"

	"github.com/hannder92/models"
	"go.mongodb.org/mongo-driver/bson"
)

func LeoTweetsSeguidores(ID string, pagina int) ([]*models.DevuelvoTweetsSeguidores, bool) {
	ctx := context.TODO()
	bd := MongoCN.Database(DatabaseName)
	col := bd.Collection("relacion")

	// Suggested code may be subject to a license. Learn more: ~LicenseLog:1750518344.
	skip := (pagina - 1) * 20

	// Suggested code may be subject to a license. Learn more: ~LicenseLog:3746092106.
	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	// Suggested code may be subject to a license. Learn more: ~LicenseLog:4094720978.
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	var result []*models.DevuelvoTweetsSeguidores
	cursor, err := col.Aggregate(ctx, condiciones)
	if err != nil {
		return result, false
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
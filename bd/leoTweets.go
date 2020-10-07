package bd

import (
	"context"
	"log"
	"time"

	"github.com/Gaboper79/gabitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoTweets(ID string, pagina int64) ([]*models.DevuelcoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("gawitterDb")
	col := db.Collection("tweet")

	var resultados []*models.DevuelcoTweets
	condicion := bson.M{
		"userid": ID,
	}
	opciones := options.Find()
	opciones.SetLimit(20)                               // cantidad de elementos por pagina
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) // ordenados por fecha en oreden descenente
	opciones.SetSkip((pagina - 1) * 20)                 // el numero de pagina por 20 tweet

	cursosr, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}
	for cursosr.Next(context.TODO()) {
		var registro models.DevuelcoTweets
		err := cursosr.Decode(&registro)
		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}

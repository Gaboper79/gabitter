package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Gaboper79/gabitter/models"
)

func LeoUsurariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("gawitterDb")
	col := db.Collection("usuarios")

	var results []*models.Usuario
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var encotrado, incluir bool
	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relacion
		r.UsuarioRelacionID = s.ID.Hex() // extraigo el string del id
		r.UsuarioID = ID

		incluir = false

		encotrado, err = ConsultoRelacion(r)
		if tipo == "new" && encotrado == false {
			incluir = true
		}
		if tipo == "follow" && encotrado == true {
			incluir = true
		}

		// si soy el mismo de la relacion
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true {
			s.Password = ""
			s.Biografia = ""
			s.Banner = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Email = ""

			results = append(results, &s)

		}
	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}

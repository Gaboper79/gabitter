package bd

import (
	"context"
	"github.com/Gaboper79/gawitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("gawitterDb")
	col := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjId, _ := result.InsertedID.(primitive.ObjectID) // aca obtengo el id devuelto por mongo al insertar el registro
	return ObjId.String(), true, nil

}

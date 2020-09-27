package bd

import (
	"context"
	"fmt"
	"github.com/Gaboper79/gabitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("gawitterDb")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("registro no encontrado" + err.Error())
		return perfil, err
	}
	return perfil, nil
}

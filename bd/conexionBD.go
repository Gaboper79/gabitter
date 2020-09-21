package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"

	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()

var clienteOptions = options.Client().ApplyURI("mongodb+srv://Gabo1979:Eluney-0803@cluster0.rf6be.mongodb.net/gawitterDb?retryWrites=true&w=majority")

/*ConnectarBD sirve para conectar a la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la DB")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

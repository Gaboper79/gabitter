package main

import (
	"github.com/Gaboper79/gabitter/bd"
	"github.com/Gaboper79/gabitter/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection()==0{
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}

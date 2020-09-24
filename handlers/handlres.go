package handlers

import (
	"github.com/Gaboper79/gabitter/middlew"
	"github.com/Gaboper79/gabitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"log"
	"net/http"
	"os"
)

func Manejadores() {
	router := mux.NewRouter()
	// rutas o endpoint
	router.HandleFunc("/registro", middlew.ChequeBD(routers.Registro)).Methods("POST")
	// pongo en listen el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // maneja los persimos de quien puede acceder.... de momento todos
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

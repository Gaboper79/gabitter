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
	router.HandleFunc("/login", middlew.ChequeBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlew.ChequeBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subiravatar", middlew.ChequeBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obteneravatar", middlew.ChequeBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirbanner", middlew.ChequeBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerbanner", middlew.ChequeBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altarelacion", middlew.ChequeBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/borrarrelacion", middlew.ChequeBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")

	router.HandleFunc("/consultorelacion", middlew.ChequeBD(middlew.ValidoJWT(routers.ConsultoRelacion))).Methods("GET")

	router.HandleFunc("/listausuarios", middlew.ChequeBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/listatweetrelacion", middlew.ChequeBD(middlew.ValidoJWT(routers.LeoTweetSeguidores))).Methods("GET")

	// pongo en listen el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router) // maneja los persimos de quien puede acceder.... de momento todos
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

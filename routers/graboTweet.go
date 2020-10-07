package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Gaboper79/gabitter/bd"

	"github.com/Gaboper79/gabitter/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		// aca que onda??
	}
	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}
	if len(registro.Mensaje) < 2 {
		http.Error(w, "debe ingresar un mensaje", 400)
		return
	}
	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "ocurrio un error al instertar el tweet", 400)
		return
	}
	if status == false {
		http.Error(w, "no se pudo insertar el tweer", 400)
		return
	}
	w.WriteHeader(201)

}

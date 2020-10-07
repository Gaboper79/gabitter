package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Gaboper79/gabitter/bd"
)

func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar un id correcto", 400)
		return
	}
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar la pagina", 400)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // convierto el string enviado en la url a numero de pag
	if err != nil {
		http.Error(w, "No se pudo leer el numero de la pag, debe ser mayor a 0", 400)
		return
	}

	pag := int64(pagina) // convierto la pag a int64
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}

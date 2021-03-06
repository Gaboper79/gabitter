package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Gaboper79/gabitter/bd"
)

func ListaUsuarios(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTempo, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parametro page mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagTempo)
	result, status := bd.LeoUsurariosTodos(IDUsuario, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(result)

}

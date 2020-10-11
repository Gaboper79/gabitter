package routers

import (
	"net/http"

	"github.com/Gaboper79/gabitter/bd"
	"github.com/Gaboper79/gabitter/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", 400)
		return
	}
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "ocurrio un error al borrar la relacion"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se pudo borrar la relacion"+err.Error(), 400)
		return
	}
	w.WriteHeader(201)
}

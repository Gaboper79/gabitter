package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Gaboper79/gabitter/bd"

	"github.com/Gaboper79/gabitter/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario) // idusuario es una variable golbal
	if err != nil {
		http.Error(w, "Error al intentar modificar el registro . reintente nuevamente "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

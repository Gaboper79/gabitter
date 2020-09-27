package routers

import (
	"encoding/json"
	"github.com/Gaboper79/gabitter/bd"
	"github.com/Gaboper79/gabitter/models"
	"net/http"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // el body es string y se lee una vez despues de destruye. y cargo tod en t
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	//aca ya esta en t todos los valores del usuario
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe ser por lo menos de 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con este email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Error al registrar el usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado registrar el usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated) //cargamos en w el status de creado.

}

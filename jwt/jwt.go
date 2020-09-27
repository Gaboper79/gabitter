package jwt

import (
	"github.com/Gaboper79/gabitter/models"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("MastersdelDesarrollo_grupoFacebook")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), // convierte con unix la hora en un numero
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

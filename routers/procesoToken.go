package routers

import (
	"errors"
	"github.com/Gaboper79/gabitter/bd"
	"github.com/Gaboper79/gabitter/models"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var Email string
var IDUsuario string

/*ProcesoToken extrae valores del token*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer") // separa bearer del token
	// sino hay dos bear y el token hay error
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims,encontrado,IDUsuario,nil
	}
	if!tkn.Valid{
		return claims,false,string(""),errors.New("token invalido")
	}
	return claims,false,string(""),err
}

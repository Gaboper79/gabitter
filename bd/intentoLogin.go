package bd

import (
	"github.com/Gaboper79/gabitter/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string,password string) (models.Usuario,bool) {

	usu, encontrado, _ :=ChequeoYaExisteUsuario(email)
	if encontrado==false{
		return usu,false
	}
	passwordByte :=[]byte(password)
	passwordBD:=[]byte(usu.Password)
	//comparo password encriptada y sin encriptas
	err:= bcrypt.CompareHashAndPassword(passwordBD,passwordByte)
	if err!= nil{
		return usu,false

	}
	return usu,true

}
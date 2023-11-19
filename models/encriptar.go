package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// funcion para encriptar password
func EncriptarPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("no se pudo encriptar la contraseña a bytes")
	}
	return string(bytes), err
}

// funcion para controlar password
func VerificarHah(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("las contraseñas no coinciden")
		return false
	}
	return true
}

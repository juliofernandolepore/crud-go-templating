package models

import (
	"gitlab.com/juliofernandolepore/crud-go-templating/config"
)

// funcion para insertar valores en la tabla usuario.

func Createuser(nombre, correo, password string) {
	conexion := config.Conexion() //para abrir la conexion a la base de datos.

	//verificar el nombre de la tabla usuario
	consulta, err := conexion.Prepare("INSERT INTO usuario (nombre, correo, password, activo) VALUES (?,?,?,?)")
	if err != nil {
		panic("error al crear usuario")
	}

	passwordSeguro, err := EncriptarPassword(password) //encriptando password
	if err != nil {
		panic("error al encryptar o cifrar la contraseña")
	}

	consulta.Exec(nombre, correo, passwordSeguro, 1)
	conexion.Close()
}

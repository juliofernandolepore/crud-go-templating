package models

import (
	"gitlab.com/juliofernandolepore/crud-go-templating/config"
)

func Updateuser(id, nombre, correo, password string) {
	conexionDB := config.Conexion()

	consulta, err := conexionDB.Prepare("UPDATE usuario SET nombre = ? , correo = ?, password = ? WHERE id = ?") //query preparada
	if err != nil {
		panic("no se puede actualizar")
	}

	passwordSeguro, err := EncriptarPassword(password)

	if err != nil {
		panic("error al cifrar la contraseña")
	}

	consulta.Exec(nombre, correo, passwordSeguro, id) //id argumento al definirse la funcion
	conexionDB.Close()

}

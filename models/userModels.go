package models

import (
	"log"

	configuracion "github.com/juliofernandolepore/crud-go-templating/config"
)

func Readusers() []Usuario {
	conexion := configuracion.Conn()

	consulta, err := conexion.Query("SELECT id, nombre, correo FROM usuario WHERE activo = 1")
	if err != nil {
		panic("error al leer usuarios")
	}

	conexion.Close()

	var arregloUsuarios []Usuario

	for consulta.Next() {
		var usuarioTemp Usuario
		consulta.Scan(&usuarioTemp.Id, &usuarioTemp.Nombre, &usuarioTemp.Correo)
		arregloUsuarios = append(arregloUsuarios, usuarioTemp)
	}
	return arregloUsuarios
}

func Readuser(id string) Usuario {
	conexion := configuracion.Conn()

	consulta := conexion.QueryRow("SELECT id, nombre, correo, password FROM usuario WHERE activo = 1 AND id = ?", id)

	conexion.Close()

	var usuario Usuario

	consulta.Scan(&usuario.Id, &usuario.Nombre, &usuario.Correo, &usuario.Password)

	return usuario
}

func Createuser(nombre, correo, password string) {
	//para abrir la conexion a la base de datos.
	conexion := configuracion.Conn()
	//verificar el nombre de la tabla usuario
	consulta, err := conexion.Prepare("INSERT INTO usuario (nombre, correo, password, activo) VALUES (?,?,?,?)")
	if err != nil {
		log.Fatalln("error al crear usuario")
	}
	//encriptando password
	passwordSeguro, err := EncriptarPassword(password)
	if err != nil {
		panic("error al encryptar o cifrar la contraseña")
	}

	consulta.Exec(nombre, correo, passwordSeguro, 1)
	conexion.Close()
}

func Deleuser(id string) {
	conexion := configuracion.Conn()
	consulta, err := conexion.Prepare("UPDATE usuario SET activo = ? WHERE ID = ?")
	if err != nil {
		panic("no se pudo preparar la consulta")
	}

	consulta.Exec(0, id)
	conexion.Close()

}

func Updateuser(id, nombre, correo, password string) {
	conexionDB := configuracion.Conn()

	consulta, err := conexionDB.Prepare("UPDATE usuario SET nombre = ? , correo = ?, password = ? WHERE id = ?") //query preparada
	if err != nil {
		panic("no se puede actualizar")
	}

	passwordSeguro, err := EncriptarPassword(password)

	if err != nil {
		panic("error al cifrar la contraseña")
	}

	consulta.Exec(nombre, correo, passwordSeguro, id)
	conexionDB.Close()

}

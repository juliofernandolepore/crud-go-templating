package models

import "gitlab.com/juliofernandolepore/crud-go-templating/config"

func Readuser(id string) Usuario { // regresa un slice de tipo usuario
	conexion := config.Conexion()

	consulta := conexion.QueryRow("SELECT id, nombre, correo, password FROM usuario WHERE activo = 1 AND id = ?", id)

	conexion.Close()

	var usuario Usuario //variable interna que sera el retorno de la funcion

	consulta.Scan(&usuario.Id, &usuario.Nombre, &usuario.Correo, &usuario.Password)

	return usuario
}

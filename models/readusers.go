package models

import (
	"gitlab.com/juliofernandolepore/crud-go-templating/config"
)

func Readusers() []Usuario { // regresa un slice de tipo usuario
	conexion := config.Conexion()

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

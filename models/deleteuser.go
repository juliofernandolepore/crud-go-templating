package models

import (
	"gitlab.com/juliofernandolepore/crud-go-templating/config"
)

func Deleuser(id string) {
	conexionDB := config.Conexion()                                                   //conexion a la base de datos
	consulta, err := conexionDB.Prepare("UPDATE usuario SET activo = ? WHERE ID = ?") //query preparada
	if err != nil {
		panic("no se pudo preparar la consulta")
	}

	consulta.Exec(0, id) //id argumento al definirse la funcion
	conexionDB.Close()

}

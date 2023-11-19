package config

import (
	"database/sql" //libreria estandar de sql
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //driver de mysql, import del paquete completo
)

// DB_CONECTION=root:PassWord123.@(localhost:33060)/basededatos1    va entre llave sino no se conecta a la DB
// esta funcion retorna una conexion a la base de datos mysql
func Conexion() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_CONECTION")) //importo la env con las credenciales de conexion
	if err != nil {
		log.Panic("No se pudo conectar a db de mysql.")
	}
	log.Println("conexion a la DB")
	return db //retorna el objeto db de tipo puntero *sql.DB
}

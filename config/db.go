package configuracion

import (
	"database/sql" //libreria estandar de sql
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //driver de mysql, import del paquete completo
)

// DB_CONECTION=root:PassWord123.@(localhost:33060)/basededatos1
func Conn() *sql.DB {
	DBConn, err := sql.Open("mysql", os.Getenv("DB_CONECTION"))
	if err != nil {
		log.Fatalln("No se pudo conectar a DB.")
	}
	return DBConn //objeto DB puntero *sql.DB
}

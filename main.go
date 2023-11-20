package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/juliofernandolepore/crud-go-templating/routes"
)

const puerto = ":3000"

// cargar variable de entorno
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("error al carga la variable de entorno")
	}
}

func main() {
	//mux object init
	r := routes.Routes()

	log.Println("servidor levantandose en puerto", puerto)

	err := http.ListenAndServe(puerto, r)
	if err != nil {
		log.Fatal("no pudo iniciar el servidor")
	}
}

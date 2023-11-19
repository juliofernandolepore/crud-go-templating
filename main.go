package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gitlab.com/juliofernandolepore/crud-go-templating/routes"
)

const puerto = ":3000"

func init() { //esta funcion para que cargue en el sistema la variable de entorno
	err := godotenv.Load(".env") //libreria de 3ero para cargar varianle de entorno
	if err != nil {
		log.Fatal("error al carga la variable de entorno")
	}
	log.Println("variable de entorno cargada con exito")
}

func main() {

	enrutador := routes.MuxRoutes() //instancia de gorilla mux (objeto r) traida desde el paquete routes

	log.Println("servidor levantandose en puerto", puerto)

	err := http.ListenAndServe(puerto, enrutador)
	if err != nil {
		log.Fatal("no pudo iniciar el servidor")
	}
}

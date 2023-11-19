package controllers

import (
	"log"
	"net/http"
	"text/template"

	"gitlab.com/juliofernandolepore/crud-go-templating/models"
)

func Readusers(w http.ResponseWriter, r *http.Request) {

	vista, err := template.ParseFiles("views/getusers.html", "views/layout.html")
	if err != nil {
		log.Fatal("no se pudo parsear el html, se detiene la aplicacion")
		return
	}

	dbUsuarios := models.Readusers()
	vista.Execute(w, dbUsuarios)
}

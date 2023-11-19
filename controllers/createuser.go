package controllers

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.com/juliofernandolepore/crud-go-templating/models"
)

func Createuser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost { //si el metodo es POST, SE GUARDAR LOS VALORES DEL FORM EN LAS VARIABLES RESPECTIVAS
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		password := r.FormValue("pass1")
		models.Createuser(nombre, correo, password) //metodo que escribe en la base de datos
		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	vista, err := template.ParseFiles("views/createuser.html", "views/layout.html")
	if err != nil {
		log.Fatal("no se pudo parsear el html de la vista, se detiene la aplicacion")
		return
	}

	vista.Execute(w, nil)
}

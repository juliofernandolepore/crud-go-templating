package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/juliofernandolepore/crud-go-templating/models"
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

// controlador sin vista - pura logica de negocio
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.Deleuser(id)
	http.Redirect(w, r, "/users/", http.StatusFound)
}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if r.Method == http.MethodPost {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		password := r.FormValue("pass1")

		models.Updateuser(id, nombre, correo, password) //argumentos del modelo update
		http.Redirect(w, r, "/users/", http.StatusFound)
	}

	vistaUpdate, err := template.ParseFiles("views/updateuser.html", "views/layout.html")
	if err != nil {
		log.Fatal("no se pudieron cargar los template para la vista del update")
	}

	usuario := models.Readuser(id)

	vistaUpdate.Execute(w, usuario)
}

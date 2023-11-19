package controllers

import (
	"html/template"
	"log"
	"net/http"

	"gitlab.com/juliofernandolepore/crud-go-templating/models"
)

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

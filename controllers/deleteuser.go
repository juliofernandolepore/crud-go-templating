package controllers

import (
	"net/http"

	"gitlab.com/juliofernandolepore/crud-go-templating/models"
)

// controlador sin vista - pura logica de negocio
func Deleteuser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.Deleuser(id)
	http.Redirect(w, r, "/users/", http.StatusFound)
}

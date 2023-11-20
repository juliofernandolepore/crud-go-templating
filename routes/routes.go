package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliofernandolepore/crud-go-templating/controllers"
)

func Routes() *mux.Router {

	r := mux.NewRouter() //inicializacion del objeto router (instancia con metodos)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/users/", http.StatusFound)
	})

	userRouter := r.PathPrefix("/users/").Subrouter()

	userRouter.HandleFunc("/", controllers.Readusers).Methods("GET")
	userRouter.HandleFunc("/create", controllers.Createuser).Methods("GET", "POST")
	userRouter.HandleFunc("/delete", controllers.Deleteuser).Methods("GET", "POST")
	userRouter.HandleFunc("/update", controllers.Updateuser).Methods("GET", "POST")

	return r // no olvidar que devuelve un objeto de tipo puntero que contiene todos los endpoints
}

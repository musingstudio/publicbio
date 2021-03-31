package publicbio

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *App) InitRoutes(router *mux.Router) {
	router.HandleFunc("/", app.handler(handleViewProfile))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("../../static/")))
}

func handleViewProfile(app *App, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	username := vars["username"]

	var p *Profile
	if username == "" {
		p = app.singleUser
	}

	if err := renderTemplate(w, "profile", p); err != nil {
		return err
	}
	return nil
}

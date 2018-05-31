package publicbio

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initRoutes(app *app) {
	app.router = mux.NewRouter()

	app.router.HandleFunc("/", app.handler(handleViewProfile))
	app.router.PathPrefix("/").Handler(http.FileServer(http.Dir("../../static/")))
}

func handleViewProfile(app *app, w http.ResponseWriter, r *http.Request) error {
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

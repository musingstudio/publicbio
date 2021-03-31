package publicbio

import (
	"github.com/writeas/impart"
	"log"
	"net/http"
)

type handlerFunc func(app *App, w http.ResponseWriter, r *http.Request) error

func (app *App) handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleError(w, r, func() error {
			return h(app, w, r)
		}())
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		return
	}

	if err, ok := err.(impart.HTTPError); ok {
		log.Printf("Error: %v", err)
		if err.Status >= 300 && err.Status < 400 {
			impart.WriteRedirect(w, err)
			return
		}
		impart.WriteError(w, err)
		return
	}
	log.Printf("Error: %v", err)

	impart.WriteError(w, impart.HTTPError{http.StatusInternalServerError, "We encountered an error we couldn't handle."})
}

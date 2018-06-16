package publicbio

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/writeas/web-core/converter"
	"io/ioutil"
	"log"
	"net/http"
)

type app struct {
	router *mux.Router
	cfg    *config

	singleUser *Profile
}

func (app *app) multiUser() bool {
	return app.singleUser == nil
}

type config struct {
	host string
	port int
}

func Serve() {
	app := &app{
		cfg: &config{},
	}

	flag.IntVar(&app.cfg.port, "p", 8080, "Port to start server on")
	flag.StringVar(&app.cfg.host, "h", "https://public.bio", "Site's base URL")
	var userFile string
	flag.StringVar(&userFile, "u", "", "Configuration file for single-user site")
	flag.Parse()

	if userFile != "" {
		f, err := ioutil.ReadFile(userFile)
		if err != nil {
			log.Fatal("File error: %v\n", err)
		}

		err = json.Unmarshal(f, &app.singleUser)
		if err != nil {
			log.Fatalf("Unable to read user config: %v", err)
		}
		fmt.Printf("Results: %v\n", app.singleUser)
	} else {
		log.Fatal("No user configuration")
	}

	initRoutes(app)

	http.Handle("/", app.router)
	log.Printf("Serving on localhost:%d", app.cfg.port)
	http.ListenAndServe(fmt.Sprintf(":%d", app.cfg.port), nil)
}

func initConverter() {
	formDecoder := schema.NewDecoder()
	formDecoder.RegisterConverter(converter.NullJSONString{}, converter.JSONNullString)
}

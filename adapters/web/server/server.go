package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HmmerHead/go-arquit/adapters/web/handler"
	app "github.com/HmmerHead/go-arquit/application"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type Webserver struct {
	Service app.ProductServiceInterface
}

func MakeNewWebserver() *Webserver {
	return &Webserver{}
}

func (w Webserver) Server() {

	router := mux.NewRouter()
	negroni := negroni.New(negroni.NewLogger())

	handler.MakeProductHandlers(router, negroni, w.Service)
	http.Handle("/", router)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

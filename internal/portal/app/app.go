package app

import (
	"ShortLand/internal/portal"
	"ShortLand/internal/portal/service"
	"ShortLand/internal/portal/service/link"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type app struct {
	router *mux.Router
	api    *portal.API
	srv    services.Service
}

func StartApp() (err error) {
	a := &app{
		router: mux.NewRouter(),
	}
	if err := a.initService(); err != nil {
		return err
	}

	a.api, err = portal.NewAPI(a.router, a.srv)
	if err != nil {
		return err
	}

	log.Println("starting server at :8080")

	return http.ListenAndServe(":8080", a.router)
}

func (a *app) initService() error {
	var srv services.Service
	srv.Link = link.NewLinkService()
	a.srv = srv
	return nil
}

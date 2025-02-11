package app

import (
	"ShortLand/internal/common/db"
	"ShortLand/internal/portal"
	"ShortLand/internal/portal/service"
	"ShortLand/internal/portal/service/link"
	"ShortLand/internal/portal/service/link/typeStorage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type app struct {
	router *mux.Router
	api    *portal.API
	srv    services.Service
	db     *gorm.DB
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

func (a *app) initService() (err error) {
	var srv services.Service
	var linkStor link.StorageLink

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	storType := os.Getenv("STORE_TYPE")
	dbConn := os.Getenv("DB_CONN")

	if dbConn != "" && storType == "inDB" {
		a.db, err = db.Connect(dbConn)
		if err != nil {
			return err
		}

		linkStor = typeStorage.NewInDB(a.db)
	}

	if dbConn == "" || storType == "inMemory" {
		log.Println("using in-memory")
		linkStor = typeStorage.NewInMemory()
	}

	srv.Link = link.NewLinkService(linkStor)
	a.srv = srv
	return nil
}

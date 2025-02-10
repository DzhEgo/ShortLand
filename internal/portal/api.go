package portal

import (
	"ShortLand/internal/model"
	"ShortLand/internal/portal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	router *mux.Router
	srv    services.Service
}

func NewAPI(srv services.Service) (*API, error) {
	ah := &API{
		router: mux.NewRouter(),
		srv:    srv,
	}

	ah.router.HandleFunc("/create", ah.createLink).Methods("POST")
	ah.router.HandleFunc("/get", ah.getOrigLink).Methods("GET")

	return ah, nil
}

func (ah *API) createLink(w http.ResponseWriter, r *http.Request) {
	var cmd model.Command

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	shorLink, err := ah.srv.Link.CreateShortLink(cmd.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(shorLink))
}

func (ah *API) getOrigLink(w http.ResponseWriter, r *http.Request) {
	var cmd model.Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	origLink, err := ah.srv.Link.GetOriginalLink(cmd.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(origLink))
}

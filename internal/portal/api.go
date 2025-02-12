package portal

import (
	"ShortLand/internal/common/model"
	"ShortLand/internal/portal/service"
	"ShortLand/internal/portal/service/link"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type API struct {
	router *mux.Router
	srv    services.Service
}

func NewAPI(router *mux.Router, srv services.Service) (*API, error) {
	ah := &API{
		router: router,
		srv:    srv,
	}

	ah.router.HandleFunc("/create", ah.createLink).Methods("POST")
	ah.router.HandleFunc("/get/{link}", ah.getOrigLink).Methods("GET")

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
		switch {
		case errors.Is(err, link.Exist):
			w.WriteHeader(http.StatusConflict)
		case errors.Is(err, link.Invalid):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(shorLink))
}

func (ah *API) getOrigLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sLink := vars["link"]

	origLink, err := ah.srv.Link.GetOriginalLink(sLink)
	if err != nil {
		switch {
		case errors.Is(err, link.Exist):
			w.WriteHeader(http.StatusGone)
		case errors.Is(err, link.Invalid):
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(origLink))
}

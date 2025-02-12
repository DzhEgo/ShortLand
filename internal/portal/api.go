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
	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	shortLink, err := ah.srv.Link.CreateShortLink(cmd.Link)
	if err != nil {
		var code int

		switch {
		case errors.Is(err, link.Exist):
			code = http.StatusConflict
		case errors.Is(err, link.Invalid):
			code = http.StatusBadRequest
		default:
			code = http.StatusInternalServerError
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(model.ErrorResponse{Code: code, Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(shortLink)
}

func (ah *API) getOrigLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sLink := vars["link"]
	w.Header().Set("Content-Type", "application/json")

	origLink, err := ah.srv.Link.GetOriginalLink(sLink)
	if err != nil {
		var code int

		switch {
		case errors.Is(err, link.Exist):
			code = http.StatusConflict
		case errors.Is(err, link.Invalid):
			code = http.StatusBadRequest
		case errors.Is(err, link.NotFound):
			code = http.StatusNotFound
		case errors.Is(err, link.Expired):
			code = http.StatusGone
		default:
			code = http.StatusInternalServerError
		}

		w.WriteHeader(code)
		json.NewEncoder(w).Encode(model.ErrorResponse{Code: code, Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(origLink)
}

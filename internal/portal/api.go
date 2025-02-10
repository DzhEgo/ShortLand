package portal

import (
	"ShortLand/internal/model"
	"ShortLand/internal/portal/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartAPIServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("/create", CreateLink).Methods("POST")
	r.HandleFunc("/get", GetOrigLink).Methods("GET")

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateLink(w http.ResponseWriter, r *http.Request) {
	var cmd model.Command

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	shorLink, err := service.NewLinkService().CreateShortLink(cmd.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(shorLink))
}

func GetOrigLink(w http.ResponseWriter, r *http.Request) {
	var cmd model.Command
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	origLink, err := service.NewLinkService().GetOriginalLink(cmd.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(origLink))
}

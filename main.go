package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Entity struct {
		Uuid				string	`json:"uuid"`
		Id          string  `json:"id"`
		WasDerivedFrom        string  `json:"wasDerivedFrom"`
    WasGeneratedBy	string `json:"wasGeneratedBy"`
		WasAttributedTo string `json:"wasAttributedTo"`
}

var entities = make(map[string]Entity)

func getBundles(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entities)
}

func addBundle(w http.ResponseWriter, r *http.Request) {
    var entity Entity
    var decoder = json.NewDecoder(r.Body)
		err := decoder.Decode(&entity)
		if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
			panic(err)
		} else {
			entity.Uuid = uuid.New().String()
    	entities[entity.Uuid] = entity
    	w.WriteHeader(http.StatusCreated)
		}
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/prov", getBundles)
    mux.HandleFunc("/prov/put", addBundle)
    http.ListenAndServe(":8080", mux)
}


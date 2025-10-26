package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var bundleCollection = make(map[string]Bundle)

func getProv(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(bundleCollection)
}

func addProv(w http.ResponseWriter, r *http.Request) {
    var bundle Bundle
    var decoder = json.NewDecoder(r.Body)
		err := decoder.Decode(&bundle)
		if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		} else {
			bundle.Uuid = uuid.New().String()
			bundle.InsertTime = time.Now().UTC()
    	bundleCollection[bundle.Uuid] = bundle
    	w.WriteHeader(http.StatusCreated)
		}
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/prov", getProv)
    mux.HandleFunc("/prov/put", addProv)
    http.ListenAndServe(":8080", mux)
}


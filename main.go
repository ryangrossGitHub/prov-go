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
      setBundleDefaults(&bundle)

    	bundleCollection[bundle.Uuid] = bundle
    	w.WriteHeader(http.StatusCreated)
		}
}

func setBundleDefaults(bundle *Bundle) {
	bundle.Uuid = uuid.New().String()
	bundle.InsertTime = time.Now().UTC()
		
	for i := range bundle.EntityCollection {
		bundle.EntityCollection[i].SharedInternal.Uuid = uuid.New().String()
		bundle.EntityCollection[i].SharedInternal.InsertTime = time.Now().UTC()
	}

  for i := range bundle.ActivityCollection {
		bundle.ActivityCollection[i].SharedInternal.Uuid = uuid.New().String()
		bundle.ActivityCollection[i].SharedInternal.InsertTime = time.Now().UTC()
	}
			
	for i := range bundle.AgentCollection {
		bundle.AgentCollection[i].SharedInternal.Uuid = uuid.New().String()
		bundle.AgentCollection[i].SharedInternal.InsertTime = time.Now().UTC()
	}
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/prov", getProv)
    mux.HandleFunc("/prov/put", addProv)
    http.ListenAndServe(":8080", mux)
}


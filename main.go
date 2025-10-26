package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type InternalMetaData struct {
		Uuid				string `json:"uuid"`
		InsertTime	time.Time `json:"insertTime"`
}

type Entity struct {
		InternalMetaData
		Id          		string `json:"id"`
		WasDerivedFrom  string `json:"wasDerivedFrom"`
    WasGeneratedBy	string `json:"wasGeneratedBy"`
		WasAttributedTo string `json:"wasAttributedTo"`
}

var entityCollection = make(map[string]Entity)

func getProv(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entityCollection)
}

func addProv(w http.ResponseWriter, r *http.Request) {
    var entity Entity
    var decoder = json.NewDecoder(r.Body)
		err := decoder.Decode(&entity)
		if err != nil {
    	w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)	
		} else {
			entity.Uuid = uuid.New().String()
			entity.InsertTime = time.Now().UTC()
    	entityCollection[entity.Uuid] = entity
    	w.WriteHeader(http.StatusCreated)
		}
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/prov", getProv)
    mux.HandleFunc("/prov/put", addProv)
    http.ListenAndServe(":8080", mux)
}


package main

import (
	"time"
)

type SharedInternal struct {
	Uuid				string `json:"uuid"` //Internal unique ID for database, can't be set by clients
	InsertTime	time.Time `json:"insertTime"` //Internal time tracking for database, can't be set by clients
}

type ID struct {
	Id          		string `json:"id"` //Client set unique ID (no enforcement)
}

type Entity struct {
	SharedInternal
	ID
	WasDerivedFrom  []string `json:"wasDerivedFrom"`  //Entity
  WasGeneratedBy	[]string `json:"wasGeneratedBy"`  //Activity
	WasAttributedTo []string `json:"wasAttributedTo"` //Agent
}

type Activity struct {
	SharedInternal
	ID
	StartedAtTime 		time.Time `json:"startedAtTime"`
	EndedAtTime 			time.Time `json:"endedAtTime"`
	Used 							[]string `json:"used"`              //Entity
	WasInformedBy 		[]string `json:"wasInformedBy"`     //Activity
	WasAssociatedWith []string `json:"wasAssociatedWith"` //Agent
}

type Agent struct {
	SharedInternal
	ID
	ActedOnBehalfOf	[]string `json:"actedOnBehalfOf"` //Agent
}

type Bundle struct {
  SharedInternal
	EntityCollection   []Entity `json:"entities"`
	ActivityCollection []Activity `json:"activities"`
	AgentCollection    []Agent `json:"agents"`
}

package model

type OpdItem struct {
	ID                    string `rethinkdb:"id,omitempty", json:"id,omitempty"`
	KodeOpd               string `rethinkdb:"kodeOpd", json:"kodeOpd"`
	NamaOpd               string `rethinkdb:"namaOpd", json:"namaOpd"`
	NumPenyedia           int    `rethinkdb:"numPenyedia", json:"numPenyedia"`
	NumPaguPenyedia       int    `rethinkdb:"numPaguPenyedia", json:"numPaguPenyedia"`
	NumSwakelola          int    `rethinkdb:"numSwakelola",json:"numSwakelola"`
	NumPaguSwakelola      int    `rethinkdb:"numPaguSwakelola",json:"numPaguSwakelola"`
	NumPenyediaDlmSwa     int    `rethinkdb:"numPenyediaDlmSwa",json:"numPenyediaDlmSwa"`
	NumPaguPenyediaDlmSwa int    `rethinkdb:"numPaguPenyediaDlmSwa",json:"numPaguPenyediaDlmSwa"`
	TotalPaket            int    `rethinkdb:"totalPaket", json:"totalPaket"`
	TotalPagu             int    `rethinkdb:"totalPagu", json:"totalPagu"`
}

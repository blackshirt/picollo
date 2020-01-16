package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

type RupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}

type RupItem struct {
	ID         string   `rethinkdb:"id,omitempty", json:"id,omitempty"`
	KodeOpd    string   `rethinkdb:"kodeOpd", json:"kodeOpd"`
	NamaOpd    string   `rethinkdb:"namaOpd",json:"namaOpd"`
	KodeRup    string   `rethinkdb:"kodeRup", json:"kodeRup"`
	Kegiatan   *string  `rethinkdb:"kegiatan,omitempty", json:"kegiatan,omitempty"`
	NamaPaket  string   `rethinkdb:"namaPaket", json:"namaPaket"`
	Pagu       string   `rethinkdb:"pagu", json:"pagu"`
	SumberDana string   `rethinkdb:"sumberDana", json:"sumberDana"`
	Waktu      string   `rethinkdb:"waktu", json:"waktu"`
	Tahun      string   `rethinkdb:"tahun", json:"tahun"`
	Kategori   Kategori `rethinkdb:"kategori", json:"kategori"`
	Metode     Metode   `rethinkdb:"metode", json:"metode"`
	State      *State   `rethinkdb:"state", json:"state"`
	Jenis      *Jenis   `rethinkdb:"jenis", json:"jenis"`
	DetilWaktu Waktu    `rethinkdb:"detilWaktu", json:"detilWaktu"`
}

type RupOptions struct {
	KodeOpd  string   `json:"kodeOpd"`
	Kategori Kategori `json:"kategori"`
	Tahun    string   `json:"tahun"`
	Metode   *Metode  `json:"metode"`
	State    *State   `json:"state"`
	Jenis    *Jenis   `json:"jenis"`
}

// // SHA256 checksum of the data
// func calculateHash(o RupItem) (string, error) {
// 	obj, err := json.Marshal(o) // return []byte, error
// 	if err != nil {
// 		return "", err
// 	}
// 	hasher := sha256.New()
// 	hasher.Write(obj)
// 	s := hasher.Sum(nil)
// 	return hex.EncodeToString(s[:]), nil
// }

// this func produces same output with above function
// SHA256 checksum of the data
func checkSum(o []RupItem) (string, error) {
	obj, err := json.Marshal(o) // return []byte, error
	if err != nil {
		return "", err
	}
	s := sha256.Sum256(obj)

	return hex.EncodeToString(s[:]), nil
}

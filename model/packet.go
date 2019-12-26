package model

type PacketItem struct {
	Id            string `rethinkdb:"id,omitempty" json:"id,omitempty"`
	Kode          string `rethinkdb:"kode" json: "kode"`
	Nama          string `rethinkdb:"nama" json: "nama"`
	Instansi      string `rethinkdb:"instansi" json: "instansi"`
	Tahap         string `rethinkdb:"tahap" json: "tahap"`
	HPS           string `rethinkdb:"hps" json: "hps"`
	TahapSekarang string `rethinkdb:"tahap_sekarang" json:"tahap_sekarang"`
}

type PacketResponse struct {
	Draw            string          `json:"draw"`
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	Data            [][]interface{} `json:"data"`
}

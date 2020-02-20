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
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	Data            [][]interface{} `json:"data"`
}

type KategoriLpse string

const (
	PengadaanBarang       KategoriLpse = "PENGADAAN_BARANG"
	PekerjaanKonstruksi   KategoriLpse = "PEKERJAAN_KONSTRUKSI"
	KonsultansiBadanUsaha KategoriLpse = "KONSULTANSI"
	KonsultansiPerorangan KategoriLpse = "KONSULTANSI_PERORANGAN"
	JasaLainnya           KategoriLpse = "JASA_LAINNYA"
)

func (e KategoriLpse) IsValid() bool {
	switch e {
	case PengadaanBarang, PekerjaanKonstruksi, KonsultansiBadanUsaha, KonsultansiPerorangan, JasaLainnya:
		return true
	}
	return false
}

func (e KategoriLpse) String() string {
	return string(e)
}

type MetodeLpse string

const (
	Lelang            MetodeLpse = "lelang"
	PengadaanLangsung MetodeLpse = "pl"
)

func (e MetodeLpse) IsValid() bool {
	switch e {
	case Lelang, PengadaanLangsung:
		return true
	}
	return false
}

func (e MetodeLpse) String() string {
	return string(e)
}

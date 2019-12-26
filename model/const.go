package model

import (
	"fmt"
	"io"
	"strconv"
)

//Model type
type ModelType string

const (
	ModelTypeRup    ModelType = "Rup"
	ModelTypeOpd    ModelType = "Opd"
	ModelTypePacket ModelType = "Packet"
)

var AllModelType = []ModelType{
	ModelTypeRup,
	ModelTypeOpd,
	ModelTypePacket,
}

func (e ModelType) IsValid() bool {
	switch e {
	case ModelTypeRup, ModelTypeOpd, ModelTypePacket:
		return true
	}
	return false
}

func (e ModelType) String() string {
	return string(e)
}

func (e *ModelType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ModelType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ModelType", str)
	}
	return nil
}

func (e ModelType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//Jenis
type Jenis string

const (
	JenisBarang      Jenis = "Barang"
	JenisKonstruksi  Jenis = "Konstruksi"
	JenisKonsultansi Jenis = "Konsultansi"
	JenisJasaLainnya Jenis = "JasaLainnya"
)

var AllJenis = []Jenis{
	JenisBarang,
	JenisKonstruksi,
	JenisKonsultansi,
	JenisJasaLainnya,
}

func (e Jenis) IsValid() bool {
	switch e {
	case JenisBarang, JenisKonstruksi, JenisKonsultansi, JenisJasaLainnya:
		return true
	}
	return false
}

func (e Jenis) String() string {
	return string(e)
}

func (e *Jenis) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Jenis(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Jenis", str)
	}
	return nil
}

func (e Jenis) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//Kategori
type Kategori string

const (
	KategoriPenyedia             Kategori = "Penyedia"
	KategoriSwakelola            Kategori = "Swakelola"
	KategoriPenyediaDlmSwakelola Kategori = "PenyediaDlmSwakelola"
)

var AllKategori = []Kategori{
	KategoriPenyedia,
	KategoriSwakelola,
	KategoriPenyediaDlmSwakelola,
}

func (e Kategori) IsValid() bool {
	switch e {
	case KategoriPenyedia, KategoriSwakelola, KategoriPenyediaDlmSwakelola:
		return true
	}
	return false
}

func (e Kategori) String() string {
	return string(e)
}

func (e *Kategori) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Kategori(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Kategori", str)
	}
	return nil
}

func (e Kategori) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//Metode
type Metode string

const (
	MetodeKontes             Metode = "Kontes"
	MetodeEPurchasing        Metode = "E-Purchasing"
	MetodePengadaanLangsung  Metode = "Pengadaan Langsung"
	MetodePenunjukanLangsung Metode = "Penunjukan Langsung"
	MetodeSwakelola          Metode = "Swakelola"
	MetodeSayembara          Metode = "Sayembara"
	MetodeSeleksi            Metode = "Seleksi"
	MetodeTender             Metode = "Tender"
	MetodeTenderCepat        Metode = "Tender Cepat"
	MetodeDikecualikan       Metode = "Dikecualikan"
)

var AllMetode = []Metode{
	MetodeKontes,
	MetodeEPurchasing,
	MetodePengadaanLangsung,
	MetodePenunjukanLangsung,
	MetodeSwakelola,
	MetodeSayembara,
	MetodeSeleksi,
	MetodeTender,
	MetodeTenderCepat,
	MetodeDikecualikan,
}

func (e Metode) IsValid() bool {
	switch e {
	case MetodeKontes, MetodeEPurchasing,
		MetodePengadaanLangsung, MetodePenunjukanLangsung,
		MetodeSwakelola, MetodeSayembara, MetodeSeleksi,
		MetodeTender, MetodeTenderCepat, MetodeDikecualikan:
		return true
	}
	return false
}

func (e Metode) String() string {
	return string(e)
}

func (e *Metode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Metode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Metode", str)
	}
	return nil
}

func (e Metode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//State
type State string

const (
	StateNotReady   State = "NotReady"
	StateInProgress State = "InProgress"
	StateWarning    State = "Warning"
	StateFinish     State = "Finish"
)

var AllState = []State{
	StateNotReady,
	StateInProgress,
	StateWarning,
	StateFinish,
}

func (e State) IsValid() bool {
	switch e {
	case StateNotReady, StateInProgress, StateWarning, StateFinish:
		return true
	}
	return false
}

func (e State) String() string {
	return string(e)
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = State(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid State", str)
	}
	return nil
}

func (e State) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Role
type Role string

const (
	RoleAdmin Role = "Admin"
	RoleUser  Role = "User"
	RoleGuest Role = "Guest"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
	RoleGuest,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser, RoleGuest:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// Waktu
type Waktu interface {
	IsWaktu()
}

type Planning struct {
	Awal  *string `json:"awal"`
	Akhir *string `json:"akhir"`
}

func (Planning) IsWaktu() {}

type RencanaWaktu struct {
	Pemilihan   *Planning `json:"pemilihan"`
	Pelaksanaan *Planning `json:"pelaksanaan"`
	Pemanfaatan *Planning `json:"pemanfaatan"`
}

func (RencanaWaktu) IsWaktu() {}

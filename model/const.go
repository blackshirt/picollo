package model

import (
	"fmt"
	"io"
	"strconv"
)

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

type Metode string

const (
	MetodeKontes             Metode = "Kontes"
	MetodeEPurchasing        Metode = "EPurchasing"
	MetodePengadaanLangsung  Metode = "PengadaanLangsung"
	MetodePenunjukanLangsung Metode = "PenunjukanLangsung"
	MetodeSwakelola          Metode = "Swakelola"
	MetodeSayembara          Metode = "Sayembara"
	MetodeSeleksi            Metode = "Seleksi"
	MetodeTender             Metode = "Tender"
	MetodeTenderCepat        Metode = "TenderCepat"
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
	case MetodeKontes, MetodeEPurchasing, MetodePengadaanLangsung, MetodePenunjukanLangsung, MetodeSwakelola, MetodeSayembara, MetodeSeleksi, MetodeTender, MetodeTenderCepat, MetodeDikecualikan:
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

type Type string

const (
	TypeRup    Type = "Rup"
	TypeOpd    Type = "Opd"
	TypePacket Type = "Packet"
)

var AllType = []Type{
	TypeRup,
	TypeOpd,
	TypePacket,
}

func (e Type) IsValid() bool {
	switch e {
	case TypeRup, TypeOpd, TypePacket:
		return true
	}
	return false
}

func (e Type) String() string {
	return string(e)
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Type(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

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

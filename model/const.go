package model

import (
	"fmt"
	"io"
	"strconv"
)

// Type represent underlying data model
type Type int

const (
	UnknownType Type = iota // 0
	Rup                     // 1
	Opd                     // 2
	Packet                  // 3
)

func TypeFrom(str string) (Type, error) {
	switch str {
	case "UnknownType":
		return UnknownType, nil
	case "Rup":
		return Rup, nil
	case "Opd":
		return Opd, nil
	case "Packet":
		return Packet, nil
	default:
		return -1, fmt.Errorf("%s is not a valid Type", str)
	}
}

func (e Type) IsValid() bool {
	switch e {
	case UnknownType, Rup, Opd, Packet:
		return true
	}
	return false
}

func (e Type) String() string {
	switch e {
	case UnknownType:
		return "UnknownType"
	case Rup:
		return "Rup"
	case Opd:
		return "Opd"
	case Packet:
		return "Packet"

	default:
		panic("invalid enum Type value")
	}
}

func (e *Type) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	var err error
	*e, err = TypeFrom(str)
	return err
}

func (e Type) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//Jenis represent jenis pengadaan
type Jenis int

const (
	UnknownJenis Jenis = iota // 0
	Barang                    // 1
	Konstruksi                // 2
	Konsultansi               // 3
	JasaLainnya               // 4
)

func (e Jenis) IsValid() bool {
	switch e {
	case UnknownJenis, Barang, Konstruksi, Konsultansi, JasaLainnya:
		return true
	}
	return false
}

func JenisFrom(str string) (Jenis, error) {
	switch str {
	case "UnknownJenis":
		return UnknownJenis, nil
	case "Barang":
		return Barang, nil
	case "Konstruksi":
		return Konstruksi, nil
	case "Konsultansi":
		return Konsultansi, nil
	case "JasaLainnya":
		return JasaLainnya, nil

	default:
		return -1, fmt.Errorf("%s is not a valid Jenis", str)
	}
}

func (e Jenis) String() string {
	switch e {
	case UnknownJenis:
		return "UnknownJenis"
	case Barang:
		return "Barang"
	case Konstruksi:
		return "Konstruksi"
	case Konsultansi:
		return "Konsultansi"
	case JasaLainnya:
		return "JasaLainnya"

	default:
		panic("Invalid enum Jenis value")
	}
}

func (e *Jenis) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums jenis must be strings")
	}

	var err error
	*e, err = JenisFrom(str)
	return err

}

func (e Jenis) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//Kategori represent kategori pengadaan
type Kategori uint

const (
	UnknownKategori              Kategori = iota // 0
	KategoriPenyedia                             // 1
	KategoriSwakelola                            // 2
	KategoriPenyediaDlmSwakelola                 // 3
)

func (e Kategori) IsValid() bool {
	switch e {
	case UnknownKategori, KategoriPenyedia, KategoriSwakelola, KategoriPenyediaDlmSwakelola:
		return true
	}
	return false
}

func KategoriFrom(str string) (Kategori, error) {
	switch str {
	case "UnknownKategori":
		return UnknownKategori, nil
	case "Penyedia":
		return KategoriPenyedia, nil
	case "Swakelola":
		return KategoriSwakelola, nil
	case "PenyediaDlmSwakelola":
		return KategoriPenyediaDlmSwakelola, nil

	default:
		return 0, fmt.Errorf("%s is not a valid Kategori", str)
	}
}

func (e Kategori) String() string {
	switch e {
	case UnknownKategori:
		return "UnknownKategori"
	case KategoriPenyedia:
		return "Penyedia"
	case KategoriSwakelola:
		return "Swakelola"
	case KategoriPenyediaDlmSwakelola:
		return "PenyediaDlmSwakelola"

	default:
		panic("Invalid enum kategori value")
	}
}

func (e *Kategori) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums jenis must be strings")
	}

	var err error
	*e, err = KategoriFrom(str)
	return err
}

func (e Kategori) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

//Metode represent metode pengadaan
type Metode uint

const (
	UnknownMetode      Metode = iota // 0
	PengadaanLangsung                // 1
	Swakelola                        // 2
	Tender                           // 3
	TenderCepat                      // 4
	Dikecualikan                     // 5
	EPurchasing                      // 6
	PenunjukanLangsung               // 7
	Kontes                           // 8
	Sayembara                        // 9
	Seleksi                          // 10
)

func (e Metode) IsValid() bool {
	switch e {
	case UnknownMetode, PengadaanLangsung, Swakelola, Tender, TenderCepat,
		Dikecualikan, EPurchasing, PenunjukanLangsung, Kontes, Sayembara,
		Seleksi:
		return true
	}
	return false
}

func MetodeFrom(str string) (Metode, error) {
	switch str {
	case "UnknownMetode":
		return UnknownMetode, nil
	case "PengadaanLangsung":
		return PengadaanLangsung, nil
	case "Swakelola":
		return Swakelola, nil
	case "Tender":
		return Tender, nil
	case "TenderCepat":
		return TenderCepat, nil
	case "Dikecualikan":
		return Dikecualikan, nil
	case "EPurchasing":
		return EPurchasing, nil
	case "PenunjukanLangsung":
		return PenunjukanLangsung, nil
	case "Kontes":
		return Kontes, nil
	case "Sayembara":
		return Sayembara, nil
	case "Seleksi":
		return Seleksi, nil

	default:
		return 0, fmt.Errorf("%s is not a valid Metode", str)
	}
}

func (e Metode) String() string {
	switch e {
	case UnknownMetode:
		return "UnknownMetode"
	case PengadaanLangsung:
		return "PengadaanLangsung"
	case Swakelola:
		return "Swakelola"
	case Tender:
		return "Tender"
	case TenderCepat:
		return "TenderCepat"
	case Dikecualikan:
		return "Dikecualikan"
	case EPurchasing:
		return "EPurchasing"
	case PenunjukanLangsung:
		return "PenunjukanLangsung"
	case Kontes:
		return "Kontes"
	case Sayembara:
		return "Sayembara"
	case Seleksi:
		return "Seleksi"

	default:
		panic("invalid enum value")
	}
}

func (e *Metode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	var err error
	*e, err = MetodeFrom(str)
	return err
}

func (e Metode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

// State represent state pengadaan in the mean of time
type State uint

const (
	NotReady   State = iota // 0
	InProgress              // 1
	Warning                 // 2
	Finish                  // 3
)

func StateFrom(str string) (State, error) {
	switch str {
	case "NotReady":
		return NotReady, nil
	case "InProgress":
		return InProgress, nil
	case "Warning":
		return Warning, nil
	case "Finish":
		return Finish, nil
	default:
		return 0, fmt.Errorf("%s is not a valid Type", str)
	}
}

func (e State) IsValid() bool {
	switch e {
	case NotReady, InProgress, Warning, Finish:
		return true
	}
	return false
}

func (e State) String() string {
	switch e {
	case NotReady:
		return "NotReady"
	case InProgress:
		return "InProgress"
	case Warning:
		return "Warning"
	case Finish:
		return "Finish"

	default:
		panic("invalid enum value")
	}
}

func (e *State) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	var err error
	*e, err = StateFrom(str)
	return err
}

func (e State) MarshalGQL(w io.Writer) {
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

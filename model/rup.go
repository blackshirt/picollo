package model

import (
	"context"
	"log"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
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

type Storage interface {
	GetRup(ctx context.Context, pkey string) (RupItem, error)      // Get rup item for specific key
	GetOpd(ctx context.Context, pkey string) (RupRekapItem, error) //// Get opd for specific key
	SaveRup(ctx context.Context, objs []RupItem) error
	SaveOpd(ctx context.Context, objs []RupRekapItem) error
	Rup(ctx context.Context, opt RupOptions) ([]RupItem, error)
	Rups(ctx context.Context, opd RupRekapItem) ([]RupItem, error)
	AllRup(ctx context.Context) ([]RupItem, error)
}

// rup item
type rethinkStorage struct {
	session *r.Session
}

func NewStorage(s *r.Session) Storage {
	return &rethinkStorage{session: s}
}

//
func (repo *rethinkStorage) GetRup(ctx context.Context, pkey string) (RupItem, error) {
	var m RupItem
	res, err := r.Table("rup_item").Get(pkey).Run(repo.session, r.RunOpts{Context: ctx})
	res.One(&m)
	return m, err
}

func (repo *rethinkStorage) GetOpd(ctx context.Context, pkey string) (RupRekapItem, error) {
	var m RupRekapItem
	res, err := r.Table("rup_rekap").Get(pkey).Run(repo.session, r.RunOpts{Context: ctx})
	res.One(&m)
	return m, err
}

func (repo *rethinkStorage) SaveRup(ctx context.Context, obj []RupItem) error {
	_, err := r.Table("rup_item").Insert(obj).RunWrite(repo.session, r.RunOpts{Context: ctx})

	return err
}

func (repo *rethinkStorage) SaveOpd(ctx context.Context, obj []RupRekapItem) error {
	_, err := r.Table("rup_rekap").Insert(obj).RunWrite(repo.session, r.RunOpts{Context: ctx})

	return err
}

func (rdb *rethinkStorage) Rups(ctx context.Context, opd RupRekapItem) ([]RupItem, error) {
	rows, err := r.Table("rup_item").GetAllByIndex("kodeOpd", opd.KodeOpd).Run(rdb.session, r.RunOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	var rups []RupItem
	err2 := rows.All(&rups)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err
	}
	return rups, nil

}

func (repo rethinkStorage) AllRup(ctx context.Context) ([]RupItem, error) {
	rows, err := r.Table("rup_item").Run(repo.session, r.RunOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	rups := make([]RupItem, 0)
	err2 := rows.All(&rups)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err
	}
	return rups, nil
}

func (repo *rethinkStorage) Rup(ctx context.Context, opt RupOptions) ([]RupItem, error) {
	rows := r.Table("rup_item").GetAllByIndex("kodeOpd", opt.KodeOpd).Filter(
		r.Row.Field("kategori").Eq(opt.Kategori)).Filter(
		r.Row.Field("tahun").Eq(opt.Tahun))
	// check for nullable option
	if opt.Metode != nil {
		rows = rows.Filter(r.Row.Field("metode").Eq(opt.Metode))
	}
	if opt.State != nil {
		rows = rows.Filter(r.Row.Field("state").Eq(opt.State))
	}
	if opt.Jenis != nil {
		rows = rows.Filter(r.Row.Field("jenis").Eq(opt.Jenis))
	}
	res, err := rows.Run(repo.session, r.RunOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	var rups []RupItem
	err2 := res.All(&rups)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err2
	}
	return rups, nil
}

// service part
type Service interface {
	Get(ctx context.Context, id string) (RupItem, error)
	Rup(ctx context.Context, opt RupOptions) ([]RupItem, error)
}

type rupService struct {
	repo Storage
}

func NewService(s Storage) Service {
	return &rupService{s}
}

// implement the interface

func (s *rupService) Get(ctx context.Context, id string) (RupItem, error) {
	return s.repo.GetRup(ctx, id)
}

func (s *rupService) Rup(ctx context.Context, opt RupOptions) ([]RupItem, error) {
	rups, err := s.repo.Rup(ctx, opt)
	return rups, err
}

type waktu string

func (w waktu) IsWaktu() {}

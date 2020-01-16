package model

import (
	"context"
	"log"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type RupStorage interface {
	Rup(ctx context.Context, id string) (RupItem, error)
	SaveRup(ctx context.Context, objs []RupItem) error
	FilteredRup(ctx context.Context, opt RupOptions) ([]RupItem, error)
	AllRup(ctx context.Context) ([]RupItem, error)
}

type OpdStorage interface {
	Opd(ctx context.Context, id string) (OpdItem, error)
	SaveOpd(ctx context.Context, objs []OpdItem) error
}

type Storage interface {
	RupStorage
	OpdStorage
	RupForOpd(ctx context.Context, opd OpdItem) ([]RupItem, error)
	Exists(ctx context.Context, t Type, key string) bool
}

// rup item
type rethinkStorage struct {
	session *r.Session
}

func NewRethinkStorage(s *r.Session) Storage {
	return &rethinkStorage{session: s}
}

//
func (repo *rethinkStorage) Rup(ctx context.Context, id string) (RupItem, error) {
	var m RupItem
	res, err := r.Table("rup_item").Get(id).Run(repo.session, r.RunOpts{Context: ctx})
	res.One(&m)
	return m, err
}

func (repo *rethinkStorage) Opd(ctx context.Context, pkey string) (OpdItem, error) {
	var m OpdItem
	res, err := r.Table("rup_rekap").Get(pkey).Run(repo.session, r.RunOpts{Context: ctx})
	res.One(&m)
	return m, err
}

func (repo *rethinkStorage) SaveRup(ctx context.Context, obj []RupItem) error {
	_, err := r.Table("rup_item").Insert(obj).RunWrite(repo.session, r.RunOpts{Context: ctx})

	return err
}

func (repo *rethinkStorage) SaveOpd(ctx context.Context, obj []OpdItem) error {
	_, err := r.Table("rup_rekap").Insert(obj).RunWrite(repo.session, r.RunOpts{Context: ctx})

	return err
}

func (rdb *rethinkStorage) RupForOpd(ctx context.Context, opd OpdItem) ([]RupItem, error) {
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

func (repo *rethinkStorage) FilteredRup(ctx context.Context, opt RupOptions) ([]RupItem, error) {
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

func (repo *rethinkStorage) Exists(ctx context.Context, t Type, key string) bool {
	var s bool
	switch t {
	case TypeRup:
		cur, err := r.Table(t.String()).GetAll(key).Count().Eq(1).Run(repo.session, r.RunOpts{Context: ctx})
		if err != nil {
			return false
		}
		if err2 := cur.One(&s); err2 != nil {
			return false
		}
		return s
	case TypeOpd:
		cur, err := r.Table(t.String()).GetAll(key).Count().Eq(1).Run(repo.session, r.RunOpts{Context: ctx})
		if err != nil {
			return false
		}
		if err2 := cur.One(&s); err2 != nil {
			return false
		}
		return s
	case TypePacket:
		cur, err := r.Table(t.String()).GetAll(key).Count().Eq(1).Run(repo.session, r.RunOpts{Context: ctx})
		if err != nil {
			return false
		}
		if err2 := cur.One(&s); err2 != nil {
			return false
		}
		return s
	default:
		return false
	}
}

// service part
type Service interface {
	Get(ctx context.Context, id string) (RupItem, error)
	Rup(ctx context.Context, opt RupOptions) ([]RupItem, error)
	Exists(ctx context.Context, t Type, key string) bool
}

type rupService struct {
	repo Storage
}

func NewService(s Storage) Service {
	return &rupService{s}
}

// implement the interface

func (s *rupService) Get(ctx context.Context, id string) (RupItem, error) {
	return s.repo.Rup(ctx, id)
}

func (s *rupService) Rup(ctx context.Context, opt RupOptions) ([]RupItem, error) {
	rups, err := s.repo.FilteredRup(ctx, opt)
	return rups, err
}

func (s *rupService) Exists(ctx context.Context, t Type, key string) bool {
	return s.Exists(ctx, t, key)
}

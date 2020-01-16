package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"picollo/model"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type Storager interface {
	AvailableType() ([]model.Type, error)
	Load(ctx context.Context, t model.Type, key string) (*Result, error)
	Save(ctx context.Context, obj interface{}) error
	Exists(ctx context.Context, t model.Type, key string) bool
	All(ctx context.Context, t model.Type) (*Result, error)
}

type Result struct {
	tipe model.Type
	data interface{}
}

type rdbStore struct {
	sess *r.Session   // rethinkdb session
	avt  []model.Type // availablye type
}

var defaultAvt []model.Type = []model.Type{model.TypeRup, model.TypeOpd, model.TypePacket}

func NewRdbStore(s *r.Session, avt []model.Type) Storager {
	if avt == nil {
		avt = defaultAvt
	}
	return rdbStore{
		sess: s,
		avt:  avt,
	}
}
func (s rdbStore) hasEmptyAvT() bool {
	if s.avt == nil || len(s.avt) == 0 {
		return true
	}
	return false
}

func (s rdbStore) AvailableType() ([]model.Type, error) {
	if s.hasEmptyAvT() {
		return nil, errors.New("has no AvT")
	}
	return s.avt, nil
}

func (s rdbStore) containsType(t model.Type) bool {
	if s.hasEmptyAvT() {
		return false
	}
	for _, item := range s.avt {
		if item == t {
			return true
		}
	}
	return false
}

func (s rdbStore) Load(ctx context.Context, t model.Type, key string) (*Result, error) {
	if !s.containsType(t) {
		return nil, errors.New("your tipe was not in store")
	}
	res := &Result{tipe: t}

	cur, err := r.Table(t.String()).Get(key).Run(s.sess, r.RunOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}
	switch t {
	case model.TypeRup:
		var item model.RupItem
		if err := cur.One(&item); err != nil {
			log.Fatal(err)
		}
		res.data = item
	case model.TypeOpd:
		item := model.OpdItem{}
		err := cur.One(&item)
		if err != nil {
			log.Fatal(err)
		}
		res.data = item

	case model.TypePacket:
		item := model.PacketItem{}

		if err := cur.One(&item); err != nil {
			log.Fatal(err)
		}
		res.data = item

	}
	return res, err
}

func (s rdbStore) All(ctx context.Context, t model.Type) (*Result, error) {
	if !s.containsType(t) {
		return nil, errors.New("your tipe was not in store")
	}
	res := &Result{tipe: t}
	rows, err := r.Table(t.String()).Run(s.sess, r.RunOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
	}

	switch res.tipe {
	case model.TypeRup:
		items := make([]model.RupItem, 0)
		if err := rows.All(&items); err != nil {
			log.Fatal(err)
		}
		res.data = items
	case model.TypeOpd:
		items := make([]model.OpdItem, 0)
		if err := rows.All(&items); err != nil {
			log.Fatal(err)
		}
		res.data = items
	case model.TypePacket:
		items := make([]model.PacketItem, 0)
		if err := rows.All(&items); err != nil {
			log.Fatal(err)
		}
		res.data = items
	}

	// filling the data with appropriate model based on tipe
	// rows, err := r.Table(t.String()).Run(s.sess, r.RunOpts{Context: ctx})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return res, err
}

func (s rdbStore) Exists(ctx context.Context, t model.Type, key string) bool {
	_, err := r.Table(t.String()).GetAll(key).Count().Eq(1).Run(s.sess, r.RunOpts{Context: ctx})
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (s rdbStore) Save(ctx context.Context, obj interface{}) error {
	switch obj.(type) {
	case *model.RupItem, []*model.RupItem:
		tbl := model.TypeRup.String()
		_, err := r.Table(tbl).Insert(obj).RunWrite(s.sess, r.RunOpts{Context: ctx})

		return err
	case *model.OpdItem, []*model.OpdItem:
		tbl := model.TypeOpd.String()
		_, err := r.Table(tbl).Insert(obj).RunWrite(s.sess, r.RunOpts{Context: ctx})

		return err
	case *model.PacketItem, []*model.PacketItem:
		tbl := model.TypePacket.String()
		_, err := r.Table(tbl).Insert(obj).RunWrite(s.sess, r.RunOpts{Context: ctx})

		return err
	default:
		return errors.New("unknown obj to insert")
	}
}

func main() {
	var ctx context.Context = nil
	s, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "picollo",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	repo := NewRdbStore(s, defaultAvt)

	// res, err := repo.AvailableType()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, item := range res {
	// 	fmt.Println(item)
	// }
	// load
	// rup, err := repo.Load(ctx, model.TypeRup, "020708eb-09b1-44de-a880-9dc769224b41")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(rup.tipe, rup.data)

	// fmt.Println("assertion")
	// val, ok := rup.data.(model.RupItem)
	// if ok {
	// 	fmt.Println(val.ID, val.NamaPaket)
	// }

	fmt.Println("All")
	res, err := repo.All(ctx, model.TypeRup)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res.tipe)
	val, ok := res.data.([]model.RupItem)
	if ok {
		for _, item := range val {
			fmt.Println(item.ID, item.NamaPaket)
		}
	}
}

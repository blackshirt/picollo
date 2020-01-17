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
	Tipe model.Type
	Data interface{}
}

type rdbStore struct {
	sess *r.Session   // rethinkdb session
	avt  []model.Type // availablye type
}

var defaultAvt []model.Type = []model.Type{model.TypeRup, model.TypeOpd, model.TypePacket}

// NewRdbStore creates RethinkDB storage
func NewRdbStore(s *r.Session, avt []model.Type) Storager {
	if avt == nil {
		avt = defaultAvt
	}
	return &rdbStore{
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

func loadRup(c *r.Cursor, key string) (*model.RupItem, error) {
	if c == nil {
		return nil, errors.New("No cursor provided")
	}
	rup := &model.RupItem{}
	if err := c.One(&rup); err != nil {
		return nil, errors.New("erros in rup cur one")
	}
	return rup, nil

}

func allRup(c *r.Cursor) ([]model.RupItem, error) {
	if c == nil {
		return nil, errors.New("No cursor provided")
	}
	rups := make([]model.RupItem, 0)
	if err := c.All(&rups); err != nil {
		return nil, errors.New("erros in rups cur all")
	}
	return rups, nil

}

func loadOpd(c *r.Cursor, key string) (*model.OpdItem, error) {
	if c == nil {
		return nil, errors.New("No cursor provided")
	}
	opd := &model.OpdItem{}
	if err := c.One(&opd); err != nil {
		return nil, errors.New("erros in opd cur one")
	}
	return opd, nil

}

func allOpd(c *r.Cursor) ([]model.OpdItem, error) {
	if c == nil {
		return nil, errors.New("No cursor provided")
	}
	opds := make([]model.OpdItem, 0)
	if err := c.All(&opds); err != nil {
		return nil, errors.New("erros in opds cur all")
	}
	return opds, nil

}

func loadPacket(c *r.Cursor, key string) (*model.PacketItem, error) {
	if c == nil {
		return nil, errors.New("No cursor provided")
	}
	pck := &model.PacketItem{}
	if err := c.One(&pck); err != nil {
		return nil, errors.New("erros in opd pck one")
	}
	return pck, nil

}

func allPacket(c *r.Cursor) ([]model.PacketItem, error) {
	if c == nil {
		return nil, errors.New("No cursor provided")
	}

	pcks := make([]model.PacketItem, 0)
	if err := c.All(&pcks); err != nil {
		return nil, errors.New("erros in opds cur all")
	}

	return pcks, nil
}

func loadItem(c *r.Cursor, t model.Type, key string) (*Result, error) {
	if !t.IsValid() {
		return nil, errors.New("errors tipe was not valid")
	}
	res := &Result{Tipe: t}
	switch t {
	case model.TypeRup:
		rup, err := loadRup(c, key)
		if err != nil {
			return nil, err
		}
		res.Data = rup
	case model.TypeOpd:
		opd, err := loadOpd(c, key)
		if err != nil {
			return nil, err
		}
		res.Data = opd
	case model.TypePacket:
		pck, err := loadPacket(c, key)
		if err != nil {
			return nil, err
		}
		res.Data = pck
	}
	return res, nil
}

func allItem(c *r.Cursor, t model.Type) (*Result, error) {
	if !t.IsValid() {
		return nil, errors.New("errors tipe was not valid")
	}

	res := &Result{Tipe: t}

	switch t {
	case model.TypeRup:
		rups, err := allRup(c)
		if err != nil {
			return nil, err
		}
		res.Data = rups

	case model.TypeOpd:
		opds, err := allOpd(c)
		if err != nil {
			return nil, err
		}
		res.Data = opds

	case model.TypePacket:
		pcks, err := allPacket(c)
		if err != nil {
			return nil, err
		}
		res.Data = pcks
	}

	return res, nil
}

func (s rdbStore) Load(ctx context.Context, t model.Type, key string) (*Result, error) {
	if !s.containsType(t) && !t.IsValid() {
		return nil, errors.New("your tipe was not in store")
	}

	cur, err := r.Table(t.String()).Get(key).Run(s.sess, r.RunOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	res, err := loadItem(cur, t, key)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s rdbStore) All(ctx context.Context, t model.Type) (*Result, error) {
	if !s.containsType(t) && !t.IsValid() {
		return nil, errors.New("your tipe was not in store")
	}
	cur, err := r.Table(t.String()).Run(s.sess, r.RunOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	res, err := allItem(cur, t)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s rdbStore) Exists(ctx context.Context, t model.Type, key string) bool {
	if !s.containsType(t) && !t.IsValid() {
		return false
	}
	_, err := r.Table(t.String()).GetAll(key).Count().Eq(1).Run(s.sess, r.RunOpts{Context: ctx})
	if err == nil {
		return true
	}

	return false
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
	// rup, err := repo.Load(ctx, model.TypeRup, "0fefa19a-d592-4240-8866-f897d580fb9e")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(rup.Tipe, rup.Data)

	// fmt.Println("assertion")
	// val, ok := rup.Data.(*model.RupItem)
	// if ok {
	// 	fmt.Println(val.ID, val.NamaPaket)
	// }

	fmt.Println("All")
	res, err := repo.All(ctx, model.TypeRup)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res.Tipe)
	val, ok := res.Data.([]model.RupItem)
	if ok {
		for _, item := range val {
			fmt.Println(item.ID, item.KodeRup, item.NamaPaket)
		}
	}
}

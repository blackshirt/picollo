package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"picollo/model"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

var (
	ErrNoCursorProvided      = errors.New("Error No cursor provided")
	ErrTypeNoExist  = errors.New("Error your tipe was not exist in store")
	ErrUnknownObject = errors.New("Error unknown obj to insert")
	ErrInQuery       = errors.New("Error in cur one/all")
	ErrInvalidType   = errors.New("Error invalid tipe")
	ErrNoAVT         = errors.New("Error has no AvT")
)

type Storager interface {
	// what type the Storager provides
	AvailableType() ([]model.Type, error)

	// Load load the model with type t and key ke  from underlying backend implementation
	Load(ctx context.Context, t model.Type, key string) (*Result, error)

	// Save saving the obj data to the backend implementation
	Save(ctx context.Context, obj interface{}) error

	// Exists check whether model with tipe t and key key was there in backend
	Exists(ctx context.Context, t model.Type, key string) bool

	// All return all model data available in the backend
	All(ctx context.Context, t model.Type) (*Result, error)
}

// Result of the operation
type Result struct {
	// tipe the result provided
	Tipe model.Type
	// underlying data
	Data interface{}
}

// rethinkdb Storager implementation
type rdbStore struct {
	sess *r.Session   // rethinkdb session
	avt  []model.Type // availablye type
}

var defaultAvt []model.Type = model.AllType

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

// AvailableType implement AvailableType interface method of Storager
func (s rdbStore) AvailableType() ([]model.Type, error) {
	if s.hasEmptyAvT() {
		return nil, ErrNoAVT
	}
	return s.avt, nil
}

// Load implement the Load method of the Storager interface, load single item
// based on tipe `t` and id `key`
func (s rdbStore) Load(ctx context.Context, t model.Type, key string) (*Result, error) {
	// check if tipe `t` was valid tipe and available in the store
	if !s.containsType(t) && !t.IsValid() {
		return nil, ErrTypeNoExist
	}

	// run the get query
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

// All return all item based on tipe `t`
func (s rdbStore) All(ctx context.Context, t model.Type) (*Result, error) {
	if !s.containsType(t) && !t.IsValid() {
		return nil, ErrTypeNoExist
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

// Exists check whether item tipe `t` and id `key` was exist in the store
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

// Save saving the object `obj` to the backend storage
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
		return ErrUnknownObject
	}
}

// hasEmptyAvT check whether the storage has available tipe
func (s rdbStore) hasEmptyAvT() bool {
	if s.avt == nil || len(s.avt) == 0 {
		return true
	}
	return false
}

// containsType check the store contains tipe `t`
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

// loadRup load single rupitem with id `key` using provided rethinkdb cursor `c`
func loadRup(c *r.Cursor, key string) (*model.RupItem, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	rup := &model.RupItem{}
	if err := c.One(&rup); err != nil {
		return nil, ErrInQuery
	}
	return rup, nil

}

// allRup load all rupitem with using provided rethinkdb cursor c
func allRup(c *r.Cursor) ([]model.RupItem, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	rups := make([]model.RupItem, 0)
	if err := c.All(&rups); err != nil {
		return nil, ErrInQuery
	}
	return rups, nil

}

// loadOpd load single opditem with id key using provided rethinkdb cursor c
func loadOpd(c *r.Cursor, key string) (*model.OpdItem, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	opd := &model.OpdItem{}
	if err := c.One(&opd); err != nil {
		return nil, ErrInQuery
	}
	return opd, nil

}

// allOpd load all opditem using provided rethinkdb cursor `c`
func allOpd(c *r.Cursor) ([]model.OpdItem, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	opds := make([]model.OpdItem, 0)
	if err := c.All(&opds); err != nil {
		return nil, ErrInQuery
	}
	return opds, nil

}

// loadPacket load single packetitem with id `key` using provided rethinkdb cursor `c`
func loadPacket(c *r.Cursor, key string) (*model.PacketItem, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	pck := &model.PacketItem{}
	if err := c.One(&pck); err != nil {
		return nil, ErrInQuery
	}
	return pck, nil

}

// allPacket load all packetitem using provided rethinkdb cursor `c`
func allPacket(c *r.Cursor) ([]model.PacketItem, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}

	pcks := make([]model.PacketItem, 0)
	if err := c.All(&pcks); err != nil {
		return nil, ErrInQuery
	}

	return pcks, nil
}

// loadItem load single specific item based on tipe `t` and id `key` using
// provided rethinkdb Cursor `c`
func loadItem(c *r.Cursor, t model.Type, key string) (*Result, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	if !t.IsValid() {
		return nil, ErrInvalidType
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

// allItem load all item with specific tipe `t` using provided rethinkdb cursor `c`
func allItem(c *r.Cursor, t model.Type) (*Result, error) {
	if c == nil {
		return nil, ErrNoCursorProvided
	}
	if !t.IsValid() {
		return nil, ErrInvalidType
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

///////////////////////////////////
// Service related definition /////
///////////////////////////////////
type Service interface {
}

type serviceManager struct {
	store *Storager

	mu  sync.RWMutex
	avS map[string]interface{}
}

type RupService interface {
	Get(ctx context.Context, key string) (model.RupItem, error)
	GetAll(ctx context.Context) ([]model.RupItem, error)
	FilteredRup(ctx context.Context, cond model.RupFilter) ([]model.RupItem, error)
	Stats(obj []model.RupItem, cond model.RupFilter) (*StatResults, error)
}

type StatResults struct {
	Obj    []interface{} // data he operates om
	Length int           // length of the data
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

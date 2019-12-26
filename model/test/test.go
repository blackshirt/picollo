package main

import (
	"context"
	"fmt"
	"log"
	"picollo/model"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

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
	repo := model.NewStorage(s)
	service := model.NewService(repo)

	opt := model.RupOptions{
		Tahun:    "2020",
		KodeOpd:  "63408", //smpn 1 mirit
		Kategori: model.KategoriPengadaanPenyediaDlmSwakelola,
		Metode:   nil,
		Jenis:    nil,
		State:    nil,
	}
	// ro := model.RupRekapItem{}
	// ro.KodeOpd = "63768"
	rups, err := service.Rup(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range rups {
		fmt.Println(item)
	}
}

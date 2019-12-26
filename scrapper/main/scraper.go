package main

import (
	"context"
	"fmt"
	"log"
	"picollo/model"
	"picollo/scrapper"
	"time"

	"github.com/gocolly/colly"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

func main() {
	col := colly.NewCollector(
		colly.AllowedDomains("sirup.lkpp.go.id"),
	)
	var ctx context.Context = nil

	col.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
	})
	s, err := r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "picollo",
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	repo := model.NewStorage(s)
	fetcher := scrapper.NewRupFetcher(col, repo)

	opt := model.RupOptions{
		Tahun:    "2020",
		KodeOpd:  "159844", //kelautan
		Kategori: model.KategoriPenyedia,
		Metode:   nil,
		Jenis:    nil,
	}
	fmt.Println(opt.Tahun, opt.KodeOpd, opt.Kategori, opt.Metode, opt.Jenis)
	fmt.Println("url:", scrapper.BuildRupURL(opt).String())
	link, err := scrapper.LinkRup(opt)
	fmt.Println("link:", link, err)

	res, err := fetcher.Fetch(ctx, opt)
	if err != nil {
		log.Fatal("Error in", err)
	}
	for _, item := range res {
		fmt.Println(item)
	}
	fetcher.Save(ctx, res)

}

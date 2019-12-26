package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"picollo/app/opd"

	"github.com/gocolly/colly"
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type client struct {
	collector  *colly.Collector
	opdService *opd.Service
}

type RupItem struct {
	Id      string `rethinkdb:"id,omitempty", json:"id,omitempty"`
	Rup     string `rethinkdb:"rup", json:"rup"`
	Nama    string `rethinkdb:"nama", json:"nama"`
	Pagu    string `rethinkdb:"pagu", json:"pagu"`
	Metode  string `rethinkdb:"metode", json:"metode"`
	Dana    string `rethinkdb:"dana", json:"dana"`
	KodeRup string `rethinkdb:"kode_rup", json:"kode_rup"`
	Waktu   string `rethinkdb:"waktu", json:"waktu"`
}

type RupS []RupItem
type RupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("sirup.lkpp.go.id"),
	)
	c.Limit(&colly.LimitRule{
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
	repo := opd.NewRepository(s)
	service := opd.NewService(repo)
	rc := client{
		collector:  c,
		opdService: &service,
	}
	data, err := rc.fetchRekapRup()
	if err != nil {
		log.Fatal(err)
	}
	var opds = make([]*opd.Opd, 0)
	for _, item := range data.AaData {
		//fmt.Println(item[0], item[1], item[2], item[3], item[4])
		opd := new(opd.Opd)
		opd.Kode = item[0]
		opd.Nama = item[1]
		opds = append(opds, opd)
	}

	res, err := r.Table("opd").Insert(opds, r.InsertOpts{Conflict: "replace"}).RunWrite(s)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res.Inserted)
	defer s.Close()

}

func (r client) fetchRupOpd(url string) (*RupResponse, error) {
	resp := new(RupResponse)
	r.collector.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &resp)
		if err != nil {
			log.Fatalln(err)
			return
		}
	})
	r.collector.Visit(url)
	r.collector.Wait()
	return resp, nil
}

func addQuerytoURL(idSatker, tahun string) *url.URL {
	basePath := "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker"
	u, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("tahun", tahun)
	q.Add("idSatker", idSatker)
	u.RawQuery = q.Encode()
	return u
}

func (r client) fetchRekapRup() (*RupResponse, error) {
	baseURL := "https://sirup.lkpp.go.id/sirup/datatablectr/datatableruprekapkldi?idKldi=D128&tahun=2019"

	resp := new(RupResponse)
	r.collector.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &resp)
		if err != nil {
			log.Fatalln(err)
		}
	})
	r.collector.Visit(baseURL)
	r.collector.Wait()
	return resp, nil
}

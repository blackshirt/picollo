// main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gocolly/colly"
)

var URL_RUP_PEROPD_PENYEDIA string = "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker?tahun=2019&idSatker=63401"
var URL_RUP_PEROPD_SWAKELOLA string = "https://sirup.lkpp.go.id/sirup/datatablectr/datarupswakelolasatker?tahun=2019&idSatker=63429"
var URL_RUP_PEROPD_PENYEDIA_DLM_SWA string = "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediaswakelolaallrekap?tahun=2019&idSatker=63429"
var URL_DETAIL_RUP string = "https://sirup.lkpp.go.id/sirup/home/detailPaketPenyediaPublic2017/20748967"

var URL_RUP_KBM string = "https://sirup.lkpp.go.id/sirup/datatablectr/datatableruprekapkldi?idKldi=D128&tahun=2019"
var REKAP_LPSE_SATKER_URL string = "http://lpse.kebumenkab.go.id/eproc4/dt/lelang?draw=2&authenticityToken=1ff307a1d53e64aa12109faac7337eb469523994"

//var URL_RUP []string{URL_RUP_PEROPD_PENYEDIA,
//URL_RUP_PEROPD_PENYEDIA_DLM,
//URL_RUP_PEROPD_PENYEDIA_DLM,}

type RupRecord struct {
	Kode                   string
	Opd                    string
	PaketPenyedia          string
	PaguPenyedia           string
	SwakelolaPaket         string
	SwakelolaPagu          string
	PenyediaSwakelolaPaket string
	PenyediaSwakelolaPagu  string
	TotalPaket             string
	TotalPagu              string
}

type PacketResponse struct {
	Draw            string          `json:"draw"`
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	Data            [][]interface{} `json:"data"`
}

type rupclient struct {
	col *colly.Collector
}
type RupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}
type data [][]string

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("sirup.lkpp.go.id"),
	)

	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
	})
	rc := rupclient{col: c}

	fileName := "opd_sirup_2018.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create %s", fileName)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Kode RUP", "Nama Paket",
		"Pagu", "Metode", "Dana", "RUP", "Waktu", "Opd",
	})
	//opd
	opds := rc.buildOpdKode()
	for kode, opd_name := range opds {
		fmt.Println("kode:opd", kode, opd_name)
		url := addQuerytoURL(kode).String()
		res := rc.fetchRupOpd(url)
		for _, item := range res {
			item := append(item, opd_name)
			writer.Write(item)
		}
	}
}
func (r rupclient) fetchRupOpd(url string) data {
	resp := new(RupResponse)
	r.col.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &resp)
		if err != nil {
			log.Fatalln(err)
		}
	})
	r.col.Visit(url)
	r.col.Wait()
	data := resp.AaData
	return data
}

func (r rupclient) buildOpdKode() (res map[string]string) {
	result := make(map[string]string)
	s, err := r.fetchRekapRup()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range s.AaData {
		result[item[0]] = item[1]
	}
	return result
}

func addQuerytoURL(idSatker string) *url.URL {
	basePath := "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker"
	u, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("tahun", "2018")
	q.Add("idSatker", idSatker)
	u.RawQuery = q.Encode()
	return u
}

func (r rupclient) fetchRekapRup() (RupResponse, error) {
	baseURL := "https://sirup.lkpp.go.id/sirup/datatablectr/datatableruprekapkldi?idKldi=D128&tahun=2018"
	var resp RupResponse
	r.col.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &resp)
		if err != nil {
			log.Fatalln(err)
		}
	})
	r.col.Visit(baseURL)
	r.col.Wait()

	return resp, nil
}

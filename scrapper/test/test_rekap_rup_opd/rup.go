// main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/queue"
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

type RupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}

func main() {
	fileName := "opd_sirup.csv"
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

	c := colly.NewCollector(
		colly.AllowedDomains("sirup.lkpp.go.id"),
	)

	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
	})

	extensions.RandomUserAgent(c)

	opd_to_fetch := buildOpdKode() // map[kode]opd
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		//fmt.Println(string(r.Body))
		var s = new(RupResponse)
		err := json.Unmarshal(r.Body, &s)
		if err != nil {
			log.Fatalln(err)
		}
		//fmt.Println(s.ITotalDisplayRecords, s.SEcho, s.AaData)
		for _, item := range s.AaData {
			item := append(item)
			writer.Write(item)

		}

	})
	for kode, _ := range opd_to_fetch {
		url := addQuerytoURL(kode).String()

		c.Visit(url)

	}
	c.Wait()
}

func addQuerytoURL(idSatker string) *url.URL {
	basePath := "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker"
	u, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("tahun", "2019")
	q.Add("idSatker", idSatker)
	u.RawQuery = q.Encode()
	return u
}

func buildOpdKode() (res map[string]string) {
	result := make(map[string]string)
	s, err := fetchRekapRup()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range s.AaData {
		result[item[0]] = item[1]
	}
	return result
}

func fetchRekapRup() (RupResponse, error) {
	baseURL := "https://sirup.lkpp.go.id/sirup/datatablectr/datatableruprekapkldi?idKldi=D128&tahun=2019"

	var err error
	var client = &http.Client{}
	var data = RupResponse{}

	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func AddRawUrl(rows []string) *queue.Queue {
	Q, _ := queue.New(
		1, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)
	for _, url := range rows {
		Q.AddURL(url)
	}
	return Q
}

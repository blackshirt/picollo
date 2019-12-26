// main.go
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

var REKAP_RUP_SATKER_URL string = "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker?tahun=2019&idSatker=63418"
var REKAP_LPSE_SATKER_URL string = "http://lpse.kebumenkab.go.id/eproc4/dt/lelang?draw=2&authenticityToken=1ff307a1d53e64aa12109faac7337eb469523994"

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
	fileName := "sirup.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create %s", fileName)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Kode Opd", "Opd", "Penyedia (Paket)",
		"Penyedia (Pagu)", "Swa (Paket)", "Swa (Pagu)", "Penyedia dlm Swa(Paket)", "Penyedia dlm Swa (Pagu)",
		"Total paket", "Total Pagu"})

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	extensions.RandomUserAgent(c)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		fmt.Println("UserAgent", r.Headers.Get("User-Agent"))
	})
	/*
		c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
			e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
				var productName, stars, price string

				productName = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")

				if productName == "" {
					// If we can't get any name, we return and go directly to the next element
					return
				}

				writer.Write([]string{
					productName,
				})
			})
		})
	*/

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
			writer.Write(item)
		}

	})
	c.Visit("https://sirup.lkpp.go.id/sirup/datatablectr/datatableruprekapkldi?idKldi=D128&tahun=2019")
	c.Wait()
}

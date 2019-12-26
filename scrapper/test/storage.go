package scrapper

import (
	"encoding/csv"
	"net/url"
	"os"

	"github.com/gocolly/colly"
)

type Storage interface {
	Save(data []byte) (err error)
}

type Scrapper interface {
	Scrape(url url.URL) error
}

type collyScrapper struct {
	client  *colly.Collector
	storage Storage
}

func (cs collyScrapper) Scrape(addr url.URL) (err error) {
	// to be implemented
	cs.client.OnResponse(func(r *colly.Response) {
		cs.storage.Save(r.Body) // r.Body was []byte
	})
	cs.client.Visit(addr.String())
}

type csvStorage struct {
	fname string
}

func (cs csvStorage) Save(data []byte) (err error) {
	file, err := os.Create(cs.fname)
	defer file.Close()
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, item := range data {
		writer.Write(string(item[:])) // Write accept []string
	}
	return nil
}

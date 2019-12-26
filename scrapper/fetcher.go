package scrapper

import (
	"encoding/json"
	"errors"
	"log"

	"picollo/model"

	"github.com/gocolly/colly"
)

type rupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}

type Scrapper interface {
	Scrape(url string) (*Results, error)
}

type Results struct {
	Body []byte
}

type collyScrapper struct {
	client *colly.Collector
	svc    model.Service
}

func New(c *colly.Collector, s model.Service) Scrapper {
	return &collyScrapper{
		client: c,
		svc:    s,
	}
}

func (f *collyScrapper) Scrape(url string) (*Results, error) {
	if url == "" { //should valid
		return nil, errors.New("null url to scrape")
	}
	res, err := f.visit(url)
	if err != nil {
		log.Fatal(err)
	}
	return res, nil
}

func (f *collyScrapper) visit(url string) (*Results, error) {
	err := f.client.Visit(url)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("collector error in visit")
	}

	res := &Results{}
	f.client.OnResponse(func(r *colly.Response) {
		res.Body = r.Body
	})
	f.client.Wait()

	return res, nil
}

//decode unmarshall from Results byte to rupResponse in the form json
func (f *collyScrapper) decode(b *Results) (*rupResponse, error) {
	if b == nil {
		return nil, errors.New("Nil results params")
	}
	resp := &rupResponse{}
	err := json.Unmarshal(b.Body, resp)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resp, nil
}

// Fetch rup with spesific option and return response
func (f *collyScrapper) fetchRup(opt model.RupOptions) (*rupResponse, error) {
	b := NewLinkBuilder(WithRupOption(opt))
	link, err := b.buildPath()
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error in fetchRup")
	}

	resp, err := f.unmarshall(link.String())
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

// Fetch rekap rup with spesific option and return response
func (f *collyScrapper) fetchRekap(opt model.RupOptions) (*rupResponse, error) {
	b := NewLinkBuilder(
		WithRupOption(opt),
		WithRekap(true),
	)
	link, err := b.buildPath()
	if err != nil {
		log.Fatal(err)
	}
	resp, err := f.unmarshall(link.String())
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func (f *collyScrapper) unmarshall(link string) (*rupResponse, error) {
	response := &rupResponse{}
	f.client.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, response)
		if err != nil {
			log.Fatalln(err)
		}
	})

	f.client.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	if err := f.client.Visit(link); err != nil {
		log.Fatal(err)
		return nil, errors.New("collector error in visit")
	}
	f.client.Wait()
	return response, nil
}

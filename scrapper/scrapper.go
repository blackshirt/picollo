package scrapper

import (
	"encoding/json"
	"errors"

	"picollo/model"

	"github.com/gocolly/colly"
)

var (
	ErrCollyinVisit  = errors.New("Error colly collector error in visit")
	ErrOnNilResults  = errors.New("Error on Nil Results params")
	ErrOnNilResponse = errors.New("Error on nil rup response params")
)

type RupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}

type Scrapper interface {
	Scrape(u string) (*Results, error)
	UnMarshall(r *Results) (*RupResponse, error)
}

type Results struct {
	Body []byte
}

// marshall Result Body to RupResponse
func (r *Results) unmarshallToResponse(res *RupResponse) error {
	if res == nil {
		return ErrOnNilResponse
	}
	err := json.Unmarshal(r.Body, res)
	return err
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

// Scrape visit the url and set the response Results body to colly response
func (f *collyScrapper) Scrape(url string) (*Results, error) {
	err := f.client.Visit(url)
	if err != nil {
		return nil, ErrCollyinVisit
	}

	res := &Results{}
	f.client.OnResponse(func(r *colly.Response) {
		res.Body = r.Body
	})
	f.client.Wait()

	return res, nil
}

// UnMarshall unmarshall Results to RupResponse in the form json
func (f *collyScrapper) UnMarshall(b *Results) (*RupResponse, error) {
	if b == nil {
		return nil, ErrOnNilResults
	}
	resp := &RupResponse{}
	err := b.unmarshallToResponse(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

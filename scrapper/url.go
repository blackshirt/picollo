package scrapper

import (
	"errors"
	"log"
	"net/url"
	"picollo/model"
	"strconv"
	"time"
)

var (
	// base link
	rupBaseUrl string = "https://sirup.lkpp.go.id/sirup/datatablectr/"
	//peropd path
	pathOpdPenyedia             string = "dataruppenyediasatker"
	pathOpdSwakelola            string = "datarupswakelolasatker"
	pathOpdPenyediaDlmSwakelola string = "dataruppenyediaswakelolaallrekap"
	//full perkategori path
	pathFullPenyedia             string = "dataruppenyediakldi"
	pathFullSwakelola            string = "datarupswakelolakldi"
	pathFullPenyediaDlmSwakelola string = "dataruppenyediaswakelolaallrekapkldi"
	//path rekap
	pathRekap string = "datatableruprekapkldi"
)

var (
	lpseBaseUrl    string = "https://lpse.kebumenkab.go.id/eproc4/dt/"
	lpsePathLelang string = "lelang"
	lpsePathPeel   string = "pl"
)

type availableQs struct {
	//all has this
	tipe model.Type

	//rup
	year         string
	useRekapLink bool
	rupKategori  model.Kategori

	//rekap rup
	idKldi string

	//peropd rup
	idSatker string

	//lpse
	lpseKategori      model.KategoriLpse
	lpseMetode        model.MetodeLpse
	rkn_nama          string
	authenticityToken string
}

func (aq *availableQs) baseUrl() string {
	var baseUrl string
	switch aq.tipe {
	case model.TypeRup:
		baseUrl = rupBaseUrl
	case model.TypeOpd:
		baseUrl = rupBaseUrl
	case model.TypePacket:
		baseUrl = lpseBaseUrl
	default:
		baseUrl = ""
	}
	return baseUrl
}

func (aq *availableQs) basePath() (*url.URL, error) {

	switch aq.tipe {
	case model.TypeRup:
		path, err := rupPerKategoriPath(aq.rupKategori, aq.useRekapLink)
		if err != nil {
			return nil, err
		}
		return path, err
	case model.TypeOpd:
		path, err := rupRekapOpdPath()
		if err != nil {
			return nil, err
		}
		return path, err
	case model.TypePacket:
		path, err := lpsePath(aq.lpseMetode)
		if err != nil {
			return nil, err
		}
		return path, err
	default:
		return nil, errors.New("invalid tipe")
	}

}

type linkBuilder struct {
	useRekap bool
	opt      *model.RupOptions
}

type LinkOption func(*linkBuilder)

func NewLinkBuilder(options ...LinkOption) *linkBuilder {
	b := &linkBuilder{}
	b.Init()

	for _, set := range options {
		set(b)
	}

	// b.parseSettingsFromEnv()
	return b
}

func UseRekap(flag bool) LinkOption {
	return func(b *linkBuilder) {
		b.useRekap = flag
	}
}

func UseRupOption(o model.RupOptions) LinkOption {
	return func(b *linkBuilder) {
		b.opt = &o
	}
}

// Year set year used to fetch, default to current year
func UseTahun(y string) LinkOption {
	return func(b *linkBuilder) {
		if y == "" {
			y = strconv.Itoa(time.Now().Year())
		}
		b.opt.Tahun = y
	}
}

func UseKodeOpd(idSatker string) LinkOption {
	return func(b *linkBuilder) {
		if idSatker == "" {
			idSatker = "wr0n9c0d3"
		}
		b.opt.KodeOpd = idSatker
	}
}

func UseCategory(m model.Kategori) LinkOption {
	return func(b *linkBuilder) {
		b.opt.Kategori = m
	}
}

// Init initialize link builder struct
func (b *linkBuilder) Init() {
	// b.Tahun = strconv.Itoa(time.Now().Year())
	if b.opt == nil {
		opt := &model.RupOptions{}
		opt.Tahun = strconv.Itoa(time.Now().Year()) // set to current year
		b.opt = opt
	}
	if b.opt.KodeOpd == "" {
		b.opt.KodeOpd = "wr0n9c0d3"
	}
}

func (b *linkBuilder) buildPath() (*url.URL, error) {
	var link *url.URL
	switch b.useRekap {
	case true: // use rekap link
		u, err := rupFullPath(b.opt.Kategori)
		if err != nil {
			log.Fatal(err)
		}
		qs := map[string]string{
			"tahun":  b.opt.Tahun,
			"idKldi": "D128", // kebumen
		}

		link = addQsToUrl(u, qs)

	case false: //peropd link
		u, err := rupOpdPath(b.opt.Kategori)
		if err != nil {
			log.Fatal(err)
		}
		qs := map[string]string{
			"tahun":    b.opt.Tahun,
			"idSatker": b.opt.KodeOpd,
		}

		link = addQsToUrl(u, qs)
	}
	return link, nil
}

//add path string to baseUrl percategory to construct url
func addPath(baseUrl, path string) (*url.URL, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	u.Path += path
	return u, nil
}

// add query string map to url
func addQsToUrl(u *url.URL, qs map[string]string) *url.URL {
	q := u.Query()
	for key, val := range qs {
		q.Add(key, val)
	}
	u.RawQuery = q.Encode()
	return u
}

func addYeartoPath(u *url.URL, year string) (*url.URL, error) {
	if u == nil {
		return nil, errors.New("Nil base path")
	}
	q := u.Query()
	if year == "" {
		year = strconv.Itoa(time.Now().Year())
	}
	q.Set("tahun", year)
	u.RawQuery = q.Encode()
	return u, nil
}

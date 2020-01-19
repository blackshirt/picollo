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
	baseUrlRUP string = "https://sirup.lkpp.go.id/sirup/datatablectr/"
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
		u, err := fullPath(b.opt.Kategori)
		if err != nil {
			log.Fatal(err)
		}
		qs := map[string]string{
			"tahun":  b.opt.Tahun,
			"idKldi": "D128", // kebumen
		}

		link = addQsToUrl(u, qs)

	case false: //peropd link
		u, err := opdPath(b.opt.Kategori)
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

//full path
func fullPath(cat model.Kategori) (*url.URL, error) {
	var link *url.URL
	if cat.IsValid() {
		switch cat {
		case model.KategoriPenyedia:
			link = addPath(baseUrlRUP, pathFullPenyedia)
		case model.KategoriSwakelola:
			link = addPath(baseUrlRUP, pathFullSwakelola)
		case model.KategoriPenyediaDlmSwakelola:
			link = addPath(baseUrlRUP, pathFullPenyediaDlmSwakelola)
		default:
			return nil, errors.New("Not valid full link")
		}
		return link, nil
	}
	return nil, errors.New("invalid categori")
}

//peropd path
func opdPath(cat model.Kategori) (*url.URL, error) {
	var link *url.URL
	if cat.IsValid() {
		switch cat {
		case model.KategoriPenyedia:
			link = addPath(baseUrlRUP, pathOpdPenyedia)
		case model.KategoriSwakelola:
			link = addPath(baseUrlRUP, pathOpdSwakelola)
		case model.KategoriPenyediaDlmSwakelola:
			link = addPath(baseUrlRUP, pathOpdPenyediaDlmSwakelola)
		default:
			return nil, errors.New("Not valid opd link")
		}
		return link, nil
	}
	return nil, errors.New("invalid categori")
}

//add path string to baseUrl percategory to construct url
func addPath(baseUrl, path string) *url.URL {
	u, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	u.Path += path
	return u
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

// Response semua rup lewat penyedia
// 0: "22605363"
// 1: "SMPN 1 BULUSPESANTREN"
// 2: "Penyediaan Bantuan Operasional Sekolah (BOS) jenjang SD/MI/SDLB dan SMP/MTS serta pesantren Salafiyah dan Satuan Pendidikan Non Islam Setara SD dan SMP"
// 3: "160000000"
// 4: "Pengadaan Langsung"
// 5: "APBD"
// 6: "22605363"
// 7: "January 2020"

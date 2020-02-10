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

func getBaseUrl(t model.Type) string {
	var baseUrl string
	switch t {
	case model.TypeRup, model.TypeOpd:
		baseUrl = rupBaseUrl
	case model.TypePacket:
		baseUrl = lpseBaseUrl
	default:
		baseUrl = ""
	}
	return baseUrl
}

type rupQs struct {
	useRekapLink bool
	kategori     model.Kategori
	year         string
	idSatker     string
	idKldi       string
}

func (rq *rupQs) rupKategoriPath() (*url.URL, error) {
	var path *url.URL
	switch rq.useRekapLink {
	case true:
		u, err := rupFullPath(rq.kategori)
		if err != nil {
			return nil, err
		}
		path = u
	case false:
		u, err := rupOpdPath(rq.kategori)
		if err != nil {
			return nil, err
		}
		path = u
	}
	return path, nil
}

func (rq *rupQs) buildRupUrl() (*url.URL, error) {
	path, err := rq.rupKategoriPath()
	if err != nil {
		return nil, err
	}

	if rq.year == "" {
		year := strconv.Itoa(time.Now().Year())
		rq.year = year
	}
	qs := make(map[string]string, 0)
	qs["tahun"] = rq.year
	switch rq.useRekapLink {
	case true:
		qs["idKldi"] = "D128"
	case false:
		qs["idSatker"] = rq.idSatker
	}

	u := addQsToUrl(path, qs)

	return u, nil
}

type lpseQs struct {
	met               model.MetodeLpse
	rkn_nama          string
	kategori          model.KategoriLpse
	authenticityToken string
}

func (lq *lpseQs) getauthenticityToken() string {
	return ""
}

func lpsePath(m model.MetodeLpse) (*url.URL, error) {
	if m.IsValid() {
		var lpsePath *url.URL
		switch m {
		case model.Lelang:
			link, err := addPath(lpseBaseUrl, lpsePathLelang)
			if err != nil {
				return nil, err
			}
			lpsePath = link
		case model.PengadaanLangsung:
			link, err := addPath(lpseBaseUrl, lpsePathPeel)
			if err != nil {
				return nil, err
			}
			lpsePath = link
		}
		return lpsePath, nil
	}
	return nil, errors.New("not valid metode lpse")
}

func (lq *lpseQs) buildLpseUrl() (*url.URL, error) {
	path, err := lpsePath(lq.met)
	if err != nil {
		return nil, err
	}
	qs := make(map[string]string, 0)
	qs["rkn_nama"] = lq.rkn_nama
	qs["kategori"] = lq.kategori.String()
	qs["authenticityToken"] = lq.authenticityToken

	u := addQsToUrl(path, qs)
	return u, nil
}

type urlBuilder struct {
	tipe       model.Type
	basePath   string
	rupFilter  *rupQs
	lpseFilter *lpseQs
}

func (ub *urlBuilder) buildURL() (*url.URL, error) {
	if ub.tipe.IsValid() {
		switch ub.tipe {
		case model.TypeRup:
			u, err := ub.rupFilter.buildRupUrl()
			if err != nil {
				return nil, err
			}
			return u, nil
		case model.TypeOpd:

		}
	}
	return nil, errors.New("invalid tipe")
}

func (ub *urlBuilder) baseUrl() (string, error) {
	baseUrl := getBaseUrl(ub.tipe)
	if baseUrl == "" {
		return "", errors.New("empty baseUrl result")
	}
	return baseUrl, nil
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

//full path
func rupFullPath(cat model.Kategori) (*url.URL, error) {
	var link *url.URL
	if cat.IsValid() {
		switch cat {
		case model.KategoriPenyedia:
			path, err := addPath(rupBaseUrl, pathFullPenyedia)
			if err != nil {
				return nil, err
			}
			link = path
		case model.KategoriSwakelola:
			path, err := addPath(rupBaseUrl, pathFullSwakelola)
			if err != nil {
				return nil, err
			}
			link = path
		case model.KategoriPenyediaDlmSwakelola:
			path, err := addPath(rupBaseUrl, pathFullPenyediaDlmSwakelola)
			if err != nil {
				return nil, err
			}
			link = path
		default:
			return nil, errors.New("Not valid full link")
		}
		return link, nil
	}
	return nil, errors.New("invalid categori")
}

//peropd path
func rupOpdPath(cat model.Kategori) (*url.URL, error) {
	var link *url.URL
	if cat.IsValid() {
		switch cat {
		case model.KategoriPenyedia:
			path, err := addPath(rupBaseUrl, pathOpdPenyedia)
			if err != nil {
				return nil, err
			}
			link = path
		case model.KategoriSwakelola:
			path, err := addPath(rupBaseUrl, pathOpdSwakelola)
			if err != nil {
				return nil, err
			}
			link = path
		case model.KategoriPenyediaDlmSwakelola:
			path, err := addPath(rupBaseUrl, pathOpdPenyediaDlmSwakelola)
			if err != nil {
				return nil, err
			}
			link = path
		default:
			return nil, errors.New("Not valid opd link")
		}
		return link, nil
	}
	return nil, errors.New("invalid categori")
}

func rupRekapOpdPath() (*url.URL, error) {
	link, err := addPath(rupBaseUrl, pathRekap)
	if err != nil {
		return nil, err
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

// Response semua rup lewat penyedia
// 0: "22605363"
// 1: "SMPN 1 BULUSPESANTREN"
// 2: "Penyediaan Bantuan Operasional Sekolah (BOS) jenjang SD/MI/SDLB dan SMP/MTS serta pesantren Salafiyah dan Satuan Pendidikan Non Islam Setara SD dan SMP"
// 3: "160000000"
// 4: "Pengadaan Langsung"
// 5: "APBD"
// 6: "22605363"
// 7: "January 2020"

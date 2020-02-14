package scrapper

import (
	"errors"
	"net/url"
	"picollo/model"
)

var (
	ErrNilURL            = errors.New("Error in operate on nil URL")
	ErrInvalidTipe       = errors.New("Error in operate on invalid tipe")
	ErrInvalidCategori   = errors.New("Error in operate on invalid categori")
	ErrBuildURL          = errors.New("Error in building url")
	ErrInvalidLpseMetode = errors.New("Error invalid lpse methode")
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
	lpseMetode model.MetodeLpse

	lpseKategori      model.KategoriLpse
	rkn_nama          string
	authenticityToken string
}

// BuildURL build full url
func (aq *availableQs) BuildURL() (*url.URL, error) {
	if aq.tipe.IsValid() {
		switch aq.tipe {
		case model.TypeRup:
			u, err := buildRupURL(aq.rupKategori, aq.useRekapLink, aq.year, aq.idSatker)
			if err != nil {
				return nil, ErrBuildURL
			}
			return u, nil
		case model.TypeOpd:
			u, err := buildOpdURL(aq.year)
			if err != nil {
				return nil, ErrBuildURL
			}
			return u, nil
		case model.TypePacket:
			u, err := buildLpseURL(aq.lpseMetode, aq.authenticityToken, aq.lpseKategori, aq.rkn_nama)
			if err != nil {
				return nil, ErrBuildURL
			}
			return u, nil
		}
	}
	return nil, ErrInvalidTipe
}

// set query string k=v to url `u`
func setQs(path *url.URL, key, value string) (*url.URL, error) {
	if path == nil {
		return nil, ErrNilURL
	}
	q := path.Query()
	q.Set(key, value)
	path.RawQuery = q.Encode()
	return path, nil
}

// add year `y` query string to url `u`
func addYeartoPath(u *url.URL, year string) (*url.URL, error) {
	u, err := setQs(u, "tahun", year)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//add path string to rawUrl to construct url
func addPath(rawUrl, path string) (*url.URL, error) {
	u, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		return nil, err
	}
	u.Path += path
	return u, nil
}

// type linkBuilder struct {
// 	useRekap bool
// 	opt      *model.RupOptions
// }

// type LinkOption func(*linkBuilder)

// func NewLinkBuilder(options ...LinkOption) *linkBuilder {
// 	b := &linkBuilder{}
// 	b.Init()

// 	for _, set := range options {
// 		set(b)
// 	}

// 	// b.parseSettingsFromEnv()
// 	return b
// }

// func UseRekap(flag bool) LinkOption {
// 	return func(b *linkBuilder) {
// 		b.useRekap = flag
// 	}
// }

// func UseRupOption(o model.RupOptions) LinkOption {
// 	return func(b *linkBuilder) {
// 		b.opt = &o
// 	}
// }

// // Year set year used to fetch, default to current year
// func UseTahun(y string) LinkOption {
// 	return func(b *linkBuilder) {
// 		if y == "" {
// 			y = strconv.Itoa(time.Now().Year())
// 		}
// 		b.opt.Tahun = y
// 	}
// }

// func UseKodeOpd(idSatker string) LinkOption {
// 	return func(b *linkBuilder) {
// 		if idSatker == "" {
// 			idSatker = "wr0n9c0d3"
// 		}
// 		b.opt.KodeOpd = idSatker
// 	}
// }

// func UseCategory(m model.Kategori) LinkOption {
// 	return func(b *linkBuilder) {
// 		b.opt.Kategori = m
// 	}
// }

// // Init initialize link builder struct
// func (b *linkBuilder) Init() {
// 	// b.Tahun = strconv.Itoa(time.Now().Year())
// 	if b.opt == nil {
// 		opt := &model.RupOptions{}
// 		opt.Tahun = strconv.Itoa(time.Now().Year()) // set to current year
// 		b.opt = opt
// 	}
// 	if b.opt.KodeOpd == "" {
// 		b.opt.KodeOpd = "wr0n9c0d3"
// 	}
// }

// func (b *linkBuilder) buildPath() (*url.URL, error) {
// 	var link *url.URL
// 	switch b.useRekap {
// 	case true: // use rekap link
// 		u, err := rupFullPath(b.opt.Kategori)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		qs := map[string]string{
// 			"tahun":  b.opt.Tahun,
// 			"idKldi": "D128", // kebumen
// 		}

// 		link = addQsToUrl(u, qs)

// 	case false: //peropd link
// 		u, err := rupOpdPath(b.opt.Kategori)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		qs := map[string]string{
// 			"tahun":    b.opt.Tahun,
// 			"idSatker": b.opt.KodeOpd,
// 		}

// 		link = addQsToUrl(u, qs)
// 	}
// 	return link, nil
// }

// // add query string map to url
// func addQsToUrl(u *url.URL, qs map[string]string) *url.URL {
// 	q := u.Query()
// 	for key, val := range qs {
// 		q.Add(key, val)
// 	}
// 	u.RawQuery = q.Encode()
// 	return u
// }

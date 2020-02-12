package scrapper

import (
	"errors"
	"net/url"

	"picollo/model"
)

func lpsePath(m model.MetodeLpse) (*url.URL, error) {
	if m.IsValid() {
		switch m {
		case model.Lelang:
			link, err := addPath(lpseBaseUrl, lpsePathLelang)
			if err != nil {
				return nil, err
			}
			return link, nil
		case model.PengadaanLangsung:
			link, err := addPath(lpseBaseUrl, lpsePathPeel)
			if err != nil {
				return nil, err
			}
			return link, nil
		}
	}
	return nil, errors.New("not valid metode lpse")
}

func buildLpseURL(m model.MetodeLpse, token string, k model.KategoriLpse, rkn string) (*url.URL, error) {
	path, err := lpsePath(m)
	if err != nil {
		return nil, err
	}
	l, _ := setAuthToken(path, token)
	p, _ := setLpseKategori(l, k)
	u, err := setRknNama(p, rkn)
	return u, err
}

func setAuthToken(p *url.URL, token string) (*url.URL, error) {
	if p == nil {
		return nil, errors.New("nil path")
	}
	l, err := setQs(p, "authenticityToken", token)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func setRknNama(p *url.URL, rkn string) (*url.URL, error) {
	if p == nil {
		return nil, errors.New("nil url")
	}
	l, err := setQs(p, "rkn_nama", rkn)
	if err != nil {
		return nil, err
	}
	return l, nil
}

func setLpseKategori(p *url.URL, k model.KategoriLpse) (*url.URL, error) {
	if p == nil {
		return nil, errors.New("nil url")
	}
	l, err := setQs(p, "kategori", k.String())
	if err != nil {
		return nil, err
	}
	return l, nil
}

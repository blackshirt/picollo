package scrapper

import (
	"net/url"

	"picollo/model"
)

// lpsePath get lpse path from metode `m`
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
	return nil, ErrInvalidLpseMetode
}

// buildLpseURL build lpse url using params provided
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

// setAuthToken set authenticityToken query string to lpse path using `token`
func setAuthToken(p *url.URL, token string) (*url.URL, error) {
	if p == nil {
		return nil, ErrNilURL
	}
	u, err := setQs(p, "authenticityToken", token)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// setRknNama set rkn_nama query string to lpse path
func setRknNama(p *url.URL, rkn string) (*url.URL, error) {
	if p == nil {
		return nil, ErrNilURL
	}
	u, err := setQs(p, "rkn_nama", rkn)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// setLpseKategori set kategori query string to lpse path
func setLpseKategori(p *url.URL, k model.KategoriLpse) (*url.URL, error) {
	if p == nil {
		return nil, ErrNilURL
	}
	u, err := setQs(p, "kategori", k.String())
	if err != nil {
		return nil, err
	}
	return u, nil
}

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

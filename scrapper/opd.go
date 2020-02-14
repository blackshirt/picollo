package scrapper

import "net/url"

// get rekap rup path
func rekapOpdPath() (*url.URL, error) {
	path, err := addPath(rupBaseUrl, pathRekap)
	if err != nil {
		return nil, err
	}
	return path, nil
}

// build rekap rup url
func buildOpdURL(year string) (*url.URL, error) {
	p, err := rekapOpdPath()
	if err != nil {
		return nil, err
	}
	l, err := addYeartoPath(p, year)
	if err != nil {
		return nil, err
	}
	u, err := setQs(l, "idKldi", "D128")
	if err != nil {
		return nil, err
	}
	return u, nil
}

package scrapper

import (
	"net/url"

	"picollo/model"
)

// get rup path based on kategori `c` and flag `useRekap`
func rupPerKategoriPath(c model.Kategori, useRekap bool) (*url.URL, error) {
	if useRekap {
		path, err := perfullPath(c)
		if err != nil {
			return nil, err
		}
		return path, nil
	}

	path, err := peropdPath(c)
	if err != nil {
		return nil, err
	}

	return path, nil
}

// build rup Url
func buildRupURL(c model.Kategori, useRekap bool, tahun string, idSatker string) (*url.URL, error) {
	if useRekap {
		u, err := addQsToRupFullPath(c, tahun)
		if err != nil {
			return nil, err
		}
		return u, nil
	}
	u, err := addQsToOpdPath(c, tahun, idSatker)
	if err != nil {
		return nil, err
	}
	return u, nil
}

//peropd path
func peropdPath(c model.Kategori) (*url.URL, error) {
	if c.IsValid() {
		switch c {
		case model.KategoriPenyedia:
			path, err := addPath(rupBaseUrl, pathOpdPenyedia)
			if err != nil {
				return nil, err
			}
			return path, nil
		case model.KategoriSwakelola:
			path, err := addPath(rupBaseUrl, pathOpdSwakelola)
			if err != nil {
				return nil, err
			}
			return path, nil
		case model.KategoriPenyediaDlmSwakelola:
			path, err := addPath(rupBaseUrl, pathOpdPenyediaDlmSwakelola)
			if err != nil {
				return nil, err
			}
			return path, nil
		}
	}
	return nil, ErrInvalidCategori
}

//full path
func perfullPath(c model.Kategori) (*url.URL, error) {
	if c.IsValid() {
		switch c {
		case model.KategoriPenyedia:
			path, err := addPath(rupBaseUrl, pathFullPenyedia)
			if err != nil {
				return nil, err
			}
			return path, nil
		case model.KategoriSwakelola:
			path, err := addPath(rupBaseUrl, pathFullSwakelola)
			if err != nil {
				return nil, err
			}
			return path, nil
		case model.KategoriPenyediaDlmSwakelola:
			path, err := addPath(rupBaseUrl, pathFullPenyediaDlmSwakelola)
			if err != nil {
				return nil, err
			}
			return path, nil
		}

	}
	return nil, ErrInvalidCategori
}

// set query string for full rup path
func addQsToRupFullPath(c model.Kategori, year string) (*url.URL, error) {
	path, err := perfullPath(c)
	if err != nil {
		return nil, err
	}
	u, _ := addYeartoPath(path, year)
	link, err := setQs(u, "idKldi", "D128")
	return link, err
}

// set query string for peropd rup path for satker `idSatker`
func addQsToOpdPath(c model.Kategori, year string, idSatker string) (*url.URL, error) {
	p, err := peropdPath(c)
	if err != nil {
		return nil, err
	}
	u, _ := addYeartoPath(p, year)
	link, err := setQs(u, "idSatker", idSatker)
	return link, err
}

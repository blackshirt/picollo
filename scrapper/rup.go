package scrapper

import (
	"net/url"

	"picollo/model"
)

func rupPerKategoriPath(cat model.Kategori, useRekap bool) (*url.URL, error) {
	if useRekap {
		path, err := rupFullPath(cat)
		if err != nil {
			return nil, err
		}
		return path, nil
	}

	path, err := rupOpdPath(cat)
	if err != nil {
		return nil, err
	}

	return path, nil
}

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
func rupOpdPath(cat model.Kategori) (*url.URL, error) {
	if cat.IsValid() {
		switch cat {
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
func rupFullPath(cat model.Kategori) (*url.URL, error) {
	if cat.IsValid() {
		switch cat {
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

func addQsToRupFullPath(c model.Kategori, year string) (*url.URL, error) {
	path, err := rupFullPath(c)
	if err != nil {
		return nil, err
	}
	u, _ := addYeartoPath(path, year)
	link, err := setQs(u, "idKldi", "D128")
	return link, err
}

func addQsToOpdPath(c model.Kategori, year string, idSatker string) (*url.URL, error) {
	p, err := rupOpdPath(c)
	if err != nil {
		return nil, err
	}
	u, _ := addYeartoPath(p, year)
	link, err := setQs(u, "idSatker", idSatker)
	return link, err
}

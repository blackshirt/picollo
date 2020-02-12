package scrapper

import (
	"errors"
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
	return nil, errors.New("invalid categori")
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
	return nil, errors.New("invalid categori")
}

func setQs(path *url.URL, key, value string) (*url.URL, error) {
	if path == nil {
		return nil, errors.New("Nil path")
	}
	q := path.Query()
	q.Set(key, value)
	path.RawQuery = q.Encode()
	return path, nil
}

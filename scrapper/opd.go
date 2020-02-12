package scrapper

import "net/url"

func rupRekapOpdPath() (*url.URL, error) {
	path, err := addPath(rupBaseUrl, pathRekap)
	if err != nil {
		return nil, err
	}
	return path, nil
}

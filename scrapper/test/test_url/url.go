package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	fmt.Println(addQuerytoURL("63401"))
	fmt.Println(addQuerytoURL("63411"))
}
func addQuerytoURL(idSatker string) *url.URL {
	basePath := "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker"
	u, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Add("tahun", "2019")
	q.Add("idSatker", idSatker)
	u.RawQuery = q.Encode()
	return u
}

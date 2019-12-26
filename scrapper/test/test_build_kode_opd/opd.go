package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type RupResponse struct {
	AaData               [][]string `json:"aaData"`
	ITotalDisplayRecords int        `json:"iTotalDisplayRecords"`
	SEcho                int        `json:"sEcho"`
}

func main() {
	hasil, err := fetchRekapRup()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hasil.AaData, hasil.ITotalDisplayRecords)
	kode := buildOpdKode()
	for i, item := range kode {
		fmt.Println(i, item)
	}
}

func buildOpdKode() (res []string) {
	var result []string
	s, err := fetchRekapRup()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range s.AaData {
		result = append(result, item[0])
	}
	return result
}

func fetchRekapRup() (RupResponse, error) {
	baseURL := "https://sirup.lkpp.go.id/sirup/datatablectr/datatableruprekapkldi?idKldi=D128&tahun=2019"

	var err error
	var client = &http.Client{}
	var data = RupResponse{}

	request, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

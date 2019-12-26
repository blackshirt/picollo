package scrapper

import (
	"fmt"
	"picollo/model"
	"testing"
)

func TestInit(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"tahun", "", "2019"},
		{"tahun 2", "2019", "2019"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			//  test empty option builder
			b := NewLinkBuilder()
			got := b.opt.Tahun
			if got != tc.want {
				t.Errorf("Init was incorrect, got: %s, want: %s.", got, tc.want)
			}
		})
	}
}

func TestTahun(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{"empty tahun", "", "2019"},
		{"tahun 1", "2019", "2019"},
		{"tahun 2", "2020", "2020"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.name), func(t *testing.T) {
			b := NewLinkBuilder(WithTahun(tc.input))
			got := b.opt.Tahun
			if got != tc.want {
				t.Errorf("Tahun builder was incorrect, got: %s, want: %s.", got, tc.want)
			}
		})
	}
}

func TestWithCategory(t *testing.T) {
	testCases := []struct {
		input model.Kategori
		want  model.Kategori
	}{
		{model.KategoriPenyediaDlmSwakelola, model.KategoriPenyediaDlmSwakelola},
		{model.KategoriPenyedia, model.KategoriPenyedia},
		{model.KategoriSwakelola, model.KategoriSwakelola},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s:", tc.input), func(t *testing.T) {
			b := NewLinkBuilder(WithCategory(tc.input))

			if got := b.opt.Kategori; got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestLinkBuilderUsingRekapLink(t *testing.T) {
	testCases := []struct {
		useRekap bool
		name     string
		category model.Kategori
		tahun    string
		want     string
	}{
		{true, "rekap penyedia 2019", model.KategoriPenyedia, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediakldi?idKldi=D128&tahun=2019"},
		{true, "rekap penyedia 2020", model.KategoriPenyedia, "2020", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediakldi?idKldi=D128&tahun=2020"},
		{true, "rekap swakelola 2019", model.KategoriSwakelola, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/datarupswakelolakldi?idKldi=D128&tahun=2019"},
		{true, "rekap swakelola 2020", model.KategoriSwakelola, "2020", "https://sirup.lkpp.go.id/sirup/datatablectr/datarupswakelolakldi?idKldi=D128&tahun=2020"},
		{true, "rekap penyedia-swakelola 2019", model.KategoriPenyediaDlmSwakelola, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediaswakelolaallrekapkldi?idKldi=D128&tahun=2019"},
		{true, "rekap penyedia-swakelola 2020", model.KategoriPenyediaDlmSwakelola, "2020", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediaswakelolaallrekapkldi?idKldi=D128&tahun=2020"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s:", tc.name), func(t *testing.T) {
			b := NewLinkBuilder(
				WithCategory(tc.category),
				WithRekap(tc.useRekap),
				WithTahun(tc.tahun))
			link, _ := b.buildPath()
			got := link.String()
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func TestLinkBuilderUsingOpdLink(t *testing.T) {
	testCases := []struct {
		name     string
		opdKode  string
		category model.Kategori
		tahun    string
		want     string
	}{
		{"opd penyedia 2019 empty", "", model.KategoriPenyedia, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker?idSatker=wr0n9c0d3&tahun=2019"},
		{"opd penyedia 2019", "159844", model.KategoriPenyedia, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker?idSatker=159844&tahun=2019"},
		{"opd penyedia 2020", "159844", model.KategoriPenyedia, "2020", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediasatker?idSatker=159844&tahun=2020"},
		{"opd swakelola 2019", "159844", model.KategoriSwakelola, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/datarupswakelolasatker?idSatker=159844&tahun=2019"},
		{"opdswakelola2020empty", "", model.KategoriSwakelola, "2020", "https://sirup.lkpp.go.id/sirup/datatablectr/datarupswakelolasatker?idSatker=wr0n9c0d3&tahun=2020"},
		{"opd penyedia-swakelola 2019", "159844", model.KategoriPenyediaDlmSwakelola, "2019", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediaswakelolaallrekap?idSatker=159844&tahun=2019"},
		{"opd penyedia-swakelola 2020", "159844", model.KategoriPenyediaDlmSwakelola, "2020", "https://sirup.lkpp.go.id/sirup/datatablectr/dataruppenyediaswakelolaallrekap?idSatker=159844&tahun=2020"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s:", tc.name), func(t *testing.T) {
			b := NewLinkBuilder(
				WithCategory(tc.category),
				WithTahun(tc.tahun),
				WithKodeOpd(tc.opdKode))
			link, _ := b.buildPath()
			got := link.String()
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

package model

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Type test
func TestTypeFromFunction(t *testing.T) {
	assert := assert.New(t)
	testCase := []struct {
		name  string // test name
		input string // input param
		want  int    // result wanted
		err   error  // error want to check
	}{
		{"UnknownType", "UnknownType", 0, nil},
		{"Invalid Type", "Invalid Type", -1, fmt.Errorf("%s is not a valid Type", "Invalid Type")},
		{"Rup", "Rup", 1, nil},
		{"Opd", "Opd", 2, nil},
		{"Packet", "Packet", 3, nil},
		{"Invalid Packet", "Invalid Packet", -1, fmt.Errorf("%s is not a valid Type", "Invalid Packet")},
	}
	for _, tc := range testCase {
		t.Run(fmt.Sprintf("%s:", tc.name), func(t *testing.T) {
			got, err := TypeFrom(tc.input) // function to be checked
			assert.Equal(tc.err, err)      // assert the error
			if int(got) != tc.want {       // make sure we get int value
				t.Errorf("got %s; want %d", got, tc.want)
			}
		})
	}
}

func TestTypeIsValidFunction(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name        string // test name
		input       string // input
		want        bool
		shouldError bool
	}{
		{"type Unknown", "UnknownType", true, false},
		{"type Rup", "Rup", true, false},
		{"type Invalid", "Invalid Type", false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := TypeFrom(tc.input)
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			got := res.IsValid()
			assert.Equal(tc.want, got)
		})
	}
}

func TestTypeUnmarshalGQLFunction(t *testing.T) {
	assert := assert.New(t)
	testCase := []struct {
		name        string
		inputTipe   Type        // input type
		inputValue  interface{} // input value
		expectedErr error
	}{
		{"unknown type", UnknownType, "UnknownType", nil},
		{"unknown string", UnknownType, "UnknownString", fmt.Errorf("%s is not a valid Type", "UnknownString")},
		{"Rup type", Rup, "Rup", nil},
		{"Unknown Rup type", Rup, "Rup Unknown", fmt.Errorf("%s is not a valid Type", "Rup Unknown")},
		{"Invalid type", Rup, 100, fmt.Errorf("enums must be strings")},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			gotErr := tc.inputTipe.UnmarshalGQL(tc.inputValue)
			assert.Equal(tc.expectedErr, gotErr)
		})
	}
}

func TestTypeMarshalGQLFunction(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name      string
		inputTipe Type
		expected  string
	}{
		{"Rup type", Rup, "Rup"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)                  // bytes.Buffer implement io.Writer
			tc.inputTipe.MarshalGQL(buf)              // MarshallGQL places quote to string
			got, err := strconv.Unquote(buf.String()) // we unquote it
			if err != nil {
				t.Fatalf("Error %s", err)
			}
			assert.Equal(tc.expected, got)
		})
	}
}

// Jenis test
// TestJenisFromFunction test JenisFrom function for Jenis type
func TestJenisFromFunction(t *testing.T) {
	// assert := assert.New(t)
	cases := []struct {
		name        string // test name
		input       string // input param
		want        int    // result wanted
		shouldError bool   // should error want to check
	}{
		{"Unknown Jenis", "UnknownJenis", 0, false},
		{"Invalid Jenis", "Invalid Jenis", -1, true}, // this should error
		{"Barang", "Barang", 1, false},
		{"Konstruksi", "Konstruksi", 2, false},
		{"Konsultansi", "Konsultansi", 3, false},
		{"Invalid Barang", "Invalid Barang", -1, true}, // this should error
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := JenisFrom(tc.input) // function to be checked
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			if int(got) != tc.want { // make sure we get int value
				t.Errorf("got %s; want %d", got, tc.want)
			}
		})
	}
}

// TestJenisIsValidFunction test IsValid function for Jenis type
func TestJenisIsValidFunction(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name        string // test name
		input       string // input
		want        bool
		shouldError bool
	}{
		{"jenis Unknown", "UnknownJenis", true, false},
		{"jenis Barang", "Barang", true, false},
		{"jenis Konstruksi", "Konstruksi", true, false},
		{"jenis Invalid", "Invalid Jenis", false, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := JenisFrom(tc.input)
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			got := res.IsValid()
			assert.Equal(tc.want, got)
		})
	}
}

func TestJenisMarshalGQLFunction(t *testing.T) {
	assert := assert.New(t)
	cases := []struct {
		name      string
		inputTipe Jenis
		expected  string
	}{
		{"Jenis type", UnknownJenis, "\"UnknownJenis\""}, // MarshallGQl add quote
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buf := new(bytes.Buffer)     // bytes.Buffer implement io.Writer
			tc.inputTipe.MarshalGQL(buf) // MarshallGQL add quote to string
			got := buf.String()
			// if err != nil {
			// 	t.Fatalf("Error %s", err)
			// }
			assert.Equal(tc.expected, got)
		})
	}
}

// TestJenisUnmarshalGQLFunction test UnmarshalGQL function for Jenis type
func TestJenisUnmarshalGQLFunction(t *testing.T) {
	// assert := assert.New(t)
	cases := []struct {
		name        string
		inputTipe   Jenis       // input type
		inputValue  interface{} // input value
		shouldError bool
	}{
		{"unknown jenis", UnknownJenis, "UnknownJenis", false},
		{"unknown string", UnknownJenis, "UnknownString", true},
		{"Barang jenis", Barang, "Barang", false},
		{"Unknown Barang", Barang, "Barang Unknown", true},
		{"Invalid jenis", Barang, 100, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.inputTipe.UnmarshalGQL(tc.inputValue)
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			// assert.Equal(tc.shouldError, err)
		})
	}
}

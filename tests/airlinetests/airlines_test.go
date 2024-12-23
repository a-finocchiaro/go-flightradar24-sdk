package airlinetests

import (
	"errors"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/tests"
)

var validPng = []byte{
	0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
	0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
	0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x01, 0x03, 0x00, 0x00, 0x00, 0x66, 0xBC, 0x3A,
	0x25, 0x00, 0x00, 0x00, 0x03, 0x50, 0x4C, 0x54,
	0x45, 0xB5, 0xD0, 0xD0, 0x63, 0x04, 0x16, 0xEA,
	0x00, 0x00, 0x00, 0x1F, 0x49, 0x44, 0x41, 0x54,
	0x68, 0x81, 0xED, 0xC1, 0x01, 0x0D, 0x00, 0x00,
	0x00, 0xC2, 0xA0, 0xF7, 0x4F, 0x6D, 0x0E, 0x37,
	0xA0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xBE, 0x0D, 0x21, 0x00, 0x00, 0x01, 0x9A,
	0x60, 0xE1, 0xD5, 0x00, 0x00, 0x00, 0x00, 0x49,
	0x45, 0x4E, 0x44, 0xAE, 0x42, 0x60, 0x82,
}

func TestGetAirlines(t *testing.T) {
	goodSubtests := []tests.TestData{
		{
			Name: "No Error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
					"version": 1727163107,
					"rows": [
						{
							"Name": "American Airlines",
							"Code": "AA",
							"ICAO": "AAL"
						},
						{
							"Name": "Delta Air Lines",
							"Code": "DL",
							"ICAO": "DAL"
						},
						{
							"Name": "United Airlines",
							"Code": "UA",
							"ICAO": "UAL"
						}
					]
				}`), nil
			},
			ExpectedError: nil,
		},
	}

	for _, subtest := range goodSubtests {
		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetAirlines(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Version != 1727163107 {
				t.Errorf("Version: got %d, want %d", res.Version, 1727163107)
			}

			if len(res.Rows) != 3 {
				t.Errorf("Rows: got length %d, want %d", len(res.Rows), 3)
			}
		})
	}
}

func TestGetAirlineLogo(t *testing.T) {
	goodSubtests := []tests.TestData{
		{
			Name: "No Error",
			Requester: func(s string) ([]byte, error) {
				return validPng, nil
			},
			ExpectedError: nil,
		},
	}

	errorSubtests := []tests.TestData{
		{
			Name: "CRC error",
			Requester: func(s string) ([]byte, error) {
				invalidPng := validPng

				// modify the last checksum element to make it invalid
				invalidPng[len(validPng)-1] = 0x83

				return invalidPng, nil
			},
			ExpectedError: common.Fr24Error{Err: "png: invalid format: invalid checksum"},
		},
		{
			Name: "Invalid Data",
			Requester: func(s string) ([]byte, error) {
				invalidImage := []byte{
					0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
					0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52,
					0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01,
					0x08, 0x06, 0x00, 0x00, 0x00, 0x1F, 0x15, 0xC4,
					0x89, 0x00, 0x00, 0x00, 0x0A, 0x49, 0x44, 0x41,
					0x54, 0x78, 0x9C, 0x63, 0x64, 0x00, 0x00, 0x00,
					0x02, 0x00, 0x01, 0x45, 0x67, 0x89, 0xAB, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4E, 0x44,
					0xAE, 0x42, 0x60, 0x82,
				}

				return invalidImage, nil
			},
			ExpectedError: common.Fr24Error{Err: "zlib: invalid checksum"},
		},
	}

	for _, subtest := range goodSubtests {
		t.Run(subtest.Name+"_CDN", func(t *testing.T) {
			res, err := client.GetAirlineLogoCdn(subtest.Requester, "UA", "UAL")

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Len() == 0 {
				t.Errorf("Expected encoded PNG image as a response, got nil.")
			}
		})

		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetAirlineLogo(subtest.Requester, "UAL")

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Len() == 0 {
				t.Errorf("Expected encoded PNG image as a response, got nil.")
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.Name+"_CDN", func(t *testing.T) {
			res, err := client.GetAirlineLogoCdn(subtest.Requester, "UA", "UAL")

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if res.Len() != 0 {
				t.Errorf("Expected PNG image bytes to be empty.")
			}
		})

		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetAirlineLogo(subtest.Requester, "UAL")

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if res.Len() != 0 {
				t.Errorf("Expected PNG image bytes to be empty.")
			}
		})
	}
}

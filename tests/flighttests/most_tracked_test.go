package flighttests

import (
	"errors"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/tests"
)

func TestGetFr24MostTracked(t *testing.T) {
	goodSubtests := []tests.TestData{
		{
			Name: "No Error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
					"version": "0.3.9",
					"update_time": 1722142873.821783,
					"data": [
						{
							"flight_id": "3663d1ec",
							"flight": "RJ267",
							"callsign": "RJA267",
							"squawk": null,
							"clicks": 236,
							"from_iata": "AMM",
							"from_city": "Amman",
							"to_iata": "DTW",
							"to_city": "Detroit",
							"model": "B788",
							"type": "Boeing 787-8 Dreamliner"
						}
					]
				}`), nil
			},
			ExpectedError: nil,
		},
	}

	errorSubtests := []tests.TestData{
		{
			Name: "Request Error",
			Requester: func(s string) ([]byte, error) {
				return []byte{}, common.Fr24Error{Err: "some error"}
			},
			ExpectedError: common.Fr24Error{Err: "some error"},
		},
	}

	for _, subtest := range goodSubtests {
		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetFR24MostTracked(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			flight := res.Data[0]

			if flight.Flight_id != "3663d1ec" {
				t.Errorf("got %s want %q", flight.Flight_id, "3663d1ec")
			}

			if flight.Flight != "RJ267" {
				t.Errorf("got %s want %q", flight.Flight, "RJ267")
			}

			if flight.Callsign != "RJA267" {
				t.Errorf("got %s want %q", flight.Callsign, "RJA267")
			}

			if flight.Squawk != "" {
				t.Errorf("got %s want %q", flight.Squawk, "")
			}

			if flight.Clicks != 236 {
				t.Errorf("got %d want %d", flight.Clicks, 236)
			}

			if flight.From_iata != "AMM" {
				t.Errorf("got %s want %q", flight.From_iata, "AMM")
			}

			if flight.From_city != "Amman" {
				t.Errorf("got %s want %q", flight.From_city, "Amman")
			}

			if flight.To_iata != "DTW" {
				t.Errorf("got %s want %q", flight.To_iata, "DTW")
			}

			if flight.To_city != "Detroit" {
				t.Errorf("got %s want %q", flight.To_city, "Detroit")
			}

			if flight.Model != "B788" {
				t.Errorf("got %s want %q", flight.Model, "B788")
			}

			if flight.Aircraft_type != "Boeing 787-8 Dreamliner" {
				t.Errorf("got %s want %q", flight.Aircraft_type, "Boeing 787-8 Dreamliner")
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetFR24MostTracked(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected (%s), got error (%v)", subtest.ExpectedError, err)
			}

			if res.Data != nil {
				t.Errorf("Non empty struct returned for an error")
			}
		})
	}
}

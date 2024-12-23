package modeltests

import (
	"encoding/json"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/flights"
)

type FlightFeedTestData struct {
	Name     string
	JsonData []byte
}

func TestFlightUnmarshal(t *testing.T) {
	goodSubTests := []FlightFeedTestData{
		{
			Name: "No error",
			JsonData: []byte(`{
					"full_count": 17238,
					"version": 4,
					"3655152d": [
						"AC0A9A",
						39.8283,
						-101.8552,
						11,
						56100,
						6,
						"",
						"F-KPUB1",
						"BALL",
						"N875TH",
						1722800892,
						"",
						"",
						"",
						0,
						8576,
						"HBAL645",
						0,
						""
					]
				}`),
		},
	}

	errorSubTests := []FlightFeedTestData{
		{
			Name: "JSON unmarshal error",
			JsonData: []byte(`{
					"full_count": 17238,
					"version": 4,
					"3655152d": [
						"AC0A9A",
						39.8283,
						-101.8552,
						11,
						56100,
						6,
						"",
						"F-KPUB1",
						"BALL",
						"N875TH",
						1722800892,
						"",
						"",
						"",
						0,
						8576,
						"HBAL645",
						0,
						""
				}`),
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var flightFeed flights.Fr24FeedData

			err := json.Unmarshal(subtest.JsonData, &flightFeed)

			if err != nil {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if flightFeed.Full_count != 17238 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 17238, flightFeed.Full_count)
			}

			if flightFeed.Version != 4 {
				t.Errorf("Expected Version to be (%d), got (%d)", 4, flightFeed.Version)
			}

			if len(flightFeed.Flights) != 1 {
				t.Errorf("Expected 1 flight to be unpacked, received (%d)", len(flightFeed.Flights))
			}

			flight := flightFeed.Flights["3655152d"]

			if flight.Icao_24bit != "AC0A9A" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Icao_24bit)
			}
			if flight.Lat != 39.8283 {
				t.Errorf("Expected CHANGEME to equal (%f), got (%f)", 39.8283, flight.Lat)
			}
			if flight.Long != -101.8552 {
				t.Errorf("Expected CHANGEME to equal (%f), got (%f)", -101.8552, flight.Long)
			}
			if flight.Heading != 11 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 11, flight.Heading)
			}
			if flight.Altitude != 56100 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 56100, flight.Altitude)
			}
			if flight.Ground_speed != 6 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 6, flight.Ground_speed)
			}
			if flight.Squawk != "" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Squawk)
			}
			if flight.Fnumber != "F-KPUB1" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Fnumber)
			}
			if flight.Aircraft_code != "BALL" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Aircraft_code)
			}
			if flight.Registration != "N875TH" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Registration)
			}
			if flight.Time != 1722800892 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 1722800892, flight.Time)
			}
			if flight.Origin_airport_iata != "" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Origin_airport_iata)
			}
			if flight.Destination_airport_iata != "" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Destination_airport_iata)
			}
			if flight.Airline_iata != "" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Airline_iata)
			}
			if flight.On_ground != 0 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 0, flight.On_ground)
			}
			if flight.Vertical_speed != 8576 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 8576, flight.Vertical_speed)
			}
			if flight.Callsign != "HBAL645" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Callsign)
			}
			if flight.SomeNum != 0 {
				t.Errorf("Expected CHANGEME to equal (%d), got (%d)", 0, flight.SomeNum)
			}
			if flight.Airline_icao != "" {
				t.Errorf("Expected CHANGEME to equal (%s), got (%s)", "", flight.Airline_icao)
			}
		})
	}

	for _, subtest := range errorSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var flightFeed flights.Fr24FeedData
			err := json.Unmarshal(subtest.JsonData, &flightFeed)

			if err == nil {
				t.Error("Expected an error to be returned, got nil")
			}

			if len(flightFeed.Flights) != 0 {
				t.Errorf("Expected (%d) call, got (%d)", 0, len(flightFeed.Flights))
			}
		})
	}
}

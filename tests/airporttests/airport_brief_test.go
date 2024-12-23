package airporttests

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/airports"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/tests"
)

func loadJsonAirportBriefData(t *testing.T) []byte {
	data, err := os.ReadFile("./testdata/airport_brief.json")

	if err != nil {
		t.Error(err.Error())
	}

	return data
}

func TestGetAirportBrief(t *testing.T) {
	jsonRes := loadJsonAirportBriefData(t)

	goodSubtests := []airportSubtests{
		{
			TestData: tests.TestData{
				Name: "No Error",
				Requester: func(s string) ([]byte, error) {
					return jsonRes, nil
				},
				ExpectedError: nil,
			},
			plugins: []string{"details"},
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
			res, err := client.GetAirportBrief(subtest.Requester, "LHR")

			if !errors.Is(err, nil) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Name != "London Heathrow Airport" {
				t.Errorf("Expected Name (%v), got Name (%v)", "London Heathrow Airport", res.Name)
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.Name+"_CDN", func(t *testing.T) {
			res, err := client.GetAirportBrief(subtest.Requester, "LHR")

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if res.Name != "" {
				t.Errorf("Expected Name to be empty, got Name (%v)", res.Name)
			}
		})
	}
}

func TestAirportBriefDataStructs(t *testing.T) {
	var airport airports.AirportBriefResponse
	jsonData := loadJsonAirportBriefData(t)

	// unmarshal the json to an airport object
	if err := json.Unmarshal(jsonData, &airport); err != nil {
		t.Errorf("Error unmarshalling airport test data (%v)", err)
	}

	subtests := []tests.JsonValidationTest{
		{
			Name:     "AirportBriefResponse",
			Actual:   airport,
			Expected: airports.AirportBriefResponse{Details: airport.Details},
		},
		{
			Name:   "AirportBriefDetails",
			Actual: airport.Details,
			Expected: airports.AirportBriefDetails{
				Name: "London Heathrow Airport",
				Code: common.IataIcaoCode{
					Iata: "LHR",
					Icao: "EGLL",
				},
				Position: airport.Details.Position,
				Timezone: airport.Details.Timezone,
				Visible:  true,
				Website:  "http://www.heathrowairport.com/",
				Stats:    airport.Details.Stats,
			},
		},
		{
			Name:   "AirportPosition",
			Actual: airport.Details.Position,
			Expected: airports.AirportPosition{
				Latitude:  51.471626,
				Longitude: -0.467081,
				Altitude:  83,
				Country:   airport.Details.Position.Country,
				Region:    airport.Details.Position.Region,
			},
		},
		{
			Name:   "AirportBriefStats",
			Actual: airport.Details.Stats,
			Expected: airports.AirportBriefStats{
				Arrivals:   airport.Details.Stats.Arrivals,
				Departures: airport.Details.Stats.Departures,
			},
		},
		{
			Name:   "ArrivalDepartureAggregateStats",
			Actual: airport.Details.Stats.Arrivals,
			Expected: airports.ArrivalDepartureAggregateStats{
				DelayIndex: 0,
				DelayAvg:   nil,
				Total:      "7055",
				Hourly:     airport.Details.Stats.Arrivals.Hourly,
				Stats:      airport.Details.Stats.Arrivals.Stats,
			},
		},
		{
			Name:   "Hourly",
			Actual: airport.Details.Stats.Arrivals.Hourly,
			Expected: airports.Hourly{
				Hour0:  "0",
				Hour1:  "0",
				Hour2:  "0",
				Hour3:  "0",
				Hour4:  "6",
				Hour5:  "108",
				Hour6:  "316",
				Hour7:  "448",
				Hour8:  "445",
				Hour9:  "439",
				Hour10: "438",
				Hour11: "403",
				Hour12: "414",
				Hour13: "439",
				Hour14: "431",
				Hour15: "393",
				Hour16: "391",
				Hour17: "399",
				Hour18: "415",
				Hour19: "420",
				Hour20: "400",
				Hour21: "383",
				Hour22: "291",
				Hour23: "76",
			},
		},
	}

	for _, subtest := range subtests {
		t.Run(subtest.Name, func(t *testing.T) {
			if !reflect.DeepEqual(subtest.Expected, subtest.Actual) {
				t.Errorf("(%s) does not equal (%s)", subtest.Expected, subtest.Actual)
			}
		})
	}
}

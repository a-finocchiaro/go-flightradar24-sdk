package fr24

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"
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
			TestData: TestData{
				name: "No Error",
				requester: func(s string) ([]byte, error) {
					return jsonRes, nil
				},
				expectedError: nil,
			},
			plugins: []string{"details"},
		},
	}

	errorSubtests := []TestData{
		{
			name: "Json Unmarshal Error",
			requester: func(s string) ([]byte, error) {
				res := jsonRes[:len(jsonRes)-1]

				return res, nil
			},
			expectedError: Fr24Error{Err: "unexpected end of JSON input"},
		},
		{
			name: "Request Error",
			requester: func(s string) ([]byte, error) {
				return []byte{}, Fr24Error{"some error"}
			},
			expectedError: Fr24Error{"some error"},
		},
	}

	for _, subtest := range goodSubtests {
		t.Run(subtest.name, func(t *testing.T) {
			res, err := GetAirportBrief(subtest.requester, "LHR")

			if !errors.Is(err, nil) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Name != "London Heathrow Airport" {
				t.Errorf("Expected name (%v), got name (%v)", "London Heathrow Airport", res.Name)
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.name+"_CDN", func(t *testing.T) {
			res, err := GetAirportBrief(subtest.requester, "LHR")

			if !errors.Is(err, subtest.expectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.expectedError, err)
			}

			if res.Name != "" {
				t.Errorf("Expected name to be empty, got name (%v)", res.Name)
			}
		})
	}
}

func TestAirportBriefDataStructs(t *testing.T) {
	var airport AirportBriefResponse
	jsonData := loadJsonAirportBriefData(t)

	// unmarshal the json to an airport object
	if err := json.Unmarshal(jsonData, &airport); err != nil {
		t.Errorf("Error unmarshalling airport test data (%v)", err)
	}

	subtests := []airportJsonSubtest{
		{
			name:     "AirportBriefResponse",
			actual:   airport,
			expected: AirportBriefResponse{Details: airport.Details},
		},
		{
			name:   "AirportBriefDetails",
			actual: airport.Details,
			expected: AirportBriefDetails{
				Name: "London Heathrow Airport",
				Code: IataIcaoCode{
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
			name:   "AirportPosition",
			actual: airport.Details.Position,
			expected: AirportPosition{
				Latitude:  51.471626,
				Longitude: -0.467081,
				Altitude:  83,
				Country:   airport.Details.Position.Country,
				Region:    airport.Details.Position.Region,
			},
		},
		{
			name:   "AirportBriefStats",
			actual: airport.Details.Stats,
			expected: AirportBriefStats{
				Arrivals:   airport.Details.Stats.Arrivals,
				Departures: airport.Details.Stats.Departures,
			},
		},
		{
			name:   "ArrivalDepartureAggregateStats",
			actual: airport.Details.Stats.Arrivals,
			expected: ArrivalDepartureAggregateStats{
				DelayIndex: 0,
				DelayAvg:   nil,
				Total:      "7055",
				Hourly:     airport.Details.Stats.Arrivals.Hourly,
				Stats:      airport.Details.Stats.Arrivals.Stats,
			},
		},
		{
			name:   "Hourly",
			actual: airport.Details.Stats.Arrivals.Hourly,
			expected: Hourly{
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
		t.Run(subtest.name, func(t *testing.T) {
			if !reflect.DeepEqual(subtest.expected, subtest.actual) {
				t.Errorf("(%s) does not equal (%s)", subtest.expected, subtest.actual)
			}
		})
	}
}

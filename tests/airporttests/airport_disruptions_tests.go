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

func loadJsonAirportDisruptionData(t *testing.T) []byte {
	data, err := os.ReadFile("./testdata/airport_disruption.json")

	if err != nil {
		t.Error(err.Error())
	}

	return data
}

func TestGetAirportDisruptions(t *testing.T) {
	jsonRes := loadJsonAirportDisruptionData(t)

	goodSubtests := []tests.TestData{
		{
			Name: "No Error",
			Requester: func(s string) ([]byte, error) {
				return jsonRes, nil
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
			res, err := client.GetAirportDisruptions(subtest.Requester)

			if !errors.Is(err, nil) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if len(res) < 1 {
				t.Error("Expected 1 or more airport disruption reports to be returned, got 0.")
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.Name+"_CDN", func(t *testing.T) {
			res, err := client.GetAirportDisruptions(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if len(res) != 0 {
				t.Error("Expected no airport disruptions to be returned.")
			}
		})
	}
}

func TestAirportDisruptionDataStructs(t *testing.T) {
	var disruptions airports.AirportDistruptionApiResponse
	jsonData := loadJsonAirportDisruptionData(t)

	if err := json.Unmarshal(jsonData, &disruptions); err != nil {
		t.Errorf("Error unmarshalling airport test data (%v)", err)
	}

	subtests := []tests.JsonValidationTest{
		{
			Name:     "AirportDisruptionData",
			Actual:   disruptions.Data,
			Expected: airports.AirportDisruptionData{Rank: disruptions.Data.Rank},
		},
		{
			Name:   "AirportDisruptionRank",
			Actual: disruptions.Data.Rank[0],
			Expected: airports.AirportDisruptionRank{
				Airport:    disruptions.Data.Rank[0].Airport,
				Arrivals:   disruptions.Data.Rank[0].Arrivals,
				Departures: disruptions.Data.Rank[0].Departures,
			},
		},
		{
			Name:   "AirportDisruptionAirport",
			Actual: disruptions.Data.Rank[0].Airport,
			Expected: airports.AirportDisruptionAirport{
				Code:      disruptions.Data.Rank[0].Airport.Code,
				Name:      "Tijuana International Airport",
				City:      "Tijuana",
				Latitude:  32.541061,
				Longitude: -116.970001,
				Country:   disruptions.Data.Rank[0].Airport.Country,
				Continent: 6,
				Timezone:  disruptions.Data.Rank[0].Airport.Timezone,
				Weather:   disruptions.Data.Rank[0].Airport.Weather,
			},
		},
		{
			Name:   "AirportDisruptionWeather",
			Actual: disruptions.Data.Rank[0].Airport.Weather,
			Expected: airports.AirportDisruptionWeather{
				Temp: disruptions.Data.Rank[0].Airport.Weather.Temp,
				Wind: disruptions.Data.Rank[0].Airport.Weather.Wind,
				Sky:  disruptions.Data.Rank[0].Airport.Weather.Sky,
			},
		},
		{
			Name:   "Temperature",
			Actual: disruptions.Data.Rank[0].Airport.Weather.Temp,
			Expected: airports.Temperature{
				Celsius:    17,
				Fahrenheit: 63,
			},
		},
		{
			Name:   "AlphaCountry",
			Actual: disruptions.Data.Rank[0].Airport.Country,
			Expected: airports.AlphaCountry{
				Name:   "MEXICO",
				Alpha2: "MX",
				Alpha3: "MEX",
			},
		},
		{
			Name:   "AirportDisruptionLiveStats",
			Actual: disruptions.Data.Rank[0].Arrivals.Live,
			Expected: airports.AirportDisruptionLiveStats{
				Index:               5,
				AverageDelayMin:     84,
				Ontime:              3,
				Delayed:             5,
				DelayedPercentage:   0.63,
				Cancelled:           0,
				CancelledPercentage: 0,
				Trend:               "static",
			},
		},
		{
			Name:   "AirportDisruptionDailyStats",
			Actual: disruptions.Data.Rank[0].Arrivals.Today,
			Expected: airports.AirportDisruptionDailyStats{
				Total:               106,
				Delayed:             31,
				DelayedPercentage:   0.29,
				Cancelled:           13,
				CancelledPercentage: 0.12,
			},
		},
		{
			Name:   "AirportDisruptionArrivalDepartureStats",
			Actual: disruptions.Data.Rank[0].Arrivals,
			Expected: airports.AirportDisruptionArrivalDepartureStats{
				Live:      disruptions.Data.Rank[0].Arrivals.Live,
				Yesterday: disruptions.Data.Rank[0].Arrivals.Yesterday,
				Today:     disruptions.Data.Rank[0].Arrivals.Today,
				Tomorrow:  disruptions.Data.Rank[0].Arrivals.Tomorrow,
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

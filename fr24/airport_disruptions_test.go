package fr24

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"
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

	goodSubtests := []TestData{
		{
			name: "No Error",
			requester: func(s string) ([]byte, error) {
				return jsonRes, nil
			},
			expectedError: nil,
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
			res, err := GetAirportDisruptions(subtest.requester)

			if !errors.Is(err, nil) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if len(res) < 1 {
				t.Error("Expected 1 or more airport disruption reports to be returned, got 0.")
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.name+"_CDN", func(t *testing.T) {
			res, err := GetAirportDisruptions(subtest.requester)

			if !errors.Is(err, subtest.expectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.expectedError, err)
			}

			if len(res) != 0 {
				t.Error("Expected no airport disruptions to be returned.")
			}
		})
	}
}

func TestAirportDisruptionDataStructs(t *testing.T) {
	var disruptions AirportDistruptionApiResponse
	jsonData := loadJsonAirportDisruptionData(t)

	if err := json.Unmarshal(jsonData, &disruptions); err != nil {
		t.Errorf("Error unmarshalling airport test data (%v)", err)
	}

	subtests := []airportJsonSubtest{
		{
			name:     "AirportDisruptionData",
			actual:   disruptions.Data,
			expected: AirportDisruptionData{Rank: disruptions.Data.Rank},
		},
		{
			name:   "AirportDisruptionRank",
			actual: disruptions.Data.Rank[0],
			expected: AirportDisruptionRank{
				Airport:    disruptions.Data.Rank[0].Airport,
				Arrivals:   disruptions.Data.Rank[0].Arrivals,
				Departures: disruptions.Data.Rank[0].Departures,
			},
		},
		{
			name:   "AirportDisruptionAirport",
			actual: disruptions.Data.Rank[0].Airport,
			expected: AirportDisruptionAirport{
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
			name:   "AirportDisruptionWeather",
			actual: disruptions.Data.Rank[0].Airport.Weather,
			expected: AirportDisruptionWeather{
				Temp: disruptions.Data.Rank[0].Airport.Weather.Temp,
				Wind: disruptions.Data.Rank[0].Airport.Weather.Wind,
				Sky:  disruptions.Data.Rank[0].Airport.Weather.Sky,
			},
		},
		{
			name:   "Temperature",
			actual: disruptions.Data.Rank[0].Airport.Weather.Temp,
			expected: Temperature{
				Celsius:    17,
				Fahrenheit: 63,
			},
		},
		{
			name:   "AlphaCountry",
			actual: disruptions.Data.Rank[0].Airport.Country,
			expected: AlphaCountry{
				Name:   "MEXICO",
				Alpha2: "MX",
				Alpha3: "MEX",
			},
		},
		{
			name:   "AirportDisruptionLiveStats",
			actual: disruptions.Data.Rank[0].Arrivals.Live,
			expected: AirportDisruptionLiveStats{
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
			name:   "AirportDisruptionDailyStats",
			actual: disruptions.Data.Rank[0].Arrivals.Today,
			expected: AirportDisruptionDailyStats{
				Total:               106,
				Delayed:             31,
				DelayedPercentage:   0.29,
				Cancelled:           13,
				CancelledPercentage: 0.12,
			},
		},
		{
			name:   "AirportDisruptionArrivalDepartureStats",
			actual: disruptions.Data.Rank[0].Arrivals,
			expected: AirportDisruptionArrivalDepartureStats{
				Live:      disruptions.Data.Rank[0].Arrivals.Live,
				Yesterday: disruptions.Data.Rank[0].Arrivals.Yesterday,
				Today:     disruptions.Data.Rank[0].Arrivals.Today,
				Tomorrow:  disruptions.Data.Rank[0].Arrivals.Tomorrow,
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

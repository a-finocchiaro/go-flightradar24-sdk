package fr24

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"
)

type airportSubtest struct {
	name     string
	expected any
	actual   any
}

func loadJsonAirportData(t *testing.T) []byte {
	data, err := os.ReadFile("../testdata/airport_res.json")

	if err != nil {
		t.Error(err.Error())
	}

	return data
}

func TestGetAirport(t *testing.T) {
	jsonRes := loadJsonAirportData(t)

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
			res, err := GetAirport(subtest.requester, "LAX")

			if !errors.Is(err, nil) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Details.Name != "Los Angeles International Airport" {
				t.Errorf("Expected name (%v), got name (%v)", "Los Angeles International Airport", res.Details.Name)
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.name+"_CDN", func(t *testing.T) {
			res, err := GetAirport(subtest.requester, "LAX")

			if !errors.Is(err, subtest.expectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.expectedError, err)
			}

			if res.Details.Name != "" {
				t.Errorf("Expected name to be empty, got name (%v)", res.Details.Name)
			}
		})
	}
}

func TestAirportDataStructs(t *testing.T) {
	var airport AirportApiResponse
	jsonData := loadJsonAirportData(t)

	// unmarshal the json to an airport object
	if err := json.Unmarshal(jsonData, &airport); err != nil {
		t.Errorf("Error unmarshalling airport test data (%v)", err)
	}

	subtests := []airportSubtest{
		{
			name:     "AirportApiResponse",
			expected: AirportApiResponse{airport.Result},
			actual:   airport,
		},
		{
			name:     "AirportResult",
			actual:   airport.Result,
			expected: AirportResult{Request: airport.Result.Request, Response: airport.Result.Response},
		},
		{
			name:     "AirportResultRequest",
			actual:   airport.Result,
			expected: AirportResult{Request: airport.Result.Request, Response: airport.Result.Response},
		},
		{
			name:   "AirportRequest",
			actual: airport.Result.Request,
			expected: AirportRequest{
				Callback: "",
				Code:     "LAX",
				Device:   "",
				Fleet:    "",
				Format:   "json",
				Limit:    100,
				Page:     1,
				Pk:       "",
				Plugin: []string{
					"details",
					"runways",
					"satelliteImage",
					"schedule",
					"scheduledRoutesStatistics",
					"weather",
				},
				PluginSetting: AirportRequestPluginSetting{
					Schedule: AirportRequestPluginSettingSchedule{
						Mode:      nil,
						Timestamp: 1727645100,
					},
					SatelliteImage: AirportRequestPluginSettingSatelliteImage{
						Scale: 1,
					},
				},
				Token: "",
			},
		},
		{
			name:     "AirportRequestResponse",
			actual:   airport.Result.Response,
			expected: AirportRequestResponse{Airport: airport.Result.Response.Airport},
		},
		{
			name:     "AirportRequestResponseAirport",
			actual:   airport.Result.Response.Airport,
			expected: AirportRequestResponseAirport{PluginData: airport.Result.Response.Airport.PluginData},
		},
		{
			name:   "AirportPluginData",
			actual: airport.Result.Response.Airport.PluginData,
			expected: AirportPluginData{
				Details:                   airport.Result.Response.Airport.PluginData.Details,
				Schedule:                  airport.Result.Response.Airport.PluginData.Schedule,
				Weather:                   airport.Result.Response.Airport.PluginData.Weather,
				AircraftCount:             airport.Result.Response.Airport.PluginData.AircraftCount,
				Runways:                   airport.Result.Response.Airport.PluginData.Runways,
				ScheduledFlightStatistics: airport.Result.Response.Airport.PluginData.ScheduledFlightStatistics,
				SatelliteImage:            airport.Result.Response.Airport.PluginData.SatelliteImage,
				SatelliteImageProperties:  airport.Result.Response.Airport.PluginData.SatelliteImageProperties,
			},
		},
		{
			name:     "AirportDetailsCode",
			actual:   airport.Result.Response.Airport.PluginData.Details.Code,
			expected: AirportDetailsCode{Iata: "LAX", Icao: "KLAX"},
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

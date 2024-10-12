package airporttests

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/client"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/airports"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/tests"
)

type airportSubtests struct {
	tests.TestData
	plugins []string
}

func loadJsonAirportData(t *testing.T) []byte {
	data, err := os.ReadFile("./testdata/airport_detailed.json")

	if err != nil {
		t.Error(err.Error())
	}

	return data
}

func TestGetAirport(t *testing.T) {
	jsonRes := loadJsonAirportData(t)

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

	errorSubtests := []airportSubtests{
		{
			TestData: tests.TestData{
				Name: "Request Error",
				Requester: func(s string) ([]byte, error) {
					return []byte{}, common.Fr24Error{Err: "some error"}
				},
				ExpectedError: common.Fr24Error{Err: "some error"},
			},
			plugins: []string{"details"},
		},
		{
			TestData: tests.TestData{
				Name: "Invalid Plugin",
				Requester: func(s string) ([]byte, error) {
					return jsonRes, nil
				},
				ExpectedError: common.Fr24Error{Err: "Plugin badPlugin not supported."},
			},
			plugins: []string{"badPlugin"},
		},
	}

	for _, subtest := range goodSubtests {
		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetAirportDetails(subtest.Requester, "LAX", subtest.plugins)

			if !errors.Is(err, nil) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Details.Name != "Los Angeles International Airport" {
				t.Errorf("Expected Name (%v), got Name (%v)", "Los Angeles International Airport", res.Details.Name)
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.Name+"_CDN", func(t *testing.T) {
			res, err := client.GetAirportDetails(subtest.Requester, "LAX", subtest.plugins)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if res.Details.Name != "" {
				t.Errorf("Expected Name to be empty, got Name (%v)", res.Details.Name)
			}
		})
	}
}

func TestAirportDataStructs(t *testing.T) {
	var airport airports.AirportApiResponse
	jsonData := loadJsonAirportData(t)

	// unmarshal the json to an airport object
	if err := json.Unmarshal(jsonData, &airport); err != nil {
		t.Errorf("Error unmarshalling airport test data (%v)", err)
	}

	subtests := []tests.JsonValidationTest{
		{
			Name:     "AirportApiResponse",
			Expected: airports.AirportApiResponse{Result: airport.Result},
			Actual:   airport,
		},
		{
			Name:     "AirportResult",
			Actual:   airport.Result,
			Expected: airports.AirportResult{Request: airport.Result.Request, Response: airport.Result.Response},
		},
		{
			Name:     "AirportResultRequest",
			Actual:   airport.Result,
			Expected: airports.AirportResult{Request: airport.Result.Request, Response: airport.Result.Response},
		},
		{
			Name:   "AirportRequest",
			Actual: airport.Result.Request,
			Expected: airports.AirportRequest{
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
				PluginSetting: airports.AirportRequestPluginSetting{
					Schedule: airports.AirportRequestPluginSettingSchedule{
						Mode:      nil,
						Timestamp: 1727645100,
					},
					SatelliteImage: airports.AirportRequestPluginSettingSatelliteImage{
						Scale: 1,
					},
				},
				Token: "",
			},
		},
		{
			Name:     "AirportRequestResponse",
			Actual:   airport.Result.Response,
			Expected: airports.AirportRequestResponse{Airport: airport.Result.Response.Airport},
		},
		{
			Name:     "AirportRequestResponseAirport",
			Actual:   airport.Result.Response.Airport,
			Expected: airports.AirportRequestResponseAirport{PluginData: airport.Result.Response.Airport.PluginData},
		},
		{
			Name:   "AirportPluginData",
			Actual: airport.Result.Response.Airport.PluginData,
			Expected: airports.AirportPluginData{
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
	}

	for _, subtest := range subtests {
		t.Run(subtest.Name, func(t *testing.T) {
			if !reflect.DeepEqual(subtest.Expected, subtest.Actual) {
				t.Errorf("(%s) does not equal (%s)", subtest.Expected, subtest.Actual)
			}
		})
	}

}

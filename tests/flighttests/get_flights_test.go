package flighttests

import (
	"errors"
	"testing"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/client"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/tests"
)

type SpyFr24FeedData struct {
	Calls [][]byte
}

func (s *SpyFr24FeedData) UnmarshalJSON(data []byte) error {
	s.Calls = append(s.Calls, data)
	return nil
}

func TestGetFlights(t *testing.T) {
	goodSubTests := []tests.TestData{
		{
			Name: "No error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
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
					],
					"365564ca": [
						"ABFBBE",
						35.0530,
						-109.7382,
						296,
						51100,
						6,
						"",
						"F-KGUP3",
						"BALL",
						"N871TH",
						1722800891,
						"",
						"",
						"",
						0,
						0,
						"HBAL641",
						0,
						""
					]
				}`), nil
			},
			ExpectedError: nil,
		},
	}

	errorSubTests := []tests.TestData{
		{
			Name: "Request Error",
			Requester: func(s string) ([]byte, error) {
				return []byte(``), common.Fr24Error{Err: "Bad Request"}
			},
			ExpectedError: common.Fr24Error{Err: "Bad Request"},
		},
		{
			Name: "JSON unmarshal error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
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
				}`), nil
			},
			ExpectedError: common.Fr24Error{Err: "invalid character '}' after array element"},
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var feedSpy SpyFr24FeedData
			err := client.GetFlights(subtest.Requester, &feedSpy)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if len(feedSpy.Calls) != 1 {
				t.Errorf("Expected %d call, got %d", 1, len(feedSpy.Calls))
			}
		})
	}

	for _, subtest := range errorSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var feedSpy SpyFr24FeedData
			err := client.GetFlights(subtest.Requester, &feedSpy)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%s), got error (%v)", subtest.ExpectedError, err)
			}

			if len(feedSpy.Calls) != 0 {
				t.Errorf("Expected %d call, got %d", 0, len(feedSpy.Calls))
			}
		})
	}
}

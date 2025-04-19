package flighttests

import (
	"errors"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/tests"
)

func TestGetFlights(t *testing.T) {
	goodSubTests := []tests.TestData{
		{
			Name: "No error",
			Requester: func(s string) ([]byte, error) {
				return []byte(dummyFeedJsonData), nil
			},
			ExpectedError: nil,
		},
	}

	errorSubTests := []tests.TestData{
		{
			Name: "Request Error",
			Requester: func(s string) ([]byte, error) {
				return []byte{}, common.Fr24Error{Err: "Bad Request"}
			},
			ExpectedError: common.Fr24Error{Err: "Bad Request"},
		},
		{
			Name: "JSON unmarshal error",
			Requester: func(s string) ([]byte, error) {
				return []byte(dummyFeedJsonDataBad), nil
			},
			ExpectedError: common.Fr24Error{Err: "invalid character '}' after array element"},
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			feedData, err := client.GetFlights(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if len(feedData.Flights) != 2 {
				t.Errorf("Expected 2 flights to be unpacked, received (%d)", len(feedData.Flights))
			}
		})
	}

	for _, subtest := range errorSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			feedData, err := client.GetFlights(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%s), got error (%v)", subtest.ExpectedError, err)
			}

			if len(feedData.Flights) != 0 {
				t.Errorf("Expected %d call, got %d", 0, len(feedData.Flights))
			}
		})
	}
}

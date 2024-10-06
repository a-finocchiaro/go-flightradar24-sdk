package internal

import (
	"errors"
	"fmt"
	"testing"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/tests"
)

type DummyJson struct {
	Model      string `json:"model"`
	TailNumber string `json:"tailno"`
}

func TestSendRequest(t *testing.T) {
	goodSubTests := []tests.TestData{
		{
			Name: "No Error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
					"model": "B747-8",
					"tailno": "N123AB"
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
			Name: "JSON Unmarshal Error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
					"model": "B747-8",
					"tailno": "N123AB"
				`), nil
			},
			ExpectedError: common.Fr24Error{Err: "unexpected end of JSON input"},
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var dummyData DummyJson
			err := SendRequest(subtest.Requester, "", &dummyData)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if dummyData.Model != "B747-8" {
				t.Errorf("Expected model (%s), received %s", "B747-8", dummyData.Model)
			}

			if dummyData.TailNumber != "N123AB" {
				t.Errorf("Expected model (%s), received %s", "N123AB", dummyData.TailNumber)
			}
		})
	}

	for _, subtest := range errorSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var dummyData DummyJson
			err := SendRequest(subtest.Requester, "", &dummyData)

			fmt.Println(dummyData)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%s), got error (%v)", subtest.ExpectedError, err)
			}
		})
	}
}

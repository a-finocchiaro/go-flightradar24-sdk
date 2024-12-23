package geographytests

import (
	"errors"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/tests"
)

func TestGetZone(t *testing.T) {
	goodSubTests := []tests.TestData{
		{
			Name: "No Error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
					"version": 4,
					"europe": {
						"tl_y": 72.57,
						"tl_x": -16.96,
						"br_y": 33.57,
						"br_x": 53.05,
						"subzones": {
							"uk": {
								"tl_y": 62.61,
								"tl_x": -13.07,
								"br_y": 49.71,
								"br_x": 3.46,
								"subzones": {
									"london": {
										"tl_y": 53.06,
										"tl_x": -2.87,
										"br_y": 50.07,
										"br_x": 3.26
                    				}
								}
							}
						}
					}
				}`), nil
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

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.GetZones(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if res.Version != 4 {
				t.Errorf("Version: got %d, want %d", res.Version, 1727163107)
			}

			if res.Europe.TlY != 72.57 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.TlY, 72.57)
			}

			if res.Europe.TlX != -16.96 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.TlX, -16.96)
			}

			if res.Europe.BrY != 33.57 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.BrY, 33.57)
			}

			if res.Europe.BrX != 53.05 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.BrX, 53.05)
			}

			if res.Europe.Subzones.Uk.TlY != 62.61 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.TlY, 62.61)
			}

			if res.Europe.Subzones.Uk.TlX != -13.07 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.TlX, -13.07)
			}

			if res.Europe.Subzones.Uk.BrY != 49.71 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.BrY, 49.71)
			}

			if res.Europe.Subzones.Uk.BrX != 3.46 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.BrX, 3.46)
			}

			if res.Europe.Subzones.Uk.Subzones.London.TlY != 53.06 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.Subzones.London.TlY, 53.06)
			}

			if res.Europe.Subzones.Uk.Subzones.London.TlX != -2.87 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.Subzones.London.TlX, -2.87)
			}

			if res.Europe.Subzones.Uk.Subzones.London.BrY != 50.07 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.Subzones.London.BrY, 50.07)
			}

			if res.Europe.Subzones.Uk.Subzones.London.BrX != 3.26 {
				t.Errorf("Rows: got length %f, want %f", res.Europe.Subzones.Uk.Subzones.London.BrX, 3.26)
			}
		})
	}

	for _, subtest := range errorSubtests {
		t.Run(subtest.Name+"_CDN", func(t *testing.T) {
			res, err := client.GetZones(subtest.Requester)

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if res.Europe.BrX != 0.0 {
				t.Errorf("Expected Name to be (0.0), got Name (%v)", res.Europe.BrX)
			}
		})
	}
}

package geographytests

import (
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
)

type TestGetBoundsByPointData struct {
	Name      string
	Latitude  float64
	Longitude float64
	Radius    float64
	Expected  string
}

func TestGetBoundsByPoint(t *testing.T) {
	goodSubTests := []TestGetBoundsByPointData{
		{
			Name:      "No Error",
			Latitude:  32.918559,
			Longitude: -97.058446,
			Radius:    500,
			Expected:  "32.923055,32.914062,-97.063802,-97.053089",
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			res := client.GetBoundsByPoint(subtest.Latitude, subtest.Longitude, subtest.Radius)

			if res != subtest.Expected {
				t.Errorf("Expected bounding box of (%s), got (%s)", subtest.Expected, res)
			}
		})
	}
}

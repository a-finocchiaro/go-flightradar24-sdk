package fr24

import (
	"errors"
	"testing"
)

type SpyFr24FeedData struct {
	Calls [][]byte
}

func (s *SpyFr24FeedData) UnmarshalJSON(data []byte) error {
	s.Calls = append(s.Calls, data)
	return nil
}

func TestGetFlights(t *testing.T) {
	goodSubTests := []TestData{
		{
			name: "No error",
			requester: func(s string) ([]byte, error) {
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
			expectedError: nil,
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.name, func(t *testing.T) {
			var feedSpy SpyFr24FeedData
			err := GetFlights(subtest.requester, &feedSpy)

			if !errors.Is(err, subtest.expectedError) {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if len(feedSpy.Calls) != 1 {
				t.Errorf("Expected %d call, got %d", 1, len(feedSpy.Calls))
			}
		})
	}
}

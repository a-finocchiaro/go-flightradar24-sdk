package modeltests

import (
	"errors"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/client"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/tests"
)

func TestRouteUnmarshal(t *testing.T) {
	goodSubTests := []tests.TestData{
		{
			Name: "No error",
			Requester: func(s string) ([]byte, error) {
				return []byte(`{
								"results": [
									{
										"id": "388a5b54",
										"label": "AY20 / FIN20 / A359 (OH-LWR)",
										"detail": {
											"lat": 33.8,
											"lon": -96.6,
											"schd_from": "DFW",
											"schd_to": "HEL",
											"ac_type": "A359",
											"route": "Dallas (DFW) ⟶ Helsinki (HEL)",
											"logo": "https://images.flightradar24.com/assets/airlines/logotypes/21.png",
											"reg": "OH-LWR",
											"callsign": "FIN20",
											"flight": "AY20",
											"operator": "FIN",
											"operator_id": 21
										},
										"type": "live",
										"match": "begins"
									},
									{
										"id": "OH-LWR",
										"label": "OH-LWR (A359) - FIN",
										"detail": {
											"owner": "FIN",
											"equip": "A359",
											"hex": "461F59",
											"operator_id": 21,
											"logo": "https://images.flightradar24.com/assets/airlines/logotypes/21.png"
										},
										"type": "aircraft",
										"match": "begins"
									}
								],
								"stats": {
									"total": {
										"all": 2,
										"airport": 0,
										"operator": 0,
										"live": 1,
										"schedule": 0,
										"aircraft": 1
									},
									"count": {
										"airport": 0,
										"operator": 0,
										"live": 1,
										"schedule": 0,
										"aircraft": 1
									}
								}
							}`), nil
			},
			ExpectedError: nil,
		},
	}

	errorSubTests := []tests.TestData{
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
			res, err := client.Search(subtest.Requester, "OH-LWR")

			if err != nil {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			detail := res.Results[0].Detail

			if len(res.Results) != 2 {
				t.Errorf("Expected Results count to be 2, got Name (%d)", len(res.Results))
			}

			if res.Results[0].ID != "388a5b54" {
				t.Errorf("Expected Id to be 388a5b54, got ID (%s)", res.Results[0].ID)
			}

			if res.Results[0].Label != "AY20 / FIN20 / A359 (OH-LWR)" {
				t.Errorf("Expected Label to be AY20 / FIN20 / A359 (OH-LWR), got  (%s)", res.Results[0].Label)
			}

			if detail.Lat != 33.8 {
				t.Errorf("Expected Detail to be 33.8, got  (%f)", detail.Lat)
			}

			if detail.Lon != -96.6 {
				t.Errorf("Expected Detail to be -96.6, got  (%f)", detail.Lon)
			}

			if detail.SchdFrom != "DFW" {
				t.Errorf("Expected Detail to be DFW, got  (%s)", detail.SchdFrom)
			}

			if detail.SchdTo != "HEL" {
				t.Errorf("Expected Detail to be HEL, got  (%s)", detail.SchdTo)
			}

			if detail.AcType != "A359" {
				t.Errorf("Expected Detail to be A359, got  (%s)", detail.AcType)
			}

			if detail.Route != "Dallas (DFW) ⟶ Helsinki (HEL)" {
				t.Errorf("Expected Detail to be Dallas (DFW) ⟶ Helsinki (HEL), got  (%s)", detail.Route)
			}

			if detail.Logo != "https://images.flightradar24.com/assets/airlines/logotypes/21.png" {
				t.Errorf("Expected Detail to be https://images.flightradar24.com/assets/airlines/logotypes/21.png, got  (%s)", detail.Logo)
			}

			if detail.Reg != "OH-LWR" {
				t.Errorf("Expected Detail to be OH-LWR, got  (%s)", detail.Reg)
			}

			if detail.Callsign != "FIN20" {
				t.Errorf("Expected Detail to be FIN20, got  (%s)", detail.Callsign)
			}

			if detail.Flight != "AY20" {
				t.Errorf("Expected Detail to be AY20, got  (%s)", detail.Flight)
			}

			if detail.Operator != "FIN" {
				t.Errorf("Expected Detail to be FIN, got  (%s)", detail.Operator)
			}

			if res.Results[0].Detail.OperatorID != 21 {
				t.Errorf("Expected Operator ID to be 21, got  (%d)", detail.OperatorID)
			}

			if res.Stats.Total.All != 2 {
				t.Errorf("Expected Stats total all to be 2, got  (%d)", res.Stats.Total.All)
			}

			if res.Stats.Total.Airport != 0 {
				t.Errorf("Expected Stats total airport to be 0, got  (%d)", res.Stats.Total.Airport)
			}

			if res.Stats.Total.Operator != 0 {
				t.Errorf("Expected Stats total operator to be 0, got  (%d)", res.Stats.Total.Operator)
			}

			if res.Stats.Total.Live != 1 {
				t.Errorf("Expected Stats total live to be 1, got  (%d)", res.Stats.Total.Live)
			}

			if res.Stats.Total.Schedule != 0 {
				t.Errorf("Expected Stats total schedule to be 0, got  (%d)", res.Stats.Total.Schedule)
			}

			if res.Stats.Total.Aircraft != 1 {
				t.Errorf("Expected Stats total aircraft to be 1, got  (%d)", res.Stats.Total.Aircraft)
			}

			if res.Stats.Count.Airport != 0 {
				t.Errorf("Expected Stats count airport to be 0, got  (%d)", res.Stats.Count.Airport)
			}

			if res.Stats.Count.Operator != 0 {
				t.Errorf("Expected Stats count operator to be 0, got  (%d)", res.Stats.Count.Operator)
			}

			if res.Stats.Count.Live != 1 {
				t.Errorf("Expected Stats count live to be 1, got  (%d)", res.Stats.Count.Live)
			}

			if res.Stats.Count.Schedule != 0 {
				t.Errorf("Expected Stats count schedule to be 0, got  (%d)", res.Stats.Count.Schedule)
			}

			if res.Stats.Count.Aircraft != 1 {
				t.Errorf("Expected Stats count aircraft to be 1, got  (%d)", res.Stats.Count.Aircraft)
			}
		})
	}

	for _, subtest := range errorSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			res, err := client.Search(subtest.Requester, "LHR")

			if !errors.Is(err, subtest.ExpectedError) {
				t.Errorf("Expected error (%v), got error (%v)", subtest.ExpectedError, err)
			}

			if len(res.Results) != 0 {
				t.Errorf("Expected Results count to be 0, got Name (%d)", len(res.Results))
			}
		})
	}
}

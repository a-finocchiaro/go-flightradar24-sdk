package modeltests

import (
	"encoding/json"
	"testing"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/airports"
)

func TestRouteUnmarshal(t *testing.T) {
	goodSubTests := []FlightFeedTestData{
		{
			Name: "No error",
			JsonData: []byte(`{
								"arrivals": {
									"United States": {
										"number": {
											"airports": 1,
											"flights": 2
										},
										"airports": {
											"SAN": {
												"name": "San Diego International Airport",
												"city": "San Diego",
												"icao": "KSAN",
												"flights": {
													"WN1322": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-19": {
																"aircraft": "B737",
																"time": "15:25",
																"timestamp": 1734621900,
																"offset": -28800
															}
														}
													},
													"WN2053": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-24": {
																"aircraft": "73W",
																"time": "16:50",
																"timestamp": 1735059000,
																"offset": -28800
															},
															"2024-12-25": {
																"aircraft": "73W",
																"time": "16:50",
																"timestamp": 1735145400,
																"offset": -28800
															}
														}
													}
												},
												"position": {
													"lat": "32.733551",
													"lon": "-117.189003"
												}
											}
										}
									}
								},
								"departures": {
									"United States": {
										"number": {
											"airports": 1,
											"flights": 2
										},
										"airports": {
											"SAN": {
												"name": "San Diego International Airport",
												"city": "San Diego",
												"icao": "KSAN",
												"flights": {
													"WN200": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-21": {
																"aircraft": "73W",
																"time": "04:55",
																"timestamp": 1734756900,
																"offset": -28800
															},
															"2024-12-23": {
																"aircraft": "73W",
																"time": "04:55",
																"timestamp": 1734929700,
																"offset": -28800
															},
															"2024-12-24": {
																"aircraft": "73W",
																"time": "04:55",
																"timestamp": 1735016100,
																"offset": -28800
															}
														}
													},
													"WN2568": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-20": {
																"aircraft": "73W",
																"time": "05:15",
																"timestamp": 1734671700,
																"offset": -28800
															}
														}
													}
												},
												"position": {
													"lat": "32.733551",
													"lon": "-117.189003"
												}
											}
										}
									}
								}
							}`),
		},
	}

	errorSubTests := []FlightFeedTestData{
		{
			Name: "Json Unmarshal Error",
			JsonData: []byte(`{
								"arrivals": {
									"United States": {
										"number": {
											"airports": 1
											"flights": 2
										}
									}
								}
							}`),
		},
	}

	for _, subtest := range goodSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var route airports.AirportRouteResponse

			err := json.Unmarshal(subtest.JsonData, &route)

			if err != nil {
				t.Errorf("Expected no errors, got error (%v)", err)
			}

			if route.Arrivals.Number.Airports != 1 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 1, route.Arrivals.Number.Airports)
			}

			if route.Arrivals.Number.Flights != 2 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 2, route.Arrivals.Number.Flights)
			}

			if route.Arrivals.Airports.Name != "San Diego International Airport" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "San Diego International Airport", route.Arrivals.Airports.Name)

			}

			if route.Arrivals.Airports.City != "San Diego" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "San Diego", route.Arrivals.Airports.City)
			}

			if route.Arrivals.Airports.Icao != "KSAN" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "KSAN", route.Arrivals.Airports.Icao)
			}

			if route.Arrivals.Airports.Iata != "SAN" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "SAN", route.Arrivals.Airports.Iata)
			}

			if len(route.Arrivals.Airports.Flights) != 2 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 2, len(route.Arrivals.Airports.Flights))
			}

			if route.Arrivals.Airports.Position.Latitude != "32.733551" {
				t.Errorf("Expected Latitude to be (%s), got (%s)", "32.733551", route.Arrivals.Airports.Position.Latitude)
			}

			if route.Arrivals.Airports.Position.Longitude != "-117.189003" {
				t.Errorf("Expected Longitude to be (%s), got (%s)", "-117.189003", route.Arrivals.Airports.Position.Longitude)
			}

			// check structure of a single flight to verify it's correct
			flight := route.Arrivals.Airports.Flights[0]

			if flight.ID != "WN1322" {
				t.Errorf("Expected Flight ID to be (%s), got (%s)", "WN1322", flight.ID)
			}

			if flight.Airline.Name != "Southwest Airlines" {
				t.Errorf("Expected Flight Airline Name to be (%s), got (%s)", "WN", flight.Airline.Name)
			}

			if flight.Airline.Iata != "WN" {
				t.Errorf("Expected Flight Airline IATA to be (%s), got (%s)", "WN", flight.Airline.Iata)
			}

			if flight.Airline.Icao != "SWA" {
				t.Errorf("Expected Flight Airline ICAO to be (%s), got (%s)", "SWA", flight.Airline.Icao)
			}

			if flight.Airline.Url != "wn-swa" {
				t.Errorf("Expected Flight Airline Url to be (%s), got (%s)", "wn-swa", flight.Airline.Url)
			}

			if flight.Utc[0].Date != "2024-12-19" {
				t.Errorf("Expected Flight UTC date to be (%s), got (%s)", "2024-12-19", flight.Utc[0].Date)
			}

			if flight.Utc[0].Aircraft != "B737" {
				t.Errorf("Expected Flight UTC date to be (%s), got (%s)", "B737", flight.Utc[0].Aircraft)
			}

			if flight.Utc[0].Time != "15:25" {
				t.Errorf("Expected Flight UTC date to be (%s), got (%s)", "15:25", flight.Utc[0].Time)
			}

			if flight.Utc[0].Timestamp != 1734621900 {
				t.Errorf("Expected Flight UTC date to be (%d), got (%d)", 1734621900, flight.Utc[0].Timestamp)
			}

			if flight.Utc[0].Offset != -28800 {
				t.Errorf("Expected Flight UTC date to be (%d), got (%d)", -28800, flight.Utc[0].Offset)
			}

			// departures
			if route.Departures.Number.Airports != 1 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 1, route.Departures.Number.Airports)
			}

			if route.Departures.Number.Flights != 2 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 2, route.Departures.Number.Flights)
			}

			if route.Departures.Airports.Name != "San Diego International Airport" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "San Diego International Airport", route.Departures.Airports.Name)

			}

			if route.Departures.Airports.City != "San Diego" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "San Diego", route.Departures.Airports.City)
			}

			if route.Departures.Airports.Icao != "KSAN" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "KSAN", route.Departures.Airports.Icao)
			}

			if route.Departures.Airports.Iata != "SAN" {
				t.Errorf("Expected FullCount to be (%s), got (%s)", "SAN", route.Departures.Airports.Iata)
			}

			if len(route.Departures.Airports.Flights) != 2 {
				t.Errorf("Expected FullCount to be (%d), got (%d)", 2, len(route.Departures.Airports.Flights))
			}

			if route.Departures.Airports.Position.Latitude != "32.733551" {
				t.Errorf("Expected Latitude to be (%s), got (%s)", "32.733551", route.Departures.Airports.Position.Latitude)
			}

			if route.Departures.Airports.Position.Longitude != "-117.189003" {
				t.Errorf("Expected Longitude to be (%s), got (%s)", "-117.189003", route.Departures.Airports.Position.Longitude)
			}
		})
	}

	for _, subtest := range errorSubTests {
		t.Run(subtest.Name, func(t *testing.T) {
			var route airports.AirportRouteResponse
			err := json.Unmarshal(subtest.JsonData, &route)

			if err == nil {
				t.Error("Expected an error to be returned, got nil")
			}

			if route.Arrivals.Country != "" {
				t.Errorf("Expected empty Country name, got (%s)", route.Arrivals.Country)
			}
		})
	}
}

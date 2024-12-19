package modeltests

import (
	"encoding/json"
	"testing"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/airports"
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
											"flights": 6
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
													},
													"WN2533": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-21": {
																"aircraft": "7M8",
																"time": "02:00",
																"timestamp": 1734746400,
																"offset": -28800
															},
															"2024-12-23": {
																"aircraft": "7M8",
																"time": "02:00",
																"timestamp": 1734919200,
																"offset": -28800
															},
															"2024-12-24": {
																"aircraft": "7M8",
																"time": "02:00",
																"timestamp": 1735005600,
																"offset": -28800
															}
														}
													},
													"WN2879": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-20": {
																"aircraft": "73H",
																"time": "15:05",
																"timestamp": 1734707100,
																"offset": -28800
															},
															"2024-12-22": {
																"aircraft": "73H",
																"time": "15:05",
																"timestamp": 1734879900,
																"offset": -28800
															},
															"2024-12-23": {
																"aircraft": "73H",
																"time": "15:05",
																"timestamp": 1734966300,
																"offset": -28800
															},
															"2024-12-26": {
																"aircraft": "73H",
																"time": "15:05",
																"timestamp": 1735225500,
																"offset": -28800
															}
														}
													},
													"WN3524": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-20": {
																"aircraft": "73W",
																"time": "01:30",
																"timestamp": 1734658200,
																"offset": -28800
															}
														}
													},
													"WN4570": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-21": {
																"aircraft": "73W",
																"time": "20:20",
																"timestamp": 1734812400,
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
											"flights": 6
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
													},
													"WN2801": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-21": {
																"aircraft": "73H",
																"time": "23:30",
																"timestamp": 1734823800,
																"offset": -28800
															}
														}
													},
													"WN3143": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-20": {
																"aircraft": "7M8",
																"time": "18:45",
																"timestamp": 1734720300,
																"offset": -28800
															},
															"2024-12-22": {
																"aircraft": "7M8",
																"time": "18:45",
																"timestamp": 1734893100,
																"offset": -28800
															},
															"2024-12-23": {
																"aircraft": "7M8",
																"time": "18:45",
																"timestamp": 1734979500,
																"offset": -28800
															},
															"2024-12-26": {
																"aircraft": "7M8",
																"time": "18:45",
																"timestamp": 1735238700,
																"offset": -28800
															}
														}
													},
													"WN3563": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-19": {
																"aircraft": "B38M",
																"time": "17:40",
																"timestamp": 1734630000,
																"offset": -28800
															}
														}
													},
													"WN3875": {
														"airline": {
															"name": "Southwest Airlines",
															"iata": "WN",
															"icao": "SWA",
															"url": "wn-swa"
														},
														"utc": {
															"2024-12-19": {
																"aircraft": "B738",
																"time": "01:45",
																"timestamp": 1734572700,
																"offset": -28800
															},
															"2024-12-25": {
																"aircraft": "7M8",
																"time": "01:45",
																"timestamp": 1735091100,
																"offset": -28800
															},
															"2024-12-26": {
																"aircraft": "7M8",
																"time": "01:45",
																"timestamp": 1735177500,
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
		})
	}
}

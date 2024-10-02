package fr24

import (
	"encoding/json"
	"fmt"
)

type AirportBriefResponse struct {
	Details AirportBriefDetails `json:"details"`
}

type AirportBriefDetails struct {
	Name     string                  `json:"name"`
	Code     IataIcaoCode            `json:"code"`
	Position AirportPosition         `json:"position"`
	Timezone TimezoneWithOffsetHours `json:"timezone"`
	Visible  bool                    `json:"visible"`
	Website  string                  `json:"website"`
	Stats    AirportBriefStats       `json:"stats"`
}

type CountryExtended struct {
	Country
	CodeLong string `json:"codeLong"`
}

type TimezoneWithOffsetHours struct {
	Timezone
	OffsetHours string `json:"offsetHours"`
}

type AirportPosition struct {
	Latitude  float64              `json:"latitude"`
	Longitude float64              `json:"longitude"`
	Altitude  int                  `json:"altitude"`
	Country   CountryExtended      `json:"country"`
	Region    AirportDetailsRegion `json:"region"`
}

type AirportBriefStats struct {
	Arrivals   ArrivalDepartureAggregateStats `json:"arrivals"`
	Departures ArrivalDepartureAggregateStats `json:"departures"`
}

type ArrivalDepartureAggregateStats struct {
	DelayIndex int      `json:"delayIndex"`
	DelayAvg   any      `json:"delayAvg"`
	Total      string   `json:"total"`
	Hourly     Hourly   `json:"hourly"`
	Stats      []string `json:"stats"`
}

type Hourly struct {
	Hour0  string `json:"0"`
	Hour1  string `json:"1"`
	Hour2  string `json:"2"`
	Hour3  string `json:"3"`
	Hour4  string `json:"4"`
	Hour5  string `json:"5"`
	Hour6  string `json:"6"`
	Hour7  string `json:"7"`
	Hour8  string `json:"8"`
	Hour9  string `json:"9"`
	Hour10 string `json:"10"`
	Hour11 string `json:"11"`
	Hour12 string `json:"12"`
	Hour13 string `json:"13"`
	Hour14 string `json:"14"`
	Hour15 string `json:"15"`
	Hour16 string `json:"16"`
	Hour17 string `json:"17"`
	Hour18 string `json:"18"`
	Hour19 string `json:"19"`
	Hour20 string `json:"20"`
	Hour21 string `json:"21"`
	Hour22 string `json:"22"`
	Hour23 string `json:"23"`
}

// Gets brief airport information from the /airports/traffic-stats endpoint from the FR24
// base URL.
// Acceps an airport IATA or ICAO code as an argument for code.
func GetAirportBrief(requester Requester, code string) (AirportBriefDetails, error) {
	var airport AirportBriefResponse

	endpoint := fmt.Sprintf("%s?airport=%s", FR24_ENDPOINTS["airport_brief"], code)

	body, err := requester(endpoint)

	if err != nil {
		return airport.Details, NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &airport); err != nil {
		return airport.Details, NewFr24Error(err)
	}

	return airport.Details, nil
}

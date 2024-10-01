package fr24

import (
	"encoding/json"
)

type AirportDistruptionApiResponse struct {
	Success bool                  `json:"success"`
	Meta    any                   `json:"meta,omitempty"`
	Data    AirportDisruptionData `json:"data"`
}

type AirportDisruptionData struct {
	Rank []AirportDisruptionRank `json:"rank"`
}

type AirportDisruptionRank struct {
	Airport    AirportDisruptionAirport               `json:"airport"`
	Arrivals   AirportDisruptionArrivalDepartureStats `json:"arrivals"`
	Departures AirportDisruptionArrivalDepartureStats `json:"departures"`
}

type AirportDisruptionAirport struct {
	Code      IataIcaoCode             `json:"code"`
	Name      string                   `json:"name"`
	City      string                   `json:"city"`
	Latitude  float64                  `json:"latitude"`
	Longitude float64                  `json:"longitude"`
	Country   AlphaCountry             `json:"country"`
	Continent int                      `json:"continent"`
	Timezone  Timezone                 `json:"timezone"`
	Weather   AirportDisruptionWeather `json:"weather"`
}

type AirportDisruptionWeather struct {
	Temp Temperature `json:"temp"`
	Wind AirportWind `json:"wind"`
	Sky  struct {
		Condition AirportCondition `json:"condition"`
	} `json:"sky"`
}

type Temperature struct {
	Celsius    int `json:"celsius"`
	Fahrenheit int `json:"fahrenheit"`
}

type AlphaCountry struct {
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
}

type AirportDisruptionLiveStats struct {
	Index               float32 `json:"index"`
	AverageDelayMin     int     `json:"averageDelayMin"`
	Ontime              int     `json:"ontime"`
	Delayed             int     `json:"delayed"`
	DelayedPercentage   float64 `json:"delayedPercentage"`
	Cancelled           int     `json:"cancelled"`
	CancelledPercentage int     `json:"cancelledPercentage"`
	Trend               string  `json:"trend"`
}

type AirportDisruptionDailyStats struct {
	Total               int     `json:"total"`
	Delayed             int     `json:"delayed"`
	DelayedPercentage   float64 `json:"delayedPercentage"`
	Cancelled           int     `json:"cancelled"`
	CancelledPercentage float64 `json:"cancelledPercentage"`
}

type AirportDisruptionArrivalDepartureStats struct {
	Live      AirportDisruptionLiveStats  `json:"live"`
	Yesterday AirportDisruptionDailyStats `json:"yesterday"`
	Today     AirportDisruptionDailyStats `json:"today"`
	Tomorrow  AirportDisruptionDailyStats `json:"tomorrow"`
}

func GetAirportDisruptions(requester Requester) ([]AirportDisruptionRank, error) {
	var disruptions AirportDistruptionApiResponse
	body, err := requester(FR24_ENDPOINTS["airport_disruptions"])

	if err != nil {
		return disruptions.Data.Rank, NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &disruptions); err != nil {
		return disruptions.Data.Rank, NewFr24Error(err)
	}

	return disruptions.Data.Rank, nil
}

/*
Logic used to query data fom FR24 to check various stats.
*/
package fr24

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

const FR24_BASE = "https://www.flightradar24.com"

type Fr24MostTrackedRes struct {
	Version     string                `json:"version"`
	Update_time float64               `json:"update_time"`
	Data        []Fr24MostTrackedData `json:"data"`
}

type Fr24MostTrackedData struct {
	Flight_id     string `json:"flight_id"`
	Flight        string `json:"flight"`
	Callsign      string `json:"callsign"`
	Squawk        string `json:"squawk"`
	Clicks        int    `json:"clicks"`
	From_iata     string `json:"from_iata"`
	From_city     string `json:"from_city"`
	To_iata       string `json:"to_iata"`
	To_city       string `json:"to_city"`
	Model         string `json:"model"`
	Aircraft_type string `json:"type"`
}

var FR24_ENDPOINTS = map[string]string{
	"most_tracked": fmt.Sprintf("%s/flights/most-tracked", FR24_BASE),
}

type (
	Requester func(string) ([]byte, error)
)

var ErrUnmarshall = errors.New("could not unmarshall json")

func GetFR24MostTracked(requester Requester) (Fr24MostTrackedRes, error) {
	var most_tracked Fr24MostTrackedRes
	body, err := requester(FR24_ENDPOINTS["most_tracked"])

	if err != nil {
		log.Fatalln(err)
		return most_tracked, err
	}

	if err := json.Unmarshal(body, &most_tracked); err != nil {
		return most_tracked, ErrUnmarshall
	}

	return most_tracked, nil
}

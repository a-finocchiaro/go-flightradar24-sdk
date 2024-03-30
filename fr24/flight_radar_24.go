/*
Logic used to query data fom FR24 to check various stats.
*/
package fr24

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/a-finocchiaro/adsb_flightradar_top10/webrequest"
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

func GetFR24MostTracked() Fr24MostTrackedRes {
	var most_tracked Fr24MostTrackedRes
	body, err := webrequest.SendRequest(FR24_ENDPOINTS["most_tracked"])

	if err != nil {
		log.Fatalln("Request failed.")
	}

	if err := json.Unmarshal(body, &most_tracked); err != nil {
		log.Fatalln("Could not unmarshal most tracked data into an Fr24MostTrackedRes object.")
	}

	return most_tracked
}

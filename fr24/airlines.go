package fr24

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
)

type AirlineRes struct {
	Version int           `json:"version"`
	Rows    []AirlineData `json:"rows"`
}

type AirlineData struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
	Icao string `json:"ICAO"`
}

func GetAirlines(requester Requester) (AirlineRes, error) {
	var airlines AirlineRes

	body, err := requester(FR24_ENDPOINTS["airlines"])

	if err != nil {
		return airlines, NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &airlines); err != nil {
		return airlines, NewFr24Error(err)
	}

	return airlines, nil
}

func GetAirlineLogo(requester Requester, icao string, iata string) (bytes.Buffer, error) {
	/**
	* Gets a logo for a requested airline based on its icao and iata code.
	 */
	var buf bytes.Buffer
	endpoint := fmt.Sprintf("%s/%s_%s.png", FR24_ENDPOINTS["airline_logo_cdn"], icao, iata)

	body, err := requester(endpoint)

	if err != nil {
		return buf, NewFr24Error(err)
	}

	img, _, err := image.Decode(bytes.NewReader(body))

	if err != nil {
		return buf, NewFr24Error(err)
	}

	// encode the bytes into a png image
	if err := png.Encode(&buf, img); err != nil {
		return buf, NewFr24Error(err)
	}

	return buf, nil
}

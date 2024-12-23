package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"

	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/airlines"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
)

func GetAirlines(requester common.Requester) (airlines.AirlineRes, error) {
	var airlines airlines.AirlineRes

	body, err := requester(common.FR24_ENDPOINTS["airlines"])

	if err != nil {
		return airlines, common.NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &airlines); err != nil {
		return airlines, common.NewFr24Error(err)
	}

	return airlines, nil
}

func GetAirlineLogoCdn(requester common.Requester, icao string, iata string) (bytes.Buffer, error) {
	/**
	* Gets a logo for a requested airline based on its icao and iata code from the CDN.
	 */
	var buf bytes.Buffer
	endpoint := fmt.Sprintf("%s/%s_%s.png", common.FR24_ENDPOINTS["airline_logo_cdn"], icao, iata)

	if err := createPng(&buf, endpoint, requester); err != nil {
		return buf, err
	}

	return buf, nil
}

func GetAirlineLogo(requester common.Requester, icao string) (bytes.Buffer, error) {
	/**
	* Get Logo from assets on Flightradar
	 */
	var buf bytes.Buffer
	endpoint := fmt.Sprintf("%s/%s_logo0.png", common.FR24_ENDPOINTS["airline_logo"], icao)

	if err := createPng(&buf, endpoint, requester); err != nil {
		return buf, err
	}

	return buf, nil
}

func createPng(buf *bytes.Buffer, endpoint string, requester common.Requester) error {
	/**
	* Sends the request to Flightradar24 for a PNG image and encodes the returned bytes into
	* a PNG logo.
	 */
	body, err := requester(endpoint)

	if err != nil {
		return common.NewFr24Error(err)
	}

	img, _, err := image.Decode(bytes.NewReader(body))

	if err != nil {
		return common.NewFr24Error(err)
	}

	// encode the bytes into a png image
	if err := png.Encode(buf, img); err != nil {
		return common.NewFr24Error(err)
	}

	return nil
}

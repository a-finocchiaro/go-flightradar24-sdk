package client

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/a-finocchiaro/adsb_flightradar_top10/fr24"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/flights"
)

func GetFlightDetails(requester fr24.Requester, flight_id string) (flights.Flight, error) {
	var flight flights.Flight

	endpoint := fmt.Sprintf("%s?airport=%s", fr24.FR24_ENDPOINTS["flight_details"], flight_id)

	if err := fr24.SendRequest(requester, endpoint, &flight); err != nil {
		return flight, fr24.NewFr24Error(err)
	}

	return flight, nil
}

func GetFlights(requester common.Requester, flightFeed flights.Fr24FeedInterface) error {
	body, err := requester(common.FR24_ENDPOINTS["all_tracked"])

	if err != nil {
		return common.NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &flightFeed); err != nil {
		return common.NewFr24Error(err)
	}

	return nil
}

func GetRandomFlight(requester common.Requester) (string, error) {
	var rand_flight flights.FeedFlightData
	var flightId string
	var feedData flights.Fr24FeedData

	err := GetFlights(requester, &feedData)

	if err != nil {
		return "", err
	}

	// find a random flight
	rand_idx := rand.Intn(len(feedData.Flights))
	idx := 0

	for flight_id, flight := range feedData.Flights {
		if idx == rand_idx {
			rand_flight = flight
			flightId = flight_id
			break
		}

		idx++
	}

	// provide a link to the flight
	return fmt.Sprintf("%s/%s/%s\n", common.FR24_BASE, rand_flight.Callsign, flightId), nil
}

func GetFR24MostTracked(requester common.Requester) (flights.Fr24MostTrackedRes, error) {
	var most_tracked flights.Fr24MostTrackedRes
	body, err := requester(common.FR24_ENDPOINTS["most_tracked"])

	if err != nil {
		log.Fatalln(err)
		return most_tracked, err
	}

	if err := json.Unmarshal(body, &most_tracked); err != nil {
		return most_tracked, common.NewFr24Error(err)
	}

	return most_tracked, nil
}

package client

import (
	"fmt"
	"math/rand"

	"github.com/a-finocchiaro/adsb_flightradar_top10/internal"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/flights"
)

func GetFlightDetails(requester common.Requester, flight_id string) (flights.Flight, error) {
	var flight flights.Flight

	endpoint := fmt.Sprintf("%s?airport=%s", common.FR24_ENDPOINTS["flight_details"], flight_id)

	if err := internal.SendRequest(requester, endpoint, &flight); err != nil {
		return flight, common.NewFr24Error(err)
	}

	return flight, nil
}

// Gets the latest copy of the flight feed from the all tracked endpoint and parses the data
// into a Fr24FeedData object.
func GetFlights(requester common.Requester) (flights.Fr24FeedData, error) {
	var flightFeed flights.Fr24FeedData

	if err := internal.SendRequest(requester, common.FR24_ENDPOINTS["all_tracked"], &flightFeed); err != nil {
		return flightFeed, common.NewFr24Error(err)
	}

	return flightFeed, nil
}

func GetRandomFlight(requester common.Requester) (string, error) {
	var rand_flight flights.FeedFlightData
	var flightId string

	feedData, err := GetFlights(requester)

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

	if err := internal.SendRequest(requester, common.FR24_ENDPOINTS["most_tracked"], &most_tracked); err != nil {
		return most_tracked, common.NewFr24Error(err)
	}

	return most_tracked, nil
}

package fr24

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
)

type FlightData struct {
	Icao_24bit               string
	Lat                      float32
	Long                     float32
	Heading                  int
	Altitude                 int
	Ground_speed             int
	Squawk                   string
	Fnumber                  string
	Aircraft_code            string
	Registration             string
	Time                     int
	Origin_airport_iata      string
	Destination_airport_iata string
	Airline_iata             string
	On_ground                int
	Vertical_speed           int
	Callsign                 string
	SomeNum                  int // figure out what this value is
	Airline_icao             string
}

func (fd *FlightData) UnmarshalJSON(data []byte) error {
	/*
	* Parses the mixed type array flight data from the feed API endpoint.
	 */
	temp := []interface{}{
		&fd.Icao_24bit,
		&fd.Lat,
		&fd.Long,
		&fd.Heading,
		&fd.Altitude,
		&fd.Ground_speed,
		&fd.Squawk,
		&fd.Fnumber,
		&fd.Aircraft_code,
		&fd.Registration,
		&fd.Time,
		&fd.Origin_airport_iata,
		&fd.Destination_airport_iata,
		&fd.Airline_iata,
		&fd.On_ground,
		&fd.Vertical_speed,
		&fd.Callsign,
		&fd.SomeNum,
		&fd.Airline_icao,
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	return nil
}

type Fr24FeedData struct {
	Full_count int                        `json:"full_count"`
	Version    int                        `json:"version"`
	Flights    map[string]json.RawMessage `json:"-"`
}

func (f *Fr24FeedData) UnmarshalJSON(data []byte) error {
	/**
	* Parses flight feed data which is returned in a very strange mixed type format.
	 */
	temp := struct {
		FullCount int                        `json:"full_count"`
		Version   int                        `json:"version"`
		Flights   map[string]json.RawMessage `json:"-"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if err := json.Unmarshal(data, &temp.Flights); err != nil {
		return err
	}

	f.Full_count = temp.FullCount
	f.Flights = temp.Flights

	return nil
}

func GetRandomFlight(requester Requester) {
	var feedData Fr24FeedData
	var rand_flight json.RawMessage
	var flightId string

	body, err := requester(FR24_ENDPOINTS["all_tracked"])

	if err != nil {
		log.Fatalln(err)
	}

	if err := json.Unmarshal(body, &feedData); err != nil {
		log.Fatalln(err)
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

	// unmarshal the flight data into a struct
	var flightData FlightData

	if err := json.Unmarshal(rand_flight, &flightData); err != nil {
		log.Fatalln(err)
	}

	// provide a link to the flight
	fmt.Println(flightData)
	fmt.Printf("%s/%s/%s\n", FR24_BASE, flightData.Callsign, flightId)
}

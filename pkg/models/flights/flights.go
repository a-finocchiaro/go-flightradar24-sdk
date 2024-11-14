package flights

import (
	"encoding/json"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
)

type Flight struct {
	Identification FlightIdentification `json:"identification"`
	Status         FlightStatus         `json:"status"`
	Level          string               `json:"level"`
	Promote        bool                 `json:"promote"`
	Aircraft       FlightAircraft       `json:"aircraft"`
	Airline        Airline              `json:"airline"`
	Owner          any                  `json:"owner"`
	Airspace       any                  `json:"airspace"`
	Airport        FlightAirportPair    `json:"airport"`
	FlightHistory  FlightHistory        `json:"flightHistory"`
	Ems            any                  `json:"ems"`
	Availability   []string             `json:"availability"`
	Time           FlightTime           `json:"time"`
	Trail          []BreadcrumbStats    `json:"trail"`
	FirstTimestamp int                  `json:"firstTimestamp"`
	S              string               `json:"s"`
}

type FlightIdentificationBase struct {
	ID     string       `json:"id"`
	Number FlightNumber `json:"number"`
}

type FlightIdentification struct {
	FlightIdentificationBase
	Row      int64  `json:"row"`
	Callsign string `json:"callsign"`
}

type FlightNumber struct {
	Default     string `json:"default"`
	Alternative string `json:"alternative"`
}

type FlightStatus struct {
	Live      bool    `json:"live"`
	Text      string  `json:"text"`
	Icon      string  `json:"icon"`
	Estimated any     `json:"estimated"`
	Ambiguous bool    `json:"ambiguous"`
	Generic   Generic `json:"generic"`
}

type Generic struct {
	Status    FlightStatusGenericStatus `json:"status"`
	EventTime FlightStatusEventTime     `json:"eventTime"`
}

type FlightStatusGenericStatus struct {
	Text  string `json:"text"`
	Color string `json:"color"`
	Type  string `json:"type"`
}

type FlightStatusEventTime struct {
	Utc   int `json:"utc"`
	Local int `json:"local"`
}

type FlightAircraft struct {
	Model        FlightAircraftModel    `json:"model"`
	CountryID    int                    `json:"countryId"`
	Registration string                 `json:"registration"`
	Age          any                    `json:"age"`
	Msn          any                    `json:"msn"`
	Images       common.MultiSizeImages `json:"images"`
	Hex          string                 `json:"hex"`
}

type FlightAircraftModel struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type Airline struct {
	Name  string              `json:"name"`
	Short string              `json:"short"`
	Code  common.IataIcaoCode `json:"code"`
	URL   string              `json:"url"`
}

type PositionAlt struct {
	common.Position
	Altitude int `json:"altitude"`
}

type FlightAirportPair struct {
	Origin      FlightAirportCurrent `json:"origin"`
	Destination FlightAirportCurrent `json:"destination"`
	Real        any                  `json:"real"`
}

type FlightAirport struct {
	Name     string                         `json:"name"`
	Code     common.IataIcaoCode            `json:"code"`
	Position PositionAlt                    `json:"position"`
	Timezone common.TimezoneWithOffsetHours `json:"timezone"`
	Visible  bool                           `json:"visible"`
	Website  any                            `json:"website"`
}

type FlightAirportCurrent struct {
	FlightAirport
	Info FlightAirportTerminalBaggageGate `json:"info"`
}

type FlightAirportTerminalBaggageGate struct {
	Terminal string `json:"terminal"`
	Baggage  string `json:"baggage"`
	Gate     string `json:"gate"`
}

type FlightHistory struct {
	Aircraft []FlightHistoryAircraft `json:"aircraft"`
}

type FlightHistoryAircraft struct {
	Identification FlightIdentificationBase `json:"identification"`
	Airport        FlightAirport            `json:"airport"`
	Time           RealTime                 `json:"time"`
}

type RealTime struct {
	Real RealDepartureTime `json:"real"`
}

type RealDepartureTime struct {
	Departure int `json:"departure"`
}

type FlightTime struct {
	Scheduled  FlightDepartureArrivalTime   `json:"scheduled"`
	Real       FlightDepartureArrivalTime   `json:"real"`
	Estimated  FlightDepartureArrivalTime   `json:"estimated"`
	Other      FlightEta                    `json:"other"`
	Historical FlightHistoricalTimeAndDelay `json:"historical"`
}

type FlightDepartureArrivalTime struct {
	Departure int `json:"departure"`
	Arrival   int `json:"arrival"`
}

type FlightEta struct {
	Eta     int `json:"eta"`
	Updated int `json:"updated"`
}

type FlightHistoricalTimeAndDelay struct {
	Flighttime string `json:"flighttime"`
	Delay      string `json:"delay"`
}

type BreadcrumbStats struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Alt int     `json:"alt"`
	Spd int     `json:"spd"`
	Ts  int     `json:"ts"`
	Hd  int     `json:"hd"`
}

type FeedFlightData struct {
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

func (fd *FeedFlightData) UnmarshalJSON(data []byte) error {
	/*
	* Parses the mixed type array flight data from the feed API endpoint.
	 */

	// flight data will always have a start byte of 91 since that is the ASCII value of
	// '[', which is the start of an array. We can safely ignore any non-arrays here, but
	// without an error since we just want to ignore this.
	if data[0] != 91 {
		return nil
	}

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
	Full_count int                       `json:"full_count"`
	Version    int                       `json:"version"`
	Flights    map[string]FeedFlightData `json:"-"`
}

func (f *Fr24FeedData) UnmarshalJSON(data []byte) error {
	/**
	* Parses flight feed data which is returned in a very strange mixed type format.
	 */
	temp := struct {
		FullCount int                        `json:"full_count"`
		Version   int                        `json:"version"`
		Flights   map[string]json.RawMessage `json:"-"`
	}{
		Flights: make(map[string]json.RawMessage),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return common.NewFr24Error(err)
	}

	if err := json.Unmarshal(data, &temp.Flights); err != nil {
		return common.NewFr24Error(err)
	}

	// remove the full_count and version keys since they should not exist in the flight data
	// this is jank, but seems to be the best way to solve this issue.
	delete(temp.Flights, "full_count")
	delete(temp.Flights, "version")

	f.Full_count = temp.FullCount
	f.Flights = make(map[string]FeedFlightData)

	// parse the json of each flight
	for flightId, flight := range temp.Flights {
		var flightData FeedFlightData

		if err := json.Unmarshal(flight, &flightData); err != nil {
			continue
		}

		f.Flights[flightId] = flightData
	}

	return nil
}

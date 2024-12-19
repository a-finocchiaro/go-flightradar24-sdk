package airports

import (
	"encoding/json"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
)

type AirportRouteResponse struct {
	Arrivals   AirportRouteCountry `json:"arrivals"`
	Departures AirportRouteCountry `json:"departures"`
}

type AirportRouteCountry struct {
	Country AirportRoute `json:"-"`
}

func (a *AirportRouteCountry) UnmarshalJSON(data []byte) error {
	// parse into temp object
	temp := struct {
		Country map[string]json.RawMessage `json:"-"`
	}{
		Country: make(map[string]json.RawMessage),
	}

	if err := json.Unmarshal(data, &temp.Country); err != nil {
		return common.NewFr24Error(err)
	}

	// Get the country name from the *first* key. It appears that there
	// is never a time that this will include more than a single country

	// This also will always correspond to the ARRIVING country. Ex:
	// if looking at routes from DFW to LHR, this will be "United Kingdom"
	for country := range temp.Country {
		a.Country.Name = country

		var d AirportRoute
		if err := json.Unmarshal(temp.Country[country], &d); err != nil {
			return common.NewFr24Error(err)
		}

		a.Country.Number = d.Number
		a.Country.Airports = d.Airports

		break
	}

	return nil
}

type AirportRoute struct {
	Name     string
	Number   AirportRouteNumber `json:"number"`
	Airports AirportRouteData   `json:"airports"`
}

type AirportRouteNumber struct {
	Airports int `json:"airports"`
	Flights  int `json:"flights"`
}

type AirportRouteData struct {
	Iata string
	AirportRouteAirport
}

func (ar *AirportRouteData) UnmarshalJSON(data []byte) error {
	temp := struct {
		Airports map[string]json.RawMessage `json:"-"`
	}{
		Airports: make(map[string]json.RawMessage),
	}

	if err := json.Unmarshal(data, &temp.Airports); err != nil {
		return common.NewFr24Error(err)
	}

	var airportIata string

	// Get the first value out of the airports
	for airport := range temp.Airports {
		airportIata = airport
		break
	}

	// save the airport IATA code
	ar.Iata = airportIata

	var a AirportRouteAirport

	if err := json.Unmarshal(temp.Airports[airportIata], &a); err != nil {
		return common.NewFr24Error(err)
	}

	// set the remaining values
	ar.Name = a.Name
	ar.City = a.City
	ar.Icao = a.Icao
	ar.Flights = a.Flights

	return nil
}

type AirportRouteAirport struct {
	Name    string                   `json:"name"`
	City    string                   `json:"city"`
	Icao    string                   `json:"icao"`
	Flights []AirportRouteFlightData `json:"flights"`
}

func (ar *AirportRouteAirport) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name    string `json:"name"`
		City    string `json:"city"`
		Icao    string `json:"icao"`
		Flights map[string]struct {
			Airline AirportRouteFlightInfoExtended    `json:"Airline"`
			Utc     map[string]AirportRouteFlightTime `json:"utc"`
		} `json:"flights"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return common.NewFr24Error(err)
	}

	for id, flightData := range temp.Flights {
		flight := AirportRouteFlightData{
			ID:      id,
			Airline: flightData.Airline,
		}

		for date, timeData := range flightData.Utc {
			timeData.Date = date
			flight.Times = append(flight.Times, timeData)
		}

		ar.Flights = append(ar.Flights, flight)
	}

	ar.City = temp.City
	ar.Icao = temp.Icao
	ar.Name = temp.Name

	return nil
}

type AirportRouteFlightData struct {
	ID      string                         `json:"id"`
	Airline AirportRouteFlightInfoExtended `json:"Airline"`
	Times   []AirportRouteFlightTime       `json:"times"`
}

type AirportRouteFlightInfoExtended struct {
	Airline AirportRouteAirline    `json:"airline"`
	Utc     AirportRouteFlightTime `json:"utc"`
}

type AirportRouteAirline struct {
	Name string `json:"name"`
	Iata string `json:"iata"`
	Icao string `json:"icao"`
	Url  string `json:"url"`
}

type AirportRouteFlightTime struct {
	Date string `json:"id"`
	AirportRouteFlightTimeAircraftInfo
}

type AirportRouteFlightTimeAircraftInfo struct {
	Aircraft  string `json:"aircraft"`
	Time      string `json:"time"`
	Timestamp int64  `json:"timestamp"`
	Offset    int    `json:"offset"`
}

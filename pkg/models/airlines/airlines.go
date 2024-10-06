package airlines

type AirlineRes struct {
	Version int           `json:"version"`
	Rows    []AirlineData `json:"rows"`
}

type AirlineData struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
	Icao string `json:"ICAO"`
}

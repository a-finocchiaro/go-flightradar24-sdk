package common

// Generic types that are used in multiple places.

type (
	Requester func(string) ([]byte, error)
)

type IataIcaoCode struct {
	Iata string `json:"iata"`
	Icao string `json:"icao"`
}

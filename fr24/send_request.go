package fr24

import (
	"encoding/json"
)

// Forwards a request to FlightRadar24 and unmarshals the response into the desired
// object.
func SendRequest[T any](requester Requester, endpoint string, obj *T) error {
	body, err := requester(endpoint)

	if err != nil {
		return NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		return NewFr24Error(err)
	}

	return nil
}

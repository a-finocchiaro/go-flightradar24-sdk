package internal

import (
	"encoding/json"

	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
)

// Forwards a request to FlightRadar24 and unmarshals the response into the desired
// object.
func SendRequest[T any](requester common.Requester, endpoint string, obj *T) error {
	body, err := requester(endpoint)

	if err != nil {
		return common.NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		return common.NewFr24Error(err)
	}

	return nil
}

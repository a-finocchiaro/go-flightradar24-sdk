package client

import (
	"fmt"

	"github.com/a-finocchiaro/go-flightradar24-sdk/internal"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/search"
)

// Searches Flightradar for a resource (live flight, scheduled flight, airport,
// aircraft, or operator)
func Search(requester common.Requester, query string) (search.SearchResultResponse, error) {
	var searchRes search.SearchResultResponse

	endpoint := fmt.Sprintf("%s?query=%s&limit=25", common.FR24_ENDPOINTS["search"], query)

	if err := internal.SendRequest(requester, endpoint, &searchRes); err != nil {
		return searchRes, common.NewFr24Error(err)
	}

	return searchRes, nil
}

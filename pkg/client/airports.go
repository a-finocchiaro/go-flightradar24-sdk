package client

import (
	"fmt"
	"slices"

	"github.com/a-finocchiaro/go-flightradar24-sdk/internal"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/airports"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
)

func getSupportedPlugins() []string {
	return []string{
		"details",
		"flightdiary",
		"schedule",
		"weather",
		"runways",
		"satelliteImage",
		"scheduledRoutesStatistics",
	}
}

// Gets detailed information from the Flightradar24 API about an airport and parses the retured
// data into structured data.
func GetAirportDetails(requester common.Requester, code string, plugins []string) (airports.AirportPluginData, error) {
	var airport airports.AirportApiResponse
	var pluginQuery string

	// verify the plugins are accepted
	for _, plugin := range plugins {
		if !slices.Contains(getSupportedPlugins(), plugin) {
			err := common.Fr24Error{Err: fmt.Sprintf("Plugin %s not supported.", plugin)}

			return airport.Result.Response.Airport.PluginData, err
		}

		// add plugin to the plugin query
		pluginQuery += fmt.Sprintf("&plugin[]=%s", plugin)
	}

	endpoint := fmt.Sprintf("%s?code=%s&limit=50%s", common.FR24_ENDPOINTS["airport_detail"], code, plugins)

	if err := internal.SendRequest(requester, endpoint, &airport); err != nil {
		return airport.Result.Response.Airport.PluginData, common.NewFr24Error(err)
	}

	return airport.Result.Response.Airport.PluginData, nil
}

// Gets brief airport information from the /airports/traffic-stats endpoint from the FR24
// base URL.
// Acceps an airport IATA or ICAO code as an argument for code.
func GetAirportBrief(requester common.Requester, code string) (airports.AirportBriefDetails, error) {
	var airport airports.AirportBriefResponse

	endpoint := fmt.Sprintf("%s?airport=%s", common.FR24_ENDPOINTS["airport_brief"], code)

	if err := internal.SendRequest(requester, endpoint, &airport); err != nil {
		return airport.Details, common.NewFr24Error(err)
	}

	return airport.Details, nil
}

// Gets a report of the current airport distruption rankings
func GetAirportDisruptions(requester common.Requester) ([]airports.AirportDisruptionRank, error) {
	var disruptions airports.AirportDistruptionApiResponse

	if err := internal.SendRequest(requester, common.FR24_ENDPOINTS["airport_disruptions"], &disruptions); err != nil {
		return disruptions.Data.Rank, common.NewFr24Error(err)
	}

	return disruptions.Data.Rank, nil
}

// Gets routes between specific airports
func GetAirportRoutes(requester common.Requester, airport1 string, airport2 string) (airports.AirportRouteResponse, error) {
	var routeInfo airports.AirportRouteResponse

	url := fmt.Sprintf("%s/data/airports/%s/routes?get-airport-arr=%s&format=json", common.FR24_BASE, airport1, airport2)

	if err := internal.SendRequest(requester, url, &routeInfo); err != nil {
		return routeInfo, common.NewFr24Error(err)
	}

	return routeInfo, nil
}

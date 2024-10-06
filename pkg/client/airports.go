package client

import (
	"encoding/json"
	"fmt"
	"slices"

	"github.com/a-finocchiaro/adsb_flightradar_top10/fr24"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/airports"
	"github.com/a-finocchiaro/adsb_flightradar_top10/pkg/models/common"
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
func GetAirportDetails(requester fr24.Requester, code string, plugins []string) (airports.AirportPluginData, error) {
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

	endpoint := fmt.Sprintf("%s?code=%s&limit=100%s", common.FR24_ENDPOINTS["airport_detail"], code, plugins)

	body, err := requester(endpoint)

	if err != nil {
		return airport.Result.Response.Airport.PluginData, common.NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &airport); err != nil {
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

	body, err := requester(endpoint)

	if err != nil {
		return airport.Details, common.NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &airport); err != nil {
		return airport.Details, common.NewFr24Error(err)
	}

	return airport.Details, nil
}

// Gets a report of the current airport distruption rankings
func GetAirportDisruptions(requester common.Requester) ([]airports.AirportDisruptionRank, error) {
	var disruptions airports.AirportDistruptionApiResponse
	body, err := requester(common.FR24_ENDPOINTS["airport_disruptions"])

	if err != nil {
		return disruptions.Data.Rank, common.NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &disruptions); err != nil {
		return disruptions.Data.Rank, common.NewFr24Error(err)
	}

	return disruptions.Data.Rank, nil
}

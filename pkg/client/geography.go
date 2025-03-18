package client

import (
	"fmt"
	"math"

	"github.com/a-finocchiaro/go-flightradar24-sdk/internal"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/geography"
)

func GetBounds(zone map[string]float64) string {
	return fmt.Sprintf("%.6f, %.6f, %.6f, %.6f", zone["tl_y"], zone["br_y"], zone["tl_x"], zone["br_x"])
}

func GetBoundsByPoint(latitude, longitude, radius float64) string {
	// Convert radius to kilometers
	halfSideInKm := math.Abs(radius) / 1000

	// Convert latitude and longitude to radians
	lat := latitude * (math.Pi / 180)
	lon := longitude * (math.Pi / 180)

	approxEarthRadius := 6371.0
	hypotenuseDistance := math.Sqrt(2 * math.Pow(halfSideInKm, 2))

	// Calculate minimum latitude and longitude (bottom-left corner)
	latMin := math.Asin(math.Sin(lat)*math.Cos(hypotenuseDistance/approxEarthRadius) +
		math.Cos(lat)*math.Sin(hypotenuseDistance/approxEarthRadius)*math.Cos(225*(math.Pi/180)))
	lonMin := lon + math.Atan2(
		math.Sin(225*(math.Pi/180))*math.Sin(hypotenuseDistance/approxEarthRadius)*math.Cos(lat),
		math.Cos(hypotenuseDistance/approxEarthRadius)-math.Sin(lat)*math.Sin(latMin),
	)

	// Calculate maximum latitude and longitude (top-right corner)
	latMax := math.Asin(math.Sin(lat)*math.Cos(hypotenuseDistance/approxEarthRadius) +
		math.Cos(lat)*math.Sin(hypotenuseDistance/approxEarthRadius)*math.Cos(45*(math.Pi/180)))
	lonMax := lon + math.Atan2(
		math.Sin(45*(math.Pi/180))*math.Sin(hypotenuseDistance/approxEarthRadius)*math.Cos(lat),
		math.Cos(hypotenuseDistance/approxEarthRadius)-math.Sin(lat)*math.Sin(latMax),
	)

	// Convert radians back to degrees
	rad2deg := 180 / math.Pi

	zone := map[string]float64{
		"tl_y": latMax * rad2deg,
		"br_y": latMin * rad2deg,
		"tl_x": lonMin * rad2deg,
		"br_x": lonMax * rad2deg,
	}

	return GetBounds(zone)
}

func GetZones(requester common.Requester) (geography.Fr24Zones, error) {
	var zones geography.Fr24Zones

	if err := internal.SendRequest(requester, common.FR24_ENDPOINTS["zones"], &zones); err != nil {
		return zones, common.NewFr24Error(err)
	}

	return zones, nil
}

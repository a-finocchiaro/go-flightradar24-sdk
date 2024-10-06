package common

import "fmt"

const FR24_API = "https://api.flightradar24.com"
const FR24_BASE = "https://www.flightradar24.com"
const FR24_DATA_CLOUD = "https://data-cloud.flightradar24.com"
const FR24_CDN = "https://cdn.flightradar24.com"
const FR24_DATA_LIVE = "https://data-live.flightradar24.com"

var FR24_ENDPOINTS = map[string]string{
	"airlines":            fmt.Sprintf("%s/_json/airlines.php", FR24_BASE),
	"airline_logo_cdn":    fmt.Sprintf("%s/assets/airlines/logotypes", FR24_CDN),
	"airline_logo":        fmt.Sprintf("%s/static/images/data/operators", FR24_BASE),
	"airport_brief":       fmt.Sprintf("%s/airports/traffic-stats", FR24_BASE),
	"airport_detail":      fmt.Sprintf("%s/common/v1/airport.json", FR24_API),
	"airport_disruptions": fmt.Sprintf("%s/webapi/v1/airport-disruptions", FR24_BASE),
	"all_tracked":         fmt.Sprintf("%s/zones/fcgi/feed.js", FR24_DATA_CLOUD),
	"flight_details":      fmt.Sprintf("%s/clickhandler", FR24_DATA_LIVE),
	"most_tracked":        fmt.Sprintf("%s/flights/most-tracked", FR24_BASE),
}

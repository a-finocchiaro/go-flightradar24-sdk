package fr24

import "fmt"

const FR24_BASE = "https://www.flightradar24.com"
const FR24_DATA_CLOUD = "https://data-cloud.flightradar24.com"

var FR24_ENDPOINTS = map[string]string{
	"most_tracked": fmt.Sprintf("%s/flights/most-tracked", FR24_BASE),
	"all_tracked":  fmt.Sprintf("%s/zones/fcgi/feed.js", FR24_DATA_CLOUD),
}

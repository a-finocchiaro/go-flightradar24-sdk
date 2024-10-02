package fr24

import (
	"encoding/json"
	"fmt"
	"slices"
)

type AirportApiResponse struct {
	Result AirportResult `json:"result"`
}

type AirportResult struct {
	Request  AirportRequest         `json:"request"`
	Response AirportRequestResponse `json:"response"`
}

type AirportRequestResponse struct {
	Airport AirportRequestResponseAirport `json:"airport"`
}

type AirportRequestResponseAirport struct {
	PluginData AirportPluginData `json:"pluginData"`
}

type AirportRequest struct {
	Callback      string                      `json:"callback"`
	Code          string                      `json:"code"`
	Device        string                      `json:"device"`
	Fleet         string                      `json:"fleet"`
	Format        string                      `json:"format"`
	Limit         int                         `json:"limit"`
	Page          int                         `json:"page"`
	Pk            string                      `json:"pk"`
	Plugin        []string                    `json:"plugin"`
	Token         string                      `json:"token"`
	PluginSetting AirportRequestPluginSetting `json:"plugin-setting"`
}

type AirportRequestPluginSetting struct {
	Schedule       AirportRequestPluginSettingSchedule       `json:"schedule"`
	SatelliteImage AirportRequestPluginSettingSatelliteImage `json:"satelliteImage"`
}

type AirportRequestPluginSettingSchedule struct {
	Mode      any `json:"mode"`
	Timestamp int `json:"timestamp"`
}

type AirportRequestPluginSettingSatelliteImage struct {
	Scale int `json:"scale"`
}

type AirportPluginData struct {
	Details                   AirportDetails                   `json:"details"`
	Schedule                  AirportSchedule                  `json:"schedule"`
	Weather                   AirportWeather                   `json:"weather"`
	AircraftCount             AirportAircraftCount             `json:"AircraftCount"`
	Runways                   AirportRunways                   `json:"AirportRunways"`
	ScheduledFlightStatistics AirportScheduledFlightStatistics `json:"scheduledFlightStatistics"`
	SatelliteImage            string                           `json:"satelliteImage"`
	SatelliteImageProperties  AirportSatelliteImageProperties  `json:"SatelliteImageProperties"`
}

// Details
type AirportDetails struct {
	Name          string                   `json:"name"`
	Code          AirportDetailsCode       `json:"code"`
	DelayIndex    AirportDetailsDelayIndex `json:"delayIndex"`
	Stats         AirportStats             `json:"stats"`
	Position      AirportDetailsPosition   `json:"position"`
	Timezone      Timezone                 `json:"timezone"`
	URL           AirportDetailsURL        `json:"url"`
	AirportImages MultiSizeImages          `json:"airportImages"`
	Visible       bool                     `json:"visible"`
}

type AirportDetailsCode struct {
	Iata string `json:"iata"`
	Icao string `json:"icao"`
}

type AirportStats struct {
	Arrivals   ArrivalDepartureStats `json:"arrivals"`
	Departures ArrivalDepartureStats `json:"departures"`
}

type ArrivalDepartureStats struct {
	DelayIndex float64                `json:"delayIndex"`
	DelayAvg   int                    `json:"delayAvg"`
	Percentage AirportStatsPercentage `json:"percentage"`
	Recent     AirportRecentStats     `json:"recent"`
	Today      AirportDailyStats      `json:"today"`
	Yesterday  AirportDailyStats      `json:"yesterday"`
	Tomorrow   AirportDailyStats      `json:"tomorrow"`
}

type AirportRecentStats struct {
	DelayIndex float64                `json:"delayIndex"`
	DelayAvg   int                    `json:"delayAvg"`
	Percentage AirportStatsPercentage `json:"percentage"`
	Quantity   AirportStatsQuantity   `json:"quantity"`
}

type AirportDailyStats struct {
	Percentage AirportStatsPercentage `json:"percentage"`
	Quantity   AirportStatsQuantity   `json:"quantity"`
}

type AirportStatsPercentage struct {
	Delayed  float32 `json:"dealyed"`
	Canceled float32 `json:"canceled"`
	Trend    string  `json:"trend"`
}

type AirportStatsQuantity struct {
	OnTime   int `json:"onTime"`
	Delayed  int `json:"delayed"`
	Canceled int `json:"canceled"`
}

type AirportDetailsDelayIndex struct {
	Arrivals   any `json:"arrivals"`
	Departures any `json:"departures"`
}

type AirportDetailsCountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
	ID   int    `json:"id"`
}

type AirportDetailsRegion struct {
	City string `json:"city"`
}

type AirportDetailsPosition struct {
	Latitude  float64               `json:"latitude"`
	Longitude float64               `json:"longitude"`
	Elevation int                   `json:"elevation"`
	Country   AirportDetailsCountry `json:"country"`
	Region    AirportDetailsRegion  `json:"region"`
}

type AirportDetailsURL struct {
	Homepage  any    `json:"homepage"`
	Webcam    any    `json:"webcam"`
	Wikipedia string `json:"wikipedia"`
}

type ImageData struct {
	Src       string `json:"src"`
	Link      string `json:"link"`
	Copyright string `json:"copyright"`
	Source    string `json:"source"`
}

type MultiSizeImages struct {
	Thumbnails []ImageData `json:"thumbnails"`
	Medium     []ImageData `json:"medium"`
	Large      []ImageData `json:"large"`
}

// Schedule
type AirportSchedule struct {
	Arrivals   AirportScheduleData `json:"arrivals"`
	Departures AirportScheduleData `json:"departures"`
}

type AirportScheduleData struct {
	Item AirportScheduleItemCounts `json:"item"`
	Page struct {
		Current int `json:"current"`
		Total   int `json:"Total"`
	}
	Timestamp int `json:"timetamp"`
	Data      []FlightArrivalDepartureData
}

type AirportScheduleItemCounts struct {
	Current int `json:"current"`
	Total   int `json:"total"`
	Limit   int `json:"limit"`
}

type FlightArrivalDepartureData struct {
	Flight Flight `json:"flight"`
}
type FlightIdentification struct {
	ID        string                     `json:"id"`
	Row       int64                      `json:"row"`
	Number    FlightIdentificationNumber `json:"number"`
	Callsign  string                     `json:"callsign"`
	Codeshare any                        `json:"codeshare"`
}
type FlightIdentificationNumber struct {
	Default     string `json:"default"`
	Alternative string `json:"alternative"`
}

type FlightStatus struct {
	Live      bool                `json:"live"`
	Text      string              `json:"text"`
	Icon      string              `json:"icon"`
	Estimated any                 `json:"estimated"`
	Ambiguous bool                `json:"ambiguous"`
	Generic   FlightStatusGeneric `json:"generic"`
}

type FlightStatusGenericStatus struct {
	Text     string `json:"text"`
	Type     string `json:"type"`
	Color    string `json:"color"`
	Diverted any    `json:"diverted"`
}

type FlightStatusGenericEventTime struct {
	Utc   int `json:"utc"`
	Local int `json:"local"`
}

type FlightStatusGeneric struct {
	Status    FlightStatusGenericStatus    `json:"status"`
	EventTime FlightStatusGenericEventTime `json:"eventTime"`
}

type FlightAircraft struct {
	Model          FlightAircraftModel        `json:"model"`
	Registration   string                     `json:"registration"`
	Country        Country                    `json:"country"`
	Hex            string                     `json:"hex"`
	Restricted     bool                       `json:"restricted"`
	SerialNo       string                     `json:"serialNo"`
	Age            FlightAircraftAge          `json:"age"`
	Availability   FlightAircraftAvailability `json:"availability"`
	OnGroundUpdate int                        `json:"onGroundUpdate"`
	HoursDiff      float32                    `json:"hoursDiff"`
	TimeDiff       float32                    `json:"timeDiff"`
}

type FlightAircraftModel struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type FlightAircraftCountry struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
}

type FlightAircraftAge struct {
	Availability bool `json:"availability"`
}

type FlightAircraftAvailability struct {
	SerialNo bool `json:"serialNo"`
	Age      bool `json:"age"`
}

type FlightAirline struct {
	Name  string       `json:"name"`
	Code  IataIcaoCode `json:"code"`
	Short string       `json:"short"`
}

type IataIcaoCode struct {
	Iata string `json:"iata"`
	Icao string `json:"icao"`
}

type FlightOwner struct {
	Name string       `json:"name"`
	Code IataIcaoCode `json:"code"`
	Logo string       `json:"logo"`
}

type FlightAirport struct {
	Origin      FlightAiportData `json:"origin"`
	Destination FlightAiportData `json:"destination"`
	Real        any              `json:"real"`
}

type FlightAiportData struct {
	Code     IataIcaoCode      `json:"code"`
	Timezone Timezone          `json:"timezone"`
	Info     FlightAirportInfo `json:"info"`
	Name     string            `json:"name"`
	Position Position          `json:"position"`
	Visible  bool              `json:"visible"`
}

type Timezone struct {
	Name     string `json:"name"`
	Offset   int    `json:"offset"`
	Abbr     string `json:"abbr"`
	AbbrName string `json:"abbrName"`
	IsDst    bool   `json:"isDst"`
}

type FlightAirportInfo struct {
	Terminal any `json:"terminal"`
	Baggage  any `json:"baggage"`
	Gate     any `json:"gate"`
}
type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
	ID   int    `json:"id"`
}
type FlightAirportRegion struct {
	City string `json:"city"`
}

type Position struct {
	Latitude  float64             `json:"latitude"`
	Longitude float64             `json:"longitude"`
	Country   Country             `json:"country"`
	Region    FlightAirportRegion `json:"region"`
}

type Scheduled struct {
	Departure int `json:"departure"`
	Arrival   int `json:"arrival"`
}
type Real struct {
	Departure int `json:"departure"`
	Arrival   any `json:"arrival"`
}
type Estimated struct {
	Departure any `json:"departure"`
	Arrival   int `json:"arrival"`
}
type Other struct {
	Eta      int `json:"eta"`
	Duration any `json:"duration"`
}
type FlightTime struct {
	Scheduled Scheduled `json:"scheduled"`
	Real      Real      `json:"real"`
	Estimated Estimated `json:"estimated"`
	Other     Other     `json:"other"`
}
type Flight struct {
	Identification FlightIdentification `json:"identification"`
	Status         FlightStatus         `json:"status"`
	Aircraft       FlightAircraft       `json:"aircraft"`
	Owner          FlightOwner          `json:"owner"`
	Airline        FlightAirline        `json:"airline"`
	Airport        FlightAirport        `json:"airport"`
	Time           FlightTime           `json:"time"`
}

// Weather
type AirportWeather struct {
	Metar     string               `json:"metar"`
	Time      int                  `json:"time"`
	Qnh       int                  `json:"qnh"`
	Dewpoint  AirportDewpoint      `json:"dewpoint"`
	Humidity  int                  `json:"humidity"`
	Pressure  AirportPressure      `json:"pressure"`
	Sky       AirportSky           `json:"sky"`
	Flight    AirportFlight        `json:"flight"`
	Wind      AirportWind          `json:"wind"`
	Temp      AirportTemp          `json:"temp"`
	Elevation FeetMeterMeasurement `json:"elevation"`
	Cached    int                  `json:"cached"`
}
type AirportDewpoint struct {
	Celsius    int `json:"celsius"`
	Fahrenheit int `json:"fahrenheit"`
}
type AirportPressure struct {
	Hg  int `json:"hg"`
	Hpa int `json:"hpa"`
}
type AirportCondition struct {
	Text string `json:"text"`
}
type AirportVisibility struct {
	Km  int `json:"km"`
	Mi  int `json:"mi"`
	Nmi int `json:"nmi"`
}
type AirportSky struct {
	Condition  AirportCondition  `json:"condition"`
	Visibility AirportVisibility `json:"visibility"`
}
type AirportFlight struct {
	Category string `json:"category"`
}
type AirportDirection struct {
	Degree int    `json:"degree"`
	Text   string `json:"text"`
}
type AirportSpeed struct {
	Kmh  int    `json:"kmh"`
	Kts  int    `json:"kts"`
	Mph  int    `json:"mph"`
	Text string `json:"text"`
}
type AirportWind struct {
	Direction AirportDirection `json:"direction"`
	Speed     AirportSpeed     `json:"speed"`
}
type AirportTemp struct {
	Celsius    int `json:"celsius"`
	Fahrenheit int `json:"fahrenheit"`
}
type FeetMeterMeasurement struct {
	M  int `json:"m"`
	Ft int `json:"ft"`
}

// Aircraft Count
type AirportAircraftCount struct {
	Ground   int                     `json:"ground"`
	OnGround AirportAircraftOnGround `json:"onGround"`
}

type AirportAircraftOnGround struct {
	Visible int `json:"visible"`
	Total   int `json:"total"`
}

// Runways
type AirportRunways struct {
	Runways []Runway `json:"runways"`
}

type Runway struct {
	Name    string               `json:"name"`
	Length  FeetMeterMeasurement `json:"length"`
	Surface RunwaySurface        `json:"surface"`
}

type RunwaySurface struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// satellite imagery
type AirportSatelliteImageProperties struct {
	Center []float32 `json:"center"`
	Zoom   int       `json:"zoom"`
	Scale  int       `json:"scale"`
}

// Scheduled Flight Statistics
type AirportScheduledFlightStatistics struct {
	TotalFlights    int      `json:"totalFlights"`
	TopRoute        TopRoute `json:"topRoute"`
	AirportsServed  int      `json:"airportsServed"`
	CountriesServed int      `json:"countriesServed"`
}
type TopRoute struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Count int    `json:"count"`
}

// Airlines
type AirportAirlines struct {
	Codeshare map[string]AirlineCodeshareData `json:"codeshare"`
}

type AirlineCodeshareData struct {
	Code IataIcaoCode `json:"code"`
}

// AircraftImages
type AircraftImages struct {
	Registration string          `json:"registration"`
	Images       MultiSizeImages `json:"images"`
}

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
func GetAirportDetails(requester Requester, code string, plugins []string) (AirportPluginData, error) {
	var airport AirportApiResponse
	var pluginQuery string

	// verify the plugins are accepted
	for _, plugin := range plugins {
		if !slices.Contains(getSupportedPlugins(), plugin) {
			err := Fr24Error{Err: fmt.Sprintf("Plugin %s not supported.", plugin)}

			return airport.Result.Response.Airport.PluginData, err
		}

		// add plugin to the plugin query
		pluginQuery += fmt.Sprintf("&plugin[]=%s", plugin)
	}

	endpoint := fmt.Sprintf("%s?code=%s&limit=100%s", FR24_ENDPOINTS["airport_detail"], code, plugins)

	body, err := requester(endpoint)

	if err != nil {
		return airport.Result.Response.Airport.PluginData, NewFr24Error(err)
	}

	if err := json.Unmarshal(body, &airport); err != nil {
		return airport.Result.Response.Airport.PluginData, NewFr24Error(err)
	}

	return airport.Result.Response.Airport.PluginData, nil
}

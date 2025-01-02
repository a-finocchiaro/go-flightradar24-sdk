package airports

import (
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/common"
	"github.com/a-finocchiaro/go-flightradar24-sdk/pkg/models/flights"
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
	Code          common.IataIcaoCode      `json:"code"`
	DelayIndex    AirportDetailsDelayIndex `json:"delayIndex"`
	Stats         AirportStats             `json:"stats"`
	Position      AirportDetailsPosition   `json:"position"`
	Timezone      common.Timezone          `json:"timezone"`
	URL           AirportDetailsURL        `json:"url"`
	AirportImages common.MultiSizeImages   `json:"airportImages"`
	Visible       bool                     `json:"visible"`
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
	Flight flights.Flight `json:"flight"`
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

// type FlightStatus struct {
// 	Live      bool                `json:"live"`
// 	Text      string              `json:"text"`
// 	Icon      string              `json:"icon"`
// 	Estimated any                 `json:"estimated"`
// 	Ambiguous bool                `json:"ambiguous"`
// 	Generic   FlightStatusGeneric `json:"generic"`
// }

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

// type FlightAircraft struct {
// 	Model          FlightAircraftModel        `json:"model"`
// 	Registration   string                     `json:"registration"`
// 	Country        common.Country             `json:"country"`
// 	Hex            string                     `json:"hex"`
// 	Restricted     bool                       `json:"restricted"`
// 	SerialNo       string                     `json:"serialNo"`
// 	Age            FlightAircraftAge          `json:"age"`
// 	Availability   FlightAircraftAvailability `json:"availability"`
// 	OnGroundUpdate int                        `json:"onGroundUpdate"`
// 	HoursDiff      float32                    `json:"hoursDiff"`
// 	TimeDiff       float32                    `json:"timeDiff"`
// }

// type FlightAircraftModel struct {
// 	Code string `json:"code"`
// 	Text string `json:"text"`
// }

// type FlightAircraftAge struct {
// 	Availability bool `json:"availability"`
// }

// type FlightAircraftAvailability struct {
// 	SerialNo bool `json:"serialNo"`
// 	Age      bool `json:"age"`
// }

// type FlightAirline struct {
// 	Name  string              `json:"name"`
// 	Code  common.IataIcaoCode `json:"code"`
// 	Short string              `json:"short"`
// }

// type FlightOwner struct {
// 	Name string              `json:"name"`
// 	Code common.IataIcaoCode `json:"code"`
// 	Logo string              `json:"logo"`
// }

// type FlightAirport struct {
// 	Origin      FlightAiportData `json:"origin"`
// 	Destination FlightAiportData `json:"destination"`
// 	Real        any              `json:"real"`
// }

type FlightAiportData struct {
	Code     common.IataIcaoCode `json:"code"`
	Timezone common.Timezone     `json:"timezone"`
	Info     FlightAirportInfo   `json:"info"`
	Name     string              `json:"name"`
	Position common.Position     `json:"position"`
	Visible  bool                `json:"visible"`
}

type FlightAirportInfo struct {
	Terminal any `json:"terminal"`
	Baggage  any `json:"baggage"`
	Gate     any `json:"gate"`
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

//	type FlightTime struct {
//		Scheduled Scheduled `json:"scheduled"`
//		Real      Real      `json:"real"`
//		Estimated Estimated `json:"estimated"`
//		Other     Other     `json:"other"`
//	}
// type Flight struct {
// 	Identification flights.FlightIdentification `json:"identification"`
// 	Status         flights.FlightStatus         `json:"status"`
// 	Aircraft       flights.FlightAircraft       `json:"aircraft"`
// 	Owner          flights.FlightOwner          `json:"owner"`
// 	Airline        flights.Airline              `json:"airline"`
// 	Airport        flights.FlightAirportPair    `json:"airport"`
// 	Time           flights.FlightTime           `json:"time"`
// }

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
	Code common.IataIcaoCode `json:"code"`
}

// AircraftImages
type AircraftImages struct {
	Registration string                 `json:"registration"`
	Images       common.MultiSizeImages `json:"images"`
}

// Airports Brief details, mainly correspond to the data returned by the airports/traffic-stats?airport= endpoint
type AirportBriefResponse struct {
	Details AirportBriefDetails `json:"details"`
}

type AirportBriefDetails struct {
	Name     string                         `json:"name"`
	Code     common.IataIcaoCode            `json:"code"`
	Position AirportPosition                `json:"position"`
	Timezone common.TimezoneWithOffsetHours `json:"timezone"`
	Visible  bool                           `json:"visible"`
	Website  string                         `json:"website"`
	Stats    AirportBriefStats              `json:"stats"`
}

type AirportPosition struct {
	Latitude  float64                `json:"latitude"`
	Longitude float64                `json:"longitude"`
	Altitude  int                    `json:"altitude"`
	Country   common.CountryExtended `json:"country"`
	Region    AirportDetailsRegion   `json:"region"`
}

type AirportBriefStats struct {
	Arrivals   ArrivalDepartureAggregateStats `json:"arrivals"`
	Departures ArrivalDepartureAggregateStats `json:"departures"`
}

type ArrivalDepartureAggregateStats struct {
	DelayIndex int      `json:"delayIndex"`
	DelayAvg   any      `json:"delayAvg"`
	Total      string   `json:"total"`
	Hourly     Hourly   `json:"hourly"`
	Stats      []string `json:"stats"`
}

type Hourly struct {
	Hour0  string `json:"0"`
	Hour1  string `json:"1"`
	Hour2  string `json:"2"`
	Hour3  string `json:"3"`
	Hour4  string `json:"4"`
	Hour5  string `json:"5"`
	Hour6  string `json:"6"`
	Hour7  string `json:"7"`
	Hour8  string `json:"8"`
	Hour9  string `json:"9"`
	Hour10 string `json:"10"`
	Hour11 string `json:"11"`
	Hour12 string `json:"12"`
	Hour13 string `json:"13"`
	Hour14 string `json:"14"`
	Hour15 string `json:"15"`
	Hour16 string `json:"16"`
	Hour17 string `json:"17"`
	Hour18 string `json:"18"`
	Hour19 string `json:"19"`
	Hour20 string `json:"20"`
	Hour21 string `json:"21"`
	Hour22 string `json:"22"`
	Hour23 string `json:"23"`
}

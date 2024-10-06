package common

// Generic types that are used in multiple places.

type (
	Requester func(string) ([]byte, error)
)

type Country struct {
	Name string `json:"name"`
	Code string `json:"code"`
	ID   int    `json:"id"`
}

// Inherits from country, adds CodeLong field
type CountryExtended struct {
	Country
	CodeLong string `json:"codeLong"`
}

type IataIcaoCode struct {
	Iata string `json:"iata"`
	Icao string `json:"icao"`
}

// Sandard image data used for multiple resource types
type ImageData struct {
	Src       string `json:"src"`
	Link      string `json:"link"`
	Copyright string `json:"copyright"`
	Source    string `json:"source"`
}

// Standard FR24 Image sizes
type MultiSizeImages struct {
	Thumbnails []ImageData `json:"thumbnails"`
	Medium     []ImageData `json:"medium"`
	Large      []ImageData `json:"large"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Country   Country `json:"country"`
	Region    Region  `json:"region"`
}

type Region struct {
	City string `json:"city"`
}

type Timezone struct {
	Name     string `json:"name"`
	Offset   int    `json:"offset"`
	Abbr     string `json:"abbr"`
	AbbrName string `json:"abbrName"`
	IsDst    bool   `json:"isDst"`
}

// Inherits from Timezone object but includes an OffsetHours field.
type TimezoneWithOffsetHours struct {
	Timezone
	OffsetHours string `json:"offsetHours"`
}

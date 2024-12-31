package search

type SearchResultResponse struct {
	Results []SearchResults   `json:"results"`
	Stats   SearchResultStats `json:"stats"`
}

type SearchResults struct {
	ID     string             `json:"id"`
	Label  string             `json:"label"`
	Detail SearchResultDetail `json:"detail,omitempty"`
	Type   string             `json:"type"`
	Match  string             `json:"match"`
}

type SearchResultDetail struct {
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	SchdFrom   string  `json:"schd_from"`
	SchdTo     string  `json:"schd_to"`
	AcType     string  `json:"ac_type"`
	Route      string  `json:"route"`
	Logo       string  `json:"logo"`
	Reg        string  `json:"reg"`
	Callsign   string  `json:"callsign"`
	Flight     string  `json:"flight"`
	Operator   string  `json:"operator"`
	OperatorID int     `json:"operator_id"`
	Owner      string  `json:"owner"`
	Equip      string  `json:"equip"`
	Hex        string  `json:"hex"`
}

type SearchResultStats struct {
	Total SearchResultsTotal `json:"total"`
	Count SearchResultsCount `json:"count"`
}

type SearchResultsTotal struct {
	All      int `json:"all"`
	Airport  int `json:"airport"`
	Operator int `json:"operator"`
	Live     int `json:"live"`
	Schedule int `json:"schedule"`
	Aircraft int `json:"aircraft"`
}

type SearchResultsCount struct {
	Airport  int `json:"airport"`
	Operator int `json:"operator"`
	Live     int `json:"live"`
	Schedule int `json:"schedule"`
	Aircraft int `json:"aircraft"`
}

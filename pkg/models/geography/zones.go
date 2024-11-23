package geography

type Fr24Zones struct {
	Version       int          `json:"version"`
	Europe        Europe       `json:"europe"`
	Northamerica  Northamerica `json:"northamerica"`
	Southamerica  BaseZone     `json:"southamerica"`
	Oceania       BaseZone     `json:"oceania"`
	Asia          Asia         `json:"asia"`
	Africa        BaseZone     `json:"africa"`
	Atlantic      BaseZone     `json:"atlantic"`
	Maldives      BaseZone     `json:"maldives"`
	Northatlantic BaseZone     `json:"northatlantic"`
}

// General
type BaseZone struct {
	TlY float64 `json:"tl_y"`
	TlX float64 `json:"tl_x"`
	BrY float64 `json:"br_y"`
	BrX float64 `json:"br_x"`
}

// Europe
type UkSubzones struct {
	London  BaseZone `json:"london"`
	Ireland BaseZone `json:"ireland"`
}

type UkSubzone struct {
	BaseZone
	Subzones UkSubzones `json:"subzones"`
}

type Europe struct {
	BaseZone
	Subzones EuropeSubzones `json:"subzones"`
}

type EuropeSubzones struct {
	Poland      BaseZone  `json:"poland"`
	Germany     BaseZone  `json:"germany"`
	Uk          UkSubzone `json:"uk"`
	Spain       BaseZone  `json:"spain"`
	France      BaseZone  `json:"france"`
	Ceur        BaseZone  `json:"ceur"`
	Scandinavia BaseZone  `json:"scandinavia"`
	Italy       BaseZone  `json:"italy"`
}

// NorthAmerica
type NorthAmericaCentralSubSubzones struct {
	NaCny BaseZone `json:"na_cny"`
	NaCla BaseZone `json:"na_cla"`
	NaCat BaseZone `json:"na_cat"`
	NaCse BaseZone `json:"na_cse"`
	NaNw  BaseZone `json:"na_nw"`
	NaNe  BaseZone `json:"na_ne"`
	NaSw  BaseZone `json:"na_sw"`
	NaSe  BaseZone `json:"na_se"`
	NaCc  BaseZone `json:"na_cc"`
}

type NorthAmericaCentralSubzone struct {
	BaseZone
	Subzones NorthAmericaCentralSubSubzones `json:"subzones"`
}

type NorthAmericaSubzones struct {
	NaN BaseZone                   `json:"na_n"`
	NaC NorthAmericaCentralSubzone `json:"na_c"`
	NaS BaseZone                   `json:"na_s"`
}

type Northamerica struct {
	BaseZone
	Subzones NorthAmericaSubzones `json:"subzones"`
}

// Asia
type AsiaSubzones struct {
	Japan BaseZone `json:"japan"`
}
type Asia struct {
	BaseZone
	Subzones AsiaSubzones `json:"subzones"`
}

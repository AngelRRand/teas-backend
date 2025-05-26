package models

type TeaCategory struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Teas        []Te   `json:"teas"`
}

type Te struct {
	ID                   int          `json:"id"`
	Nombre               string       `json:"name"`
	Origen               string       `json:"origin"`
	ConsumptionCountries []string     `json:"consumption_countries"`
	Ingredients          []Ingredient `json:"ingredients"`
	Sabor                string       `json:"flavor_profile"`
	Beneficio            string       `json:"benefits"`
	Image                string       `json:"image"`
	ImageWhite           string       `json:"imageWhite"`
}

type Ingredient struct {
	Name               string `json:"name"`
	Origin             string `json:"origin"`
	Processing         string `json:"processing"`
	FlavorContribution string `json:"flavor_contribution"`
}

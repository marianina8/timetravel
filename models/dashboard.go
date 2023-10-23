package models

type Dashboard struct {
	Destination   string `json:"destination"`
	Present       string `json:"present"`
	LastDeparture string `json:"lastDeparture"`
	Completed     bool   `json:"completed"`
}

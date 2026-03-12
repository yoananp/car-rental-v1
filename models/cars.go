package models

type Car struct {
	ID           int    `json:"id"`
	Brand        string `json:"brand"`
	Type         string `json:"type"`
	Transmission string `json:"transmission"`
	PlateNumber  string `json:"plate_number"`
	PricePerDay  int    `json:"price_per_day"`
	Available    bool   `json:"available"`
}

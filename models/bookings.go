package models

type Booking struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	CarID      int    `json:"car_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`
}

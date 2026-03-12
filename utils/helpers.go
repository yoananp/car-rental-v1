package utils

import (
	"database/sql"
	"errors"
	"time"
)

// Cek apakah mobil bisa dibooking pada tanggal tertentu
func IsCarCanBooked(db *sql.DB, carId int, startDate string, endDate string) error {
	var count int
	err := db.QueryRow(
		`SELECT COUNT(*) FROM bookings
		WHERE car_id = ? AND
		((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?))`,
		carId, startDate, startDate, endDate, endDate).Scan(&count)
	if err != nil {
		return errors.New("error while checking car availability")
	}
	if count > 0 {
		return errors.New("car is not available for the selected dates")
	}
	return nil
}

// Hitung jumlah hari booking
func CalculateDays(startDate string, endDate string) (int, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return 0, errors.New("invalid start date format")
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return 0, errors.New("invalid end date format")
	}
	if start.After(end) {
		return 0, errors.New("start date cannot be after end date")
	}
	days := int(end.Sub(start).Hours() / 24)
	return days, nil
}

// Hitung total harga booking
func CalculateTotalPrice(dailyRate float64, days int) float64 {
	return dailyRate * float64(days)
}

// Validasi tanggal booking
func ValidateBookingDates(startDate string, endDate string) error {
	layout := "2006-01-02"
	start, err := time.Parse(layout, startDate)
	if err != nil {
		return errors.New("invalid start date format, use YYYY-MM-DD")
	}
	end, err := time.Parse(layout, endDate)
	if err != nil {
		return errors.New("invalid end date format, use YYYY-MM-DD")
	}
	if start.After(end) {
		return errors.New("start date must be before end date")
	}
	if start.Equal(end) {
		return errors.New("start date and end date cannot be the same")
	}
	return nil
}

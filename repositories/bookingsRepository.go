package repository

import (
	"database/sql"
	"errors"

	database "github.com/yoananp/car-rental-v1/database/sql_migration"
)

// Ambil semua booking
func GetAllBooking(db *sql.DB) ([]database.Booking, error) {
	rows, err := db.Query("SELECT id, customer_id, car_id, start_date, end_date FROM bookings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []database.Booking
	for rows.Next() {
		var b database.Booking
		err := rows.Scan(&b.ID, &b.CustomerID, &b.CarID, &b.StartDate, &b.EndDate)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

// Tambah booking
func InsertBooking(db *sql.DB, booking database.Booking) error {
	_, err := db.Exec(
		"INSERT INTO bookings(customer_id, car_id, start_date, end_date, total_price, status) VALUES($1,$2,$3,$4,$5,$6)",
		booking.CustomerID, booking.CarID, booking.StartDate, booking.EndDate, booking.TotalPrice, booking.Status,
	)
	return err
}

// Update booking
func UpdateBooking(db *sql.DB, booking database.Booking, bookingId int) error {
	res, err := db.Exec(
		"UPDATE bookings SET customer_id=$1, car_id=$2, start_date=$3, end_date=$4, total_price=$5, status=$6 WHERE id=$7",
		booking.CustomerID, booking.CarID, booking.StartDate, booking.EndDate, booking.TotalPrice, booking.Status, bookingId,
	)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Booking not found")
	}
	return nil
}

// Hapus booking
func DeleteBooking(db *sql.DB, booking database.Booking) error {
	res, err := db.Exec("DELETE FROM bookings WHERE id=$1", booking.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Booking not found")
	}
	return nil
}

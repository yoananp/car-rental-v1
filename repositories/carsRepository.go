package repository

import (
	"database/sql"
	"errors"

	database "github.com/yoananp/car-rental-v1/database/sql_migration"
)

// Ambil semua mobil
func GetAllCar(db *sql.DB) ([]database.Car, error) {
	rows, err := db.Query("SELECT id, brand, type, transmission, plate_number, price_per_day, available FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []database.Car
	for rows.Next() {
		var c database.Car
		err := rows.Scan(&c.ID, &c.Brand, &c.Type, &c.Transmission, &c.PlateNumber, &c.PricePerDay, &c.Available)
		if err != nil {
			return nil, err
		}
		cars = append(cars, c)
	}
	return cars, nil
}

// Tambah mobil
func InsertCar(db *sql.DB, car database.Car) error {
	_, err := db.Exec(
		"INSERT INTO cars(brand, type, transmission, plate_number, price_per_day, available) VALUES($1,$2,$3,$4,$5,$6)",
		car.Brand, car.Type, car.Transmission, car.PlateNumber, car.PricePerDay, car.Available,
	)
	return err
}

// Update mobil
func UpdateCar(db *sql.DB, car database.Car, carId int) error {
	res, err := db.Exec(
		"UPDATE cars SET brand=$1, type=$2, transmission=$3, plate_number=$4, price_per_day=$5, available=$6 WHERE id=$7",
		car.Brand, car.Type, car.Transmission, car.PlateNumber, car.PricePerDay, car.Available, carId,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Car not found")
	}
	return nil
}

// Hapus mobil
func DeleteCar(db *sql.DB, car database.Car) error {
	res, err := db.Exec("DELETE FROM cars WHERE id=$1", car.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Car not found")
	}
	return nil
}

// Cek harga mobil
func GetCarPrice(db *sql.DB, id int) (int, error) {
	var price int
	err := db.QueryRow("SELECT price_per_day FROM cars WHERE id=$1", id).Scan(&price)
	if err != nil {
		return 0, errors.New("Car not found")
	}
	return price, nil
}

// Cek apakah mobil bisa dibooking (sesuai tabel bookings)
func IsCarCanBooked(db *sql.DB, carId int, dateStart, dateFinish string) error {
	var id int
	query := `SELECT id FROM bookings 
			  WHERE car_id=$1 AND (status='Pending' OR status='Confirmed') 
			  AND ( $2 <= end_date AND $3 >= start_date )`
	err := db.QueryRow(query, carId, dateStart, dateFinish).Scan(&id)
	if err != nil {
		// err = sql.ErrNoRows => mobil bisa dibooking
		return nil
	}
	return errors.New("Car is already booked for this period")
}

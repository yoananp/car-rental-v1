package repository

import (
	"database/sql"
	"errors"

	database "github.com/yoananp/car-rental-v1/database/sql_migration"
)

// Ambil semua customer
func GetAllCustomer(db *sql.DB) ([]database.Customer, error) {
	rows, err := db.Query("SELECT id, name, nik, phone_number FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []database.Customer
	for rows.Next() {
		var c database.Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.NIK, &c.PhoneNumber); err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// Tambah customer
func InsertCustomer(db *sql.DB, customer database.Customer) error {
	_, err := db.Exec(
		"INSERT INTO customers(name, nik, phone_number) VALUES($1,$2,$3)",
		customer.Name, customer.NIK, customer.PhoneNumber,
	)
	return err
}

// Update customer
func UpdateCustomer(db *sql.DB, customer database.Customer, customerId int) error {
	res, err := db.Exec(
		"UPDATE customers SET name=$1, nik=$2, phone_number=$3 WHERE id=$4",
		customer.Name, customer.NIK, customer.PhoneNumber, customerId,
	)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Customer not found")
	}
	return nil
}

// Hapus customer
func DeleteCustomer(db *sql.DB, customer database.Customer) error {
	res, err := db.Exec("DELETE FROM customers WHERE id=$1", customer.ID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Customer not found")
	}
	return nil
}

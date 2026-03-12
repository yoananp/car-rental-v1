package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DbConnection *sql.DB // Variabel global untuk koneksi database

// ConnectDB membuat koneksi ke Postgres
func ConnectDB(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Database connected!")
	DbConnection = db // Menyimpan koneksi ke variabel global
	return db, nil
}

// DbMigrate placeholder untuk migrasi DB
func DbMigrate(db *sql.DB) {
	// Implementasi migrasi database
	fmt.Println("Running migrations... (implement your migration logic here)")
}

// Struktur untuk tabel yang digunakan
type Car struct {
	ID           int
	Brand        string
	Type         string
	Transmission string
	PlateNumber  string
	PricePerDay  int
	Available    bool
}

type Customer struct {
	ID          int
	Name        string
	NIK         string
	PhoneNumber string
}

type Booking struct {
	ID         int
	CustomerID int
	CarID      int
	StartDate  string
	EndDate    string
	TotalPrice int
	Status     string
}

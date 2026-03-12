package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/yoananp/car-rental-v1/database/sql_migration"
	repository "github.com/yoananp/car-rental-v1/repositories"
)

// GetAllCar ambil semua mobil
func GetAllCar(c *gin.Context) {
	cars, err := repository.GetAllCar(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

// InsertCar tambah mobil baru
func InsertCar(c *gin.Context) {
	var car database.Car
	// Bind JSON to Car struct
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.InsertCar(database.DbConnection, car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Car inserted"})
}

// UpdateCar update mobil berdasarkan ID
func UpdateCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	var car database.Car
	// Bind JSON to Car struct
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = repository.UpdateCar(database.DbConnection, car, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Car updated"})
}

// DeleteCar hapus mobil berdasarkan ID
func DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	// Create Car object
	var car database.Car
	car.ID = id

	err = repository.DeleteCar(database.DbConnection, car)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Car deleted"})
}

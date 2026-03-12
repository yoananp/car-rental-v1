package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/yoananp/car-rental-v1/database/sql_migration"
	repository "github.com/yoananp/car-rental-v1/repositories"
)

// GetAllCustomer ambil semua customer
func GetAllCustomer(c *gin.Context) {
	customers, err := repository.GetAllCustomer(database.DbConnection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

// InsertCustomer tambah customer baru
func InsertCustomer(c *gin.Context) {
	var customer database.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.InsertCustomer(database.DbConnection, customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Customer inserted successfully"})
}

// UpdateCustomer update customer
func UpdateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	var customer database.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repository.UpdateCustomer(database.DbConnection, customer, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully"})
}

// DeleteCustomer hapus customer
func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
		return
	}

	customer := database.Customer{ID: id}
	if err := repository.DeleteCustomer(database.DbConnection, customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

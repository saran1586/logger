package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {

	var employee Employee

	err := c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	DB.Create(&employee)

	c.JSON(http.StatusCreated, employee)
}

func GetEmployees(c *gin.Context) {

	var employees []Employee

	DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func UpdateEmployee(c *gin.Context) {

	id := c.Param("id")

	var employee Employee

	if err := DB.First(&employee, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Employee not found",
		})

		return
	}

	var update Employee

	c.ShouldBindJSON(&update)

	employee.Name = update.Name
	employee.Salary = update.Salary

	DB.Save(&employee)

	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {

	id := c.Param("id")

	var employee Employee

	if err := DB.First(&employee, id).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Employee not found",
		})

		return
	}

	DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee deleted",
	})
}

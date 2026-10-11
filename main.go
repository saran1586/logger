package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	InitLogger()

	ConnectDB()

	r := gin.Default()

	r.POST("/employees", CreateEmployee)
	r.GET("/employees", GetEmployees)
	r.PUT("/employees/:id", UpdateEmployee)
	r.DELETE("/employees/:id", DeleteEmployee)

	Logger.Println("Server started on port 8080")

	log.Println("Server running on :8080")

	r.Run(":8080")
}

package main

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {

	file, err := os.OpenFile(
		"employee.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	Logger = log.New(
		file,
		"INFO: ",
		log.Ldate|log.Ltime,
	)
}

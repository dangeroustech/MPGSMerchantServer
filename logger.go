package main

import (
	"log"
	"os"
)

//Logger - logs things
func Logger(msg string) {
	f, err := os.OpenFile("sessions.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "MPGS_API", log.LstdFlags)
	logger.Println(msg)
}

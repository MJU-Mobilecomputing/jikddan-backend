package main

import (
	"log"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/app"
)

func main() {
	application, err := app.InitApplication()
	if err != nil {
		log.Fatal(err)
	}
	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}
}

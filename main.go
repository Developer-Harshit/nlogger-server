package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	XATA_URL string
	XATA_API_KEY string
	XATA_TABLE string
)

func main() {
	XATA_URL = os.Getenv("XATA_DATABASE_URL")
	XATA_API_KEY = os.Getenv("XATA_API_KEY")
	XATA_TABLE = os.Getenv("XATA_TABLE_NAME")



	if XATA_URL == "" || XATA_API_KEY == "" || XATA_TABLE == "" {
		panic("INVALID OR UNSET CREDENTIALS")
	}

	assignHandlers()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Server listening on PORT: " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error while starting server:", err)
	}
}

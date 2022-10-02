package main

import (
	"fmt"
	"net/http"

	"github.com/emilieanthony/bike-rental-service/src/handlers"
)

func main(){
	fmt.Println("Starting API...")
	http.HandleFunc("/rentals", handlers.RentalHandler)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)	
}
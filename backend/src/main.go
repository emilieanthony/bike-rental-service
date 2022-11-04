package main

import (
	"fmt"

	"github.com/emilieanthony/rental-service/backend/src/handlers"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Starting API...")
	server := gin.Default()

	// routes
	server.POST("/rentals", handlers.CreateRental)
	server.GET("/rentals/:id", handlers.GetRentalsById)
	server.GET("/rentals", handlers.GetAllRentals)

	fmt.Println("Listening on port 3000...")
	server.Run(":3000")
}






package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type rentalBody struct{
	UserEmail string `json:"userEmail"`
	BikeID string `json:"bikeID"`
}

type rental struct {
	ID string
	UserEmail string
	BikeID string
}

var rentals []rental

func CreateRental(ctx *gin.Context){

	var body rentalBody

	if err := ctx.BindJSON(&body); err != nil {
		ctx.String(http.StatusBadRequest, "invalid input")
		return
	}
	newRental := rental {
		ID: uuid.New().String(),
		UserEmail: body.UserEmail,
		BikeID: body.BikeID,
	}

	rentals = append(rentals, newRental)
	ctx.JSON(http.StatusCreated, newRental)
}

func GetRentalsById(ctx *gin.Context){
	id := ctx.Param("id")
	
	for _, r := range rentals {
		if r.ID == id {
			ctx.JSON(http.StatusOK, r)
			return
		}
	}
	ctx.String(http.StatusNotFound, "rental with id %s not found", id)
}
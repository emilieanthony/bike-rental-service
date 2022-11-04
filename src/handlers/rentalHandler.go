package handlers

import (
	"encoding/json"
	"net/http"
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

func RentalHandler(w http.ResponseWriter, req *http.Request){
	
	if req.Method == http.MethodPost{
		var body rentalBody
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = createRental(body.UserEmail, body.BikeID) 
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}
	if req.Method == http.MethodGet{
		// TODO: Get by id - refactor to gin
		j, err := json.Marshal(rentals)
		if err != nil{
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(j)
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
	return
}

func createRental(userEmail string, bikeID string) error {
	rentals = append(rentals, rental{
		ID: "",
		UserEmail: userEmail,
		BikeID: bikeID,
	})
	return nil
}

func getRentalById (bikeID string) *rental {
	for _, s := range rentals {
		if s.BikeID == bikeID{
			return &s
		}
	} 
	return nil
}
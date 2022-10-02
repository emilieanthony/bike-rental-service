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
	userEmail string
	bikeID string
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

	w.WriteHeader(http.StatusNotImplemented)
	return

}

func createRental(userEmail string, bikeID string) error {
	rentals = append(rentals, rental{
		ID: "",
		userEmail: userEmail,
		bikeID: bikeID,
	})
	return nil
}
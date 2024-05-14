package mupVozilaHandlers

import (
	"encoding/json"
	"mupvozila_service/data"
	"net/http"
)

// IssueLicenseHandler handles requests to issue driver's licenses
func IssueLicenseHandler(w http.ResponseWriter, r *http.Request) {
	var license data.License
	err := json.NewDecoder(r.Body).Decode(&license)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.InsertLicense(&license)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the inserted license ID
	json.NewEncoder(w).Encode(result.InsertedID)
}

// RegisterVehicleHandler handles requests to register vehicles
func RegisterVehicleHandler(w http.ResponseWriter, r *http.Request) {
	var vehicle data.RegisterVehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.InsertVehicle(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the inserted vehicle ID
	json.NewEncoder(w).Encode(result.InsertedID)
}

package mupVozilaHandlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	result, err := data.InsertRegisteredVehicle(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the inserted vehicle ID
	json.NewEncoder(w).Encode(result.InsertedID)
}

// GetAllLicensesHandler handles requests to retrieve all driver's licenses
func GetAllLicensesHandler(w http.ResponseWriter, r *http.Request) {
	licenses, err := data.GetAllLicenses()
	if err != nil {
		log.Println("Error retrieving licenses:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved licenses
	json.NewEncoder(w).Encode(licenses)
}

// GetAllVehiclesHandler handles requests to retrieve all registered vehicles
func GetAllVehiclesHandler(w http.ResponseWriter, r *http.Request) {
	vehicles, err := data.GetAllVehicles()
	if err != nil {
		log.Println("Error retrieving vehicles:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved vehicles
	json.NewEncoder(w).Encode(vehicles)
}

// GetLicenseByIDHandler handles requests to retrieve a driver's license by its ID
func GetLicenseByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	licenseID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	license, err := data.GetLicenseByID(licenseID)
	if err != nil {
		log.Println("Error retrieving license:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved license
	json.NewEncoder(w).Encode(license)
}

// GetVehicleByIDHandler handles requests to retrieve a vehicle by its ID
func GetVehicleByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vehicle, err := data.GetVehicleByID(vehicleID)
	if err != nil {
		log.Println("Error retrieving vehicle:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved vehicle
	json.NewEncoder(w).Encode(vehicle)
}

// GetLicencesByUserID retrieves licenses by user ID
func GetLicencesByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		licences, err := data.GetLicencesByUserID(dbClient, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(licences)
	}
}

// GetVehicleByRegistrationHandler handles requests to retrieve a vehicle by its registration
func GetVehicleByRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registration := params["registration"]

	vehicle, err := data.GetVehicleByRegistration(registration)
	if err != nil {
		log.Println("Error retrieving vehicle:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if vehicle == nil {
		w.Write([]byte("OK"))
		return
	}

	// Respond with the retrieved vehicle
	json.NewEncoder(w).Encode(vehicle)
}

// GetAllRegistrationsHandler handles requests to retrieve all registered vehicles
func GetAllRegistrationsHandler(w http.ResponseWriter, r *http.Request) {
	vehicles, err := data.GetAllRegistrations()
	if err != nil {
		log.Println("Error retrieving registrations:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved registrations
	json.NewEncoder(w).Encode(vehicles)
}

// InsertCarHandler handles requests to insert a car into the database
func InsertCarHandler(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := data.InsertCar(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the inserted car ID
	json.NewEncoder(w).Encode(result.InsertedID)
}

// GetAllCarsHandler handles requests to retrieve all cars
func GetAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars, err := data.GetAllCars()
	if err != nil {
		log.Println("Error retrieving cars:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved cars
	json.NewEncoder(w).Encode(cars)
}

// GetCarsByUserIDHandler handles requests to retrieve cars by user ID
func GetCarsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := primitive.ObjectIDFromHex(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	cars, err := data.GetCarsByUserID(userID)
	if err != nil {
		log.Println("Error retrieving cars by user ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved cars
	json.NewEncoder(w).Encode(cars)
}

// DeleteCarByIDHandler handles requests to delete a car by its ID
func DeleteCarByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	carID, err := primitive.ObjectIDFromHex(params["car_id"])
	if err != nil {
		http.Error(w, "Invalid carID", http.StatusBadRequest)
		return
	}

	err = data.DeleteCarByID(carID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

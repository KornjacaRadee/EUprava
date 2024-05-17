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

	result, err := data.InsertVehicle(&vehicle)
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

func GetLicencesByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("a daj brate")
		log.Println("This is id: ", mux.Vars(r)["id"])

		userID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		log.Println("This is id in object: ", userID)
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

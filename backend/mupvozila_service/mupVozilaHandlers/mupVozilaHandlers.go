package mupVozilaHandlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"mupvozila_service/client"
	"mupvozila_service/data"
	"net/http"
)

type UpdateValidityRequest struct {
	IsValid  bool   `json:"is_valid"`
	category string `json:"category"`
}

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

// GetLicensesByUserJMBGHandler handles requests to retrieve licenses by user's JMBG
func GetLicensesByUserJMBGHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userJMBG := params["jmbg"]

	licenses, err := data.GetLicensesByUserJMBG(userJMBG)
	if err != nil {
		log.Println("Error retrieving licenses:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved licenses
	json.NewEncoder(w).Encode(licenses)
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

func UpdateLicenseValidityHandler(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID := params["id"]

		category := params["category"]

		var req UpdateValidityRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Updating licenses for user ID: == and category: %s to IsValid: %v\n", category, req.IsValid)

		result, err := data.UpdateLicenseValidityByUserIDAndCategory(userID, category, req.IsValid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)
	}
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

	log.Println("Fetching cars for user ID:", userID) // Log the user ID being queried
	cars, err := data.GetCarsByUserID(userID)
	if err != nil {
		log.Println("Error retrieving cars by user ID:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Cars retrieved:", cars) // Log the retrieved cars
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

// DeleteRegistrationHandler handles requests to delete a registration by ID
func DeleteRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	registrationID := params["id"] // Assuming your URL path parameter is named "id"

	// Convert registration ID from string to MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(registrationID)
	if err != nil {
		http.Error(w, "Invalid registration ID", http.StatusBadRequest)
		return
	}

	// Call the data layer function to delete the registration
	err = data.DeleteRegistrationByID(objectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message or appropriate response
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// DeleteLicenseByIDHandler handles requests to delete a license by its ID
func DeleteLicenseByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	licenseID, err := primitive.ObjectIDFromHex(params["license_id"])
	if err != nil {
		http.Error(w, "Invalid license ID", http.StatusBadRequest)
		return
	}

	err = data.DeleteLicenseByID(licenseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// UpdateCarHandler handles requests to update a car
func UpdateCarHandler(w http.ResponseWriter, r *http.Request) {
	var car data.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.UpdateCar(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// UpdateRegistrationHandler handles requests to update a registration
func UpdateRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var registration data.RegisterVehicle
	err := json.NewDecoder(r.Body).Decode(&registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.UpdateRegistration(&registration)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// GetCarByLicensePlateHandler vraća automobil na osnovu registracionih tablica
func GetCarByLicensePlateHandler(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		licensePlate := params["license_plate"]

		car, err := data.GetCarByLicensePlate(r.Context(), licensePlate)
		if err != nil {
			if err.Error() == "car with license plate "+licensePlate+" not found" {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(car); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

// UpdateLicenseHandler handles requests to update a license
func UpdateLicenseHandler(w http.ResponseWriter, r *http.Request) {
	var license data.License
	err := json.NewDecoder(r.Body).Decode(&license)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.UpdateLicense(&license)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// GetLicencesByUserID retrieves licenses by user ID
func GetLicencesByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		log.Println("Fetching licenses for user ID:", userID) // Log the user ID being queried
		licences, err := data.GetLicencesByUserID(dbClient, userID)
		if err != nil {
			log.Println("Error retrieving licenses by user ID:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("Licenses retrieved:", licences) // Log the retrieved licenses
		json.NewEncoder(w).Encode(licences)
	}
}

// Handler struct for dependencies
type Handler struct {
	saobracajnaClient *client.SaobracajnaPolicijaClient
}

func NewHandler(saobracajnaClient *client.SaobracajnaPolicijaClient) *Handler {
	return &Handler{
		saobracajnaClient: saobracajnaClient,
	}
}

// GetNesreceByVozacHandler handles requests to get nesrece by vozac
func (h *Handler) GetNesreceByVozacHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vozac := params["vozac"]

	nesrece, err := h.saobracajnaClient.GetAllNesreceByVozac(context.Background(), vozac)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(nesrece); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

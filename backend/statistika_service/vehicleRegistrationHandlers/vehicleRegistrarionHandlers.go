package vehicleRegistrationHandlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"statistika_service/config"
	"statistika_service/data"
)

type KeyProduct struct{}
type VehicleRegistrationHandler struct {
	logger *config.Logger
	repo   *data.VehicleRegistrationRepo
}

func NewVehicleRegistrationHandler(l *config.Logger, r *data.VehicleRegistrationRepo) *VehicleRegistrationHandler {
	return &VehicleRegistrationHandler{l, r}
}

func (a *VehicleRegistrationHandler) GetAllVehicles(rw http.ResponseWriter, h *http.Request) {
	registrations, err := a.repo.GetAll()
	if err != nil {
		a.logger.Print("Database exception: ", err)
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		return
	}

	if registrations == nil {
		http.Error(rw, "No registrations found", http.StatusNotFound)
		return
	}

	// Convert registrations to JSON
	jsonData, err := json.Marshal(registrations)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		a.logger.Fatalf("Unable to convert to json: %v", err)
		return
	}

	// Set Content-Type header
	rw.Header().Set("Content-Type", "application/json")

	// Write JSON response
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonData)
}

func (a *VehicleRegistrationHandler) GetVehicle(rw http.ResponseWriter, h *http.Request) {

	vars := mux.Vars(h)
	id := vars["id"]

	registrations, err := a.repo.GetRegistrationByID(id)
	if err != nil {
		a.logger.Print("Database exception: ", err)
	}

	if registrations.Id.Hex() != id {
		http.Error(rw, "Accommodation not found", 404)
		return
	}

	err = registrations.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		a.logger.Fatalf("Unable to convert to json :", err)
		return
	}
}

func (a *VehicleRegistrationHandler) PostVehicleRegistration(rw http.ResponseWriter, h *http.Request) {
	/*	tokenString := h.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(rw, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Remove 'Bearer ' prefix if present
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		role, err := getRoleFromToken(tokenString)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error extracting user role: %v", err), http.StatusUnauthorized)
			return
		}

		// Check if the user has the required role
		if role != "host" {
			http.Error(rw, "Unauthorized: Insufficient privileges", http.StatusUnauthorized)
			return
		}

		// Extract user ID from the token
		userID, err := getUserIdFromToken(tokenString)
		if err != nil {
			http.Error(rw, fmt.Sprintf("Error extracting user ID: %v", err), http.StatusUnauthorized)
			return
		}
	*/
	// Create new accommodation with the extracted user ID as the owner
	registration := h.Context().Value(KeyProduct{}).(*data.VehicleRegistration)
	///ovo vjerovatno nije dobro
	// Insert the accommodation
	erra := a.repo.InsertRegistration(registration)
	if erra != nil {
		http.Error(rw, "Unable to post accommodation", http.StatusBadRequest)
		a.logger.Fatalf("Unable to post accommodation", erra)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (a *VehicleRegistrationHandler) PatchVehicleRegistration(rw http.ResponseWriter, h *http.Request) {
	/*tokenString := h.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(rw, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	// Remove 'Bearer ' prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	role, err := getRoleFromToken(tokenString)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error extracting user role: %v", err), http.StatusUnauthorized)
		return
	}

	// Check if the user has the required role
	if role != "host" {
		http.Error(rw, "Unauthorized: Insufficient privileges", http.StatusUnauthorized)
		return
	}*/
	vars := mux.Vars(h)
	id := vars["id"]
	registration := h.Context().Value(KeyProduct{}).(*data.VehicleRegistration)

	a.repo.UpdateRegistration(id, registration)
	rw.WriteHeader(http.StatusOK)
}

func (a *VehicleRegistrationHandler) DeleteVehicleRegistration(rw http.ResponseWriter, h *http.Request) {
	/**tokenString := h.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(rw, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	// Remove 'Bearer ' prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	role, err := getRoleFromToken(tokenString)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error extracting user role: %v", err), http.StatusUnauthorized)
		return
	}

	// Extract user ID from the token
	userID, err := getUserIdFromToken(tokenString)
	if err != nil {
		http.Error(rw, fmt.Sprintf("Error extracting user ID: %v", err), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(h)
	id := vars["id"]

	// Check if the user has the required role or is the owner of the accommodation
	if role != "host" {
		http.Error(rw, "Unauthorized: Insufficient privileges", http.StatusUnauthorized)
		return
	}

	// Provjeri da li je korisnik vlasnik smjestaja
	accommodation, err := a.repo.GetByID(id) // Use the new GetByID function
	if err != nil {
		http.Error(rw, "Error getting accommodation", http.StatusInternalServerError)
		a.logger.Fatalf("Error getting accommodation", err)
		return
	}

	idUser, _ := primitive.ObjectIDFromHex(userID)

	if accommodation.Owner.Id != idUser {
		http.Error(rw, "Unauthorized: User is not the owner of the accommodation", http.StatusUnauthorized)
		return
	}
	*/
	vars := mux.Vars(h)
	id := vars["id"]

	a.repo.DeleteRegistration(id)
	rw.WriteHeader(http.StatusOK)
}
func (a *VehicleRegistrationHandler) MiddlewareVehicleDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		registration := &data.VehicleRegistration{}
		err := registration.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			a.logger.Fatalf("Unable to decode json", err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyProduct{}, registration)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
func (a *VehicleRegistrationHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		a.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

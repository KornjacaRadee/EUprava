package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"saobracajna_policija/data"
)

type SaobracajnaPolicijaHandler struct {
	repo   *data.SaobracajnaPolicijaRepo
	logger *log.Logger
}

func NewSaobracajnaPolicijaHandler(repo *data.SaobracajnaPolicijaRepo, logger *log.Logger) *SaobracajnaPolicijaHandler {
	return &SaobracajnaPolicijaHandler{
		repo:   repo,
		logger: logger,
	}
}

func (h *SaobracajnaPolicijaHandler) GetLicensesByUserJMBG(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]

	licenses, err := h.repo.GetLicensesByUserJMBG(r.Context(), jmbg)
	if err != nil {
		http.Error(w, "Error retrieving licenses", http.StatusInternalServerError)
		return
	}

	if licenses == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(licenses)
}

func (h *SaobracajnaPolicijaHandler) GetAllNesreceByVozac(w http.ResponseWriter, r *http.Request) {
	vozac := mux.Vars(r)["vozac"]
	nesrece, err := h.repo.GetAllNesreceByVozac(r.Context(), vozac)
	if err != nil {
		h.logger.Println("Error fetching accidents by driver:", err)
		http.Error(w, "Error fetching accidents", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nesrece)
}

func (h *SaobracajnaPolicijaHandler) GetAllCars(w http.ResponseWriter, r *http.Request) {
	cars, err := h.repo.GetAllCars(r.Context())
	if err != nil {
		http.Error(w, "Greška pri dobijanju vozila", http.StatusInternalServerError)
		return
	}

	// Vraćanje odgovora sa JSON sadržajem
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
func (h *SaobracajnaPolicijaHandler) GetCarByLicensePlate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	licensePlate := vars["license_plate"]

	// Log the received license plate
	h.logger.Printf("Received request for car with license plate: %s", licensePlate)

	// Retrieve the car by license plate
	car, err := h.repo.GetCarByLicensePlate(r.Context(), licensePlate)
	if err != nil {
		// Log the error for debugging
		h.logger.Printf("Error retrieving car by license plate %s: %v", licensePlate, err)
		http.Error(w, "Error retrieving car", http.StatusInternalServerError)
		return
	}

	// Check if the car was found
	if car == nil {
		h.logger.Printf("Car with license plate %s not found", licensePlate)
		http.NotFound(w, r)
		return
	}

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the car details as JSON
	if err := json.NewEncoder(w).Encode(car); err != nil {
		// Log the error and send an internal server error response
		h.logger.Printf("Error encoding car to JSON: %v", err)
		http.Error(w, "Error encoding car to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *SaobracajnaPolicijaHandler) GetUserByJMBG(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	jmbg := vars["jmbg"]

	user, err := h.repo.GetUserByJMBG(r.Context(), jmbg)
	if err != nil {
		h.logger.Printf("Error retrieving user by JMBG %s: %v", jmbg, err)
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		h.logger.Printf("Error encoding user to JSON: %v", err)
		http.Error(w, "Error encoding user to JSON", http.StatusInternalServerError)
		return
	}
}

func (h *SaobracajnaPolicijaHandler) CreateNesreca(w http.ResponseWriter, r *http.Request) {
	var nesreca data.Nesreca
	err := json.NewDecoder(r.Body).Decode(&nesreca)
	if err != nil {
		http.Error(w, "Greška pri parsiranju zahteva", http.StatusBadRequest)
		return
	}

	err = h.repo.CreateNesreca(r.Context(), &nesreca)
	if err != nil {
		http.Error(w, "Greška pri kreiranju nesreće", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SaobracajnaPolicijaHandler) CreatePrekrsaj(w http.ResponseWriter, r *http.Request) {
	var prekrsaj data.Prekrsaj
	err := json.NewDecoder(r.Body).Decode(&prekrsaj)
	if err != nil {
		http.Error(w, "Greška pri parsiranju zahteva", http.StatusBadRequest)
		return
	}

	err = h.repo.CreatePrekrsaj(r.Context(), &prekrsaj)
	if err != nil {
		http.Error(w, "Greška pri kreiranju prekršaja", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SaobracajnaPolicijaHandler) GetNesrece(w http.ResponseWriter, r *http.Request) {
	nesrece, err := h.repo.GetNesrece(r.Context())
	if err != nil {
		http.Error(w, "Greška pri dobijanju nesreća", http.StatusInternalServerError)
		return
	}

	// Vraćanje odgovora sa JSON sadržajem
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nesrece)
}

func (h *SaobracajnaPolicijaHandler) GetPrekrsaji(w http.ResponseWriter, r *http.Request) {
	prekrsaji, err := h.repo.GetPrekrsaji(r.Context())
	if err != nil {
		http.Error(w, "Greška pri dobijanju prekršaja", http.StatusInternalServerError)
		return
	}

	// Vraćanje odgovora sa JSON sadržajem
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prekrsaji)
}

func (h *SaobracajnaPolicijaHandler) DeleteNesreca(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteNesreca(r.Context(), id)
	if err != nil {
		http.Error(w, "Error deleting nesreca", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SaobracajnaPolicijaHandler) DeletePrekrsaj(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.repo.DeletePrekrsaj(r.Context(), id)
	if err != nil {
		http.Error(w, "Error deleting prekrsaj", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

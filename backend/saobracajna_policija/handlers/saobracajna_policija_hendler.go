package handler

import (
	"encoding/json"
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

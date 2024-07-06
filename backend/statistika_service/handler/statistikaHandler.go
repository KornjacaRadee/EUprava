package handler

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"statistika_service/client"
	"statistika_service/data"
)

func CreatePrekrsajiStatistika(saobracajClient *client.SaobracajnaPolicijaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Prekrsaj data
		prekrsaji, err := saobracajClient.FetchPrekrsaji(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch prekrsaji: %v", prekrsaji), http.StatusInternalServerError)
			return
		}

		vehicleCounts := make(map[string]int)
		locationCounts := make(map[string]int)
		totalCrashes := 0

		for _, prekrsaj := range prekrsaji {
			vehicleCounts[prekrsaj.Vozilo]++
			locationCounts[prekrsaj.Lokacija]++
			totalCrashes++
		}

		// Find the most common vehicle and location
		var mostCommonVehicle string
		var mostCommonLocation string
		maxVehicleCount := 0
		maxLocationCount := 0

		for vehicle, count := range vehicleCounts {
			if count > maxVehicleCount {
				maxVehicleCount = count
				mostCommonVehicle = vehicle
			}
		}

		for location, count := range locationCounts {
			if count > maxLocationCount {
				maxLocationCount = count
				mostCommonLocation = location
			}
		}

		statistika := data.StatistikaPrekrsaja{
			ID:                    primitive.NewObjectID(),
			GodisnjiBrojPrekrsaja: fmt.Sprintf("%d", totalCrashes),
			NajcescaLokacija:      mostCommonLocation,
			NajcesceVozilo:        mostCommonVehicle,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(statistika)
	}
}
func CreateNesrecaStatistika(saobracajClient *client.SaobracajnaPolicijaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Nesreca data
		nesrece, err := saobracajClient.FetchNesrece(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch nesrece: %v", nesrece), http.StatusInternalServerError)
			return
		}

		vehicleCounts := make(map[string]int)
		locationCounts := make(map[string]int)
		totalCrashes := 0

		for _, nesreca := range nesrece {
			vehicleCounts[nesreca.Vozilo]++
			locationCounts[nesreca.Lokacija]++
			totalCrashes++
		}

		// Find the most common vehicle and location
		var mostCommonVehicle string
		var mostCommonLocation string
		maxVehicleCount := 0
		maxLocationCount := 0

		for vehicle, count := range vehicleCounts {
			if count > maxVehicleCount {
				maxVehicleCount = count
				mostCommonVehicle = vehicle
			}
		}

		for location, count := range locationCounts {
			if count > maxLocationCount {
				maxLocationCount = count
				mostCommonLocation = location
			}
		}

		statistika := data.StatistikaNesreca{
			ID:                  primitive.NewObjectID(),
			GodisnjiBrojNesreca: fmt.Sprintf("%d", totalCrashes),
			NajcescaLokacija:    mostCommonLocation,
			NajcesceVozilo:      mostCommonVehicle,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(statistika)
	}
}
func CreateVozackihDozvolaStatistika(mupClient *client.MupVozilaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch License data
		licenses, err := mupClient.FetchLicenses(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch licenses: %v", err), http.StatusInternalServerError)
			return
		}

		categoryCounts := make(map[string]int)
		totalLicenses := 0
		totalPoints := 0

		for _, license := range licenses {
			categoryCounts[license.Category]++
			totalLicenses++
			totalPoints += license.Points
		}

		// Find the most common category
		var mostCommonCategory string
		maxCategoryCount := 0

		for category, count := range categoryCounts {
			if count > maxCategoryCount {
				maxCategoryCount = count
				mostCommonCategory = category
			}
		}

		statistika := data.StatistikaVozackihDozvola{
			ID:                         primitive.NewObjectID(),
			GodisnjiBrojIzdatihDozvola: fmt.Sprintf("%d", totalLicenses),
			NajcescaKategorija:         mostCommonCategory,
			GodisnjiBrojKaznenihBodova: fmt.Sprintf("%d", totalPoints),
		}

		// Encode the statistika object into JSON and send as response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statistika); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
func CreateRegistrovanihVozilaStatistika(mupClient *client.MupVozilaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Registered Vehicles data
		registeredVehicles, err := mupClient.FetchRegisteredVehicles(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch registered vehicles: %v", err), http.StatusInternalServerError)
			return
		}

		vehicleCounts := make(map[string]int)
		totalRegisteredVehicles := 0
		totalDays := 0

		for _, vehicle := range registeredVehicles {
			vehicleCounts[vehicle.Name]++
			totalRegisteredVehicles++
			totalDays += int(vehicle.ValidUntilDate.Sub(vehicle.IssuingDate).Hours() / 24)
		}

		// Find the most common registered vehicle
		var mostCommonVehicle string
		maxVehicleCount := 0

		for vehicle, count := range vehicleCounts {
			if count > maxVehicleCount {
				maxVehicleCount = count
				mostCommonVehicle = vehicle
			}
		}

		// Calculate average validity period
		averageValidityPeriod := totalDays / totalRegisteredVehicles

		statistika := data.StatistikaRegistrovanihVozila{
			ID:                         primitive.NewObjectID(),
			BrojRegistrovanihVozila:    fmt.Sprintf("%d", totalRegisteredVehicles),
			ProsecniPeriodVazenja:      fmt.Sprintf("%d", averageValidityPeriod),
			NajcesceRegistrovanoVozilo: mostCommonVehicle,
		}

		// Encode the statistika object into JSON and send as response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statistika); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
func CreateStatistikaAuta(mupClient *client.MupVozilaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Car data
		cars, err := mupClient.FetchCars(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch cars: %v", err), http.StatusInternalServerError)
			return
		}

		modelCounts := make(map[string]int)
		yearCounts := make(map[int]int)
		totalCars := 0

		for _, car := range cars {
			modelCounts[car.Model]++
			yearCounts[car.Year]++
			totalCars++
		}

		// Find the most common model and year of production
		var mostCommonModel string
		var mostCommonYear int
		maxModelCount := 0
		maxYearCount := 0

		for model, count := range modelCounts {
			if count > maxModelCount {
				maxModelCount = count
				mostCommonModel = model
			}
		}

		for year, count := range yearCounts {
			if count > maxYearCount {
				maxYearCount = count
				mostCommonYear = year
			}
		}

		statistika := data.StatistikaAuta{
			ID:                        primitive.NewObjectID(),
			UkupanBrojVozila:          fmt.Sprintf("%d", totalCars),
			NajcesciModel:             mostCommonModel,
			NajcescaGodinaProizvodnje: fmt.Sprintf("%d", mostCommonYear),
		}

		// Encode the statistika object into JSON and send as response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statistika); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
func CreateNaloziZaPretresStatistika(lawCourtClient *client.LawCourtClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch House Search Warrants data
		nalozi, err := lawCourtClient.FetchHouseSearchWarrants(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch nalozi za pretres: %v", err), http.StatusInternalServerError)
			return
		}

		razlogCounts := make(map[string]int)
		korisnikCounts := make(map[string]int)
		locationCounts := make(map[string]int)
		totalNalozi := 0

		for _, nalog := range nalozi {
			razlogCounts[nalog.Title]++
			korisnikCounts[nalog.UserID.Hex()]++ // Convert ObjectID to string with Hex()
			locationCounts[nalog.Address]++
			totalNalozi++
		}

		// Find the most common razlog, korisnik and location
		var mostCommonRazlog string
		var mostCommonKorisnik string
		var mostCommonLocation string
		maxRazlogCount := 0
		maxKorisnikCount := 0
		maxLocationCount := 0

		for razlog, count := range razlogCounts {
			if count > maxRazlogCount {
				maxRazlogCount = count
				mostCommonRazlog = razlog
			}
		}

		for korisnik, count := range korisnikCounts {
			if count > maxKorisnikCount {
				maxKorisnikCount = count
				mostCommonKorisnik = korisnik
			}
		}

		for location, count := range locationCounts {
			if count > maxLocationCount {
				maxLocationCount = count
				mostCommonLocation = location
			}
		}

		statistika := data.StatistikaNalogaZaPretres{
			ID:                         primitive.NewObjectID(),
			NajcesciRazlogPretresa:     mostCommonRazlog,
			NajcesciKorisnik:           mostCommonKorisnik,
			BrojIzdatihNalogaZaPretres: fmt.Sprintf("%d", totalNalozi),
			NajcescaLokacija:           mostCommonLocation,
		}

		// Encode the statistika object into JSON and send as response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statistika); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
func CreateStatistikaSaslusanja(lawCourtClient *client.LawCourtClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Hearings data
		hearings, err := lawCourtClient.FetchHearings(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch hearings: %v", err), http.StatusInternalServerError)
			return
		}

		razlogCounts := make(map[string]int)
		totalSaslusanja := len(hearings)
		totalTrajanje := 0

		for _, hearing := range hearings {
			razlogCounts[hearing.Title]++
			totalTrajanje += int(hearing.Duration.Minutes()) // Convert duration to minutes for total calculation
		}

		// Find the most common razlog
		var mostCommonRazlog string
		maxRazlogCount := 0

		for razlog, count := range razlogCounts {
			if count > maxRazlogCount {
				maxRazlogCount = count
				mostCommonRazlog = razlog
			}
		}

		// Calculate average duration
		var prosecnoTrajanjeSalusanja string
		if totalSaslusanja > 0 {
			prosecnoTrajanjeSalusanja = fmt.Sprintf("%.2f", float64(totalTrajanje)/float64(totalSaslusanja))
		} else {
			prosecnoTrajanjeSalusanja = "0"
		}

		statistika := data.StatistikaSaslusanja{
			ID:                        primitive.NewObjectID(),
			NajcesciRazlogSaslusanja:  mostCommonRazlog,
			ProsecnoTrajanjeSalusanja: prosecnoTrajanjeSalusanja,
			BrojGodisnjihSaslusanja:   fmt.Sprintf("%d", totalSaslusanja),
		}

		// Encode the statistika object into JSON and send as response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statistika); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}
func CreateStatistikaPravnogZahteva(lawCourtClient *client.LawCourtClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Fetch Legal Requests data
		legalRequests, err := lawCourtClient.FetchLegalRequests(r.Context())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch legal requests: %v", err), http.StatusInternalServerError)
			return
		}

		razlogCounts := make(map[string]int)
		korisnikCounts := make(map[string]int)
		totalZahteva := len(legalRequests)

		for _, zahtev := range legalRequests {
			razlogCounts[zahtev.Title]++
			korisnikCounts[zahtev.UserID.Hex()]++
		}

		// Find the most common razlog and korisnik
		var mostCommonRazlog string
		var mostCommonKorisnik string
		maxRazlogCount := 0
		maxKorisnikCount := 0

		for razlog, count := range razlogCounts {
			if count > maxRazlogCount {
				maxRazlogCount = count
				mostCommonRazlog = razlog
			}
		}

		for korisnik, count := range korisnikCounts {
			if count > maxKorisnikCount {
				maxKorisnikCount = count
				mostCommonKorisnik = korisnik
			}
		}

		statistika := data.StatistikaPravnogZahteva{
			ID:                   primitive.NewObjectID(),
			NajcesciRazlog:       mostCommonRazlog,
			NajcesciKorisnik:     mostCommonKorisnik,
			BrojGodisnjihZahteva: fmt.Sprintf("%d", totalZahteva),
		}

		// Encode the statistika object into JSON and send as response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(statistika); err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	}
}

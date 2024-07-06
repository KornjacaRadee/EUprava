package court_handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"law_court/client"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"law_court/data"
)

type CreateLegalEntityRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Licence     bool               `bson:"licence" json:"licence"`
	Category    string             `bson:"category" json:"category"`
	Description string             `bson:"description" json:"description"`
	IssueDate   time.Time          `bson:"issueDate" json:"issueDate"`
	DueToDate   time.Time          `bson:"dueToDate" json:"dueToDate"`
	JMBG        string             `json:"jmbg"`
}

type SearchWarrantRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	IssueDate   time.Time          `bson:"issueDate" json:"issueDate"`
	DueToDate   time.Time          `bson:"dueToDate" json:"dueToDate"`
	JMBG        string             `json:"jmbg"`
	Address     string             `bson:"address" json:"address"`
}
type Hearing struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title         string             `bson:"title" json:"title"`
	Description   string             `bson:"description" json:"description"`
	ScheduledAt   time.Time          `bson:"scheduledAt" json:"scheduledAt"`
	Duration      time.Duration      `bson:"duration" json:"duration"`
	LegalEntityID primitive.ObjectID `bson:"legalEntityId" json:"legalEntityId"`
}
type LegalRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	UserJMBG    string             `bson:"userJMBG" json:"userJMBG"`
	RequestDate time.Time          `bson:"requestDate" json:"requestDate"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
}

func CreateLegalRequest(dbClient *mongo.Client, saobracajClient *client.SaobracajnaPolicijaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var legalRequest data.LegalRequest
		if err := json.NewDecoder(r.Body).Decode(&legalRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//userID := r.Header.Get("UserID")
		//
		//userObjID, err := primitive.ObjectIDFromHex(userID)
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusBadRequest)
		//	return
		//}

		prekrsaji, err := saobracajClient.FetchPrekrsaji(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch prekrsaji", http.StatusInternalServerError)
			return
		}

		userJMBG := legalRequest.UserJMBG // Assuming UserID holds the JMBG
		userHasPrekrsaj := false
		for _, prekrsaj := range prekrsaji {
			if prekrsaj.Vozac == userJMBG {
				userHasPrekrsaj = true
				break
			}
		}

		if !userHasPrekrsaj {
			legalEntity := data.LegalEntity{
				Title:       "Potvrda o kaznjavanju",
				Description: "Covek nema na sebi ni jedan prekrsaj",
				IssueDate:   time.Now().Local(),
				DueToDate:   time.Now().Local().Add(time.Hour * 24 * 10),
				UserID:      legalRequest.UserID,
			}
			if err := data.CreateLegalEntity(dbClient, &legalEntity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(legalEntity)
			return
		} else {
			legalEntity := data.LegalEntity{
				Title:       "Potvrda o kaznjavanju",
				Description: "Covek ima prekrsaj i to nije lepo",
				IssueDate:   time.Now().Local(),
				DueToDate:   time.Now().Local().Add(time.Hour * 24 * 10),
				UserID:      legalRequest.UserID,
			}
			if err := data.CreateLegalEntity(dbClient, &legalEntity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(legalEntity)
			return
		}

		//legalRequest.UserID = userObjID
		//legalRequest.RequestDate = time.Now()
		//
		//if err := data.CreateLegalRequest(dbClient, &legalRequest); err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}
		//
		//w.WriteHeader(http.StatusCreated)
	}
}

func GetLegalRequestsByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		legalRequests, err := data.GetLegalRequestsByUserID(dbClient, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(legalRequests); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CreateLegalEntity(dbClient *mongo.Client, authClient *client.AuthClient, mupvozilaClient *client.MupvozilaClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateLegalEntityRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		_, err := getRoleFromToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		user, err := authClient.GetUserByJMBG(context.Background(), req.JMBG)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		legalEntity := data.LegalEntity{
			Title:       req.Title,
			Description: req.Description,
			IssueDate:   req.IssueDate,
			DueToDate:   req.DueToDate,
			UserID:      user.ID,
		}

		err = data.CreateLegalEntity(dbClient, &legalEntity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// FIKSIRAN USERID DOK KOLEGA NE NAMESTI SVOJ SERVIS!!!!!~
		if req.Licence {
			// Update the licence validity
			err = mupvozilaClient.UpdateLicenceValidity(context.Background(), "000000000000000000000000", req.Category)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(legalEntity)
	}
}

func GetLegalEntity(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		legalEntity, err := data.GetLegalEntity(dbClient, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if legalEntity == nil {
			http.Error(w, "Legal entity not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(legalEntity)
	}
}

func GetLegalEntitiesByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("a daj brate")
		log.Println("This is id: ", mux.Vars(r)["id"])

		userID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		log.Println("This is id in object: ", userID)
		if err != nil {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		legalEntities, err := data.GetLegalEntitiesByUserID(dbClient, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(legalEntities)
	}
}

func UpdateLegalEntity(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var legalEntity data.LegalEntity
		if err := json.NewDecoder(r.Body).Decode(&legalEntity); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err := data.UpdateLegalEntity(dbClient, id, &legalEntity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteLegalEntity(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err := data.DeleteLegalEntity(dbClient, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func CreateHouseSearchWarrant(dbClient *mongo.Client, authClient *client.AuthClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warrant data.HouseSearchWarrant
		var req CreateLegalEntityRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		_, err := getRoleFromToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := authClient.GetUserByJMBG(context.Background(), req.JMBG)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		warrant.Title = req.Title
		warrant.Description = req.Description
		warrant.IssueDate = req.IssueDate
		warrant.DueToDate = req.DueToDate
		warrant.UserID = user.ID

		err = data.CreateHouseSearchWarrant(dbClient, &warrant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetHouseSearchWarrant(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		warrant, err := data.GetHouseSearchWarrant(dbClient, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if warrant == nil {
			http.Error(w, "House search warrant not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(warrant)
	}
}

func GetHouseSearchWarrantsByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if userID.IsZero() {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		warrants, err := data.GetHouseSearchWarrantsByUserID(dbClient, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(warrants)
	}
}

func UpdateHouseSearchWarrant(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var warrant data.HouseSearchWarrant
		if err := json.NewDecoder(r.Body).Decode(&warrant); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err := data.UpdateHouseSearchWarrant(dbClient, id, &warrant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteHouseSearchWarrant(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		err := data.DeleteHouseSearchWarrant(dbClient, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func ScheduleHearing(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hearing data.Hearing
		if err := json.NewDecoder(r.Body).Decode(&hearing); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := data.ScheduleHearing(dbClient, &hearing)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func GetHearing(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["Id"])
		if id.IsZero() {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		hearing, err := data.GetHearing(dbClient, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if hearing == nil {
			http.Error(w, "Hearing not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(hearing)
	}
}

func GetHearingsByUserID(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("a daj brate")
		log.Println("This is id: ", mux.Vars(r)["id"])

		entityID, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
		log.Println("This is id in object: ", entityID)
		if entityID.IsZero() {
			http.Error(w, "Invalid userID", http.StatusBadRequest)
			return
		}

		hearings, err := data.GetHearingsByUserID(dbClient, entityID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(hearings)
	}
}

const jwtSecret = "g3HtH5KZNq3KcWglpIc3eOBHcrxChcY/7bTKG8a5cHtjn2GjTqUaMbxR3DBIr+44"

func getRoleFromToken(r *http.Request) (string, error) {
	// Extract the token from the Authorization header
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", errors.New("missing Authorization header")
	}

	// Remove 'Bearer ' prefix if present
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Provide the secret key used to sign the token
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("Invalid token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	// Extract user role from claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid token claims")
	}

	// Get user role
	role, ok := claims["roles"].(string)
	if !ok {
		return "", fmt.Errorf("User role not found in token claims")
	}

	return role, nil
}
func GetAllHouseSearchWarrants(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		houseSearchWarrants, err := data.GetAllHouseSearchWarrants(dbClient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(houseSearchWarrants)
	}
}
func GetAllHearings(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hearings, err := data.GetAllHearings(dbClient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(hearings)
	}
}
func GetAllLegalRequests(dbClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		legalRequests, err := data.GetAllLegalRequests(dbClient)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(legalRequests)
	}
}

package court_handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"law_court/client"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"law_court/data"
)

type CreateLegalEntityRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
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

func CreateLegalEntity(dbClient *mongo.Client, authClient *client.AuthClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var legalEntity data.LegalEntity
		var req CreateLegalEntityRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		user, err := authClient.GetUserByJMBG(context.Background(), req.JMBG)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		legalEntity.Title = req.Title
		legalEntity.Description = req.Description
		legalEntity.IssueDate = req.IssueDate
		legalEntity.DueToDate = req.DueToDate
		legalEntity.UserID = user.ID

		err = data.CreateLegalEntity(dbClient, &legalEntity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(legalEntity)
		w.WriteHeader(http.StatusCreated)
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

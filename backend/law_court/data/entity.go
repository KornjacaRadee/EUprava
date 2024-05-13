// data/court.go

package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LegalEntity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	IssueDate   time.Time          `bson:"issueDate" json:"issueDate"`
	DueToDate   time.Time          `bson:"dueToDate" json:"dueToDate"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
}

type HouseSearchWarrant struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	IssueDate   time.Time          `bson:"issueDate" json:"issueDate"`
	DueToDate   time.Time          `bson:"dueToDate" json:"dueToDate"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId"`
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

func CreateLegalEntity(dbClient *mongo.Client, legalEntity *LegalEntity) error {
	legalEntityCollection := dbClient.Database("authDB").Collection("legal_entities")

	_, err := legalEntityCollection.InsertOne(context.Background(), legalEntity)
	return err
}

func GetLegalEntitiesByUserID(dbClient *mongo.Client, userID primitive.ObjectID) ([]LegalEntity, error) {
	legalEntityCollection := dbClient.Database("authDB").Collection("legal_entities")

	var legalEntities []LegalEntity
	cursor, err := legalEntityCollection.Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var legalEntity LegalEntity
		if err := cursor.Decode(&legalEntity); err != nil {
			return nil, err
		}
		legalEntities = append(legalEntities, legalEntity)
	}

	return legalEntities, nil
}

func GetLegalEntity(dbClient *mongo.Client, id primitive.ObjectID) (*LegalEntity, error) {
	legalEntityCollection := dbClient.Database("authDB").Collection("legal_entities")

	var legalEntity LegalEntity
	err := legalEntityCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&legalEntity)
	if err != nil {
		return nil, err
	}

	return &legalEntity, nil
}

func UpdateLegalEntity(dbClient *mongo.Client, id primitive.ObjectID, legalEntity *LegalEntity) error {
	legalEntityCollection := dbClient.Database("authDB").Collection("legal_entities")

	update := bson.D{{"$set", legalEntity}}
	_, err := legalEntityCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	return err
}

func DeleteLegalEntity(dbClient *mongo.Client, id primitive.ObjectID) error {
	legalEntityCollection := dbClient.Database("authDB").Collection("legal_entities")

	_, err := legalEntityCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func CreateHouseSearchWarrant(dbClient *mongo.Client, warrant *HouseSearchWarrant) error {
	warrantCollection := dbClient.Database("authDB").Collection("house_search_warrants")

	_, err := warrantCollection.InsertOne(context.Background(), warrant)
	return err
}

func GetHouseSearchWarrantsByUserID(dbClient *mongo.Client, userID primitive.ObjectID) ([]HouseSearchWarrant, error) {
	warrantCollection := dbClient.Database("authDB").Collection("house_search_warrants")
	log.Println("Id is:", userID)
	var warrants []HouseSearchWarrant
	cursor, err := warrantCollection.Find(context.Background(), bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var warrant HouseSearchWarrant
		if err := cursor.Decode(&warrant); err != nil {
			return nil, err
		}
		warrants = append(warrants, warrant)
	}

	return warrants, nil
}

func GetHouseSearchWarrant(dbClient *mongo.Client, id primitive.ObjectID) (*HouseSearchWarrant, error) {
	warrantCollection := dbClient.Database("authDB").Collection("house_search_warrants")

	var warrant HouseSearchWarrant
	err := warrantCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&warrant)
	if err != nil {
		return nil, err
	}

	return &warrant, nil
}

func UpdateHouseSearchWarrant(dbClient *mongo.Client, id primitive.ObjectID, warrant *HouseSearchWarrant) error {
	warrantCollection := dbClient.Database("authDB").Collection("house_search_warrants")

	update := bson.D{{"$set", warrant}}
	_, err := warrantCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	return err
}

func DeleteHouseSearchWarrant(dbClient *mongo.Client, id primitive.ObjectID) error {
	warrantCollection := dbClient.Database("authDB").Collection("house_search_warrants")

	_, err := warrantCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func ScheduleHearing(dbClient *mongo.Client, hearing *Hearing) error {
	hearingCollection := dbClient.Database("authDB").Collection("hearings")

	_, err := hearingCollection.InsertOne(context.Background(), hearing)
	return err
}

func GetHearingsByUserID(dbClient *mongo.Client, entityID primitive.ObjectID) ([]Hearing, error) {
	hearingCollection := dbClient.Database("authDB").Collection("hearings")

	var hearings []Hearing
	cursor, err := hearingCollection.Find(context.Background(), bson.M{"legalEntityId": entityID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var hearing Hearing
		if err := cursor.Decode(&hearing); err != nil {
			return nil, err
		}
		hearings = append(hearings, hearing)
	}

	return hearings, nil
}

func GetHearing(dbClient *mongo.Client, id primitive.ObjectID) (*Hearing, error) {
	hearingCollection := dbClient.Database("authDB").Collection("hearings")

	var hearing Hearing
	err := hearingCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&hearing)
	if err != nil {
		return nil, err
	}

	return &hearing, nil
}

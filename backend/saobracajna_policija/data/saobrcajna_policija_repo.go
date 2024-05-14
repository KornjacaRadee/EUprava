package data

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type SaobracajnaPolicijaRepo struct {
	cli    *mongo.Client
	logger *log.Logger
	client *http.Client
}

var (
	dbHost = os.Getenv("SAOBRACAJNA_POLICIJA_DB_HOST")
	dbPort = os.Getenv("SAOBRACAJNA_POLICIJA_DB_PORT")
	dbName = os.Getenv("SAOBRACAJNA_POLICIJA_DB_NAME")
)

func NewSaobracajnaPolicijaRepo(ctx context.Context, logger *log.Logger) *SaobracajnaPolicijaRepo {
	dburi := fmt.Sprintf("mongodb://%s:%s/", dbHost, dbPort)

	client, err := mongo.NewClient(options.Client().ApplyURI(dburi))
	if err != nil {
		return nil
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			MaxConnsPerHost:     10,
		},
	}

	// Return repository with logger and DB client
	return &SaobracajnaPolicijaRepo{
		cli:    client,
		logger: logger,
		client: httpClient,
	}
}

func (pr *SaobracajnaPolicijaRepo) CreateNesreca(ctx context.Context, nesreca *Nesreca) error {
	collection := pr.cli.Database(dbName).Collection("nesrece")
	_, err := collection.InsertOne(ctx, nesreca)
	if err != nil {
		return err
	}
	return nil
}

func (pr *SaobracajnaPolicijaRepo) CreatePrekrsaj(ctx context.Context, prekrsaj *Prekrsaj) error {
	collection := pr.cli.Database(dbName).Collection("prekrsaji")
	_, err := collection.InsertOne(ctx, prekrsaj)
	if err != nil {
		return err
	}
	return nil
}

func (pr *SaobracajnaPolicijaRepo) GetNesrece(ctx context.Context) ([]Nesreca, error) {
	collection := pr.cli.Database(dbName).Collection("nesrece")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var nesrece []Nesreca
	if err := cursor.All(ctx, &nesrece); err != nil {
		return nil, err
	}

	return nesrece, nil
}

func (pr *SaobracajnaPolicijaRepo) GetPrekrsaji(ctx context.Context) ([]Prekrsaj, error) {
	collection := pr.cli.Database(dbName).Collection("prekrsaji")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var prekrsaji []Prekrsaj
	if err := cursor.All(ctx, &prekrsaji); err != nil {
		return nil, err
	}

	return prekrsaji, nil
}

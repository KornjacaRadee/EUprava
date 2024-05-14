package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"saobracajna_policija/data"
	"saobracajna_policija/handlers"

	"time"
)

func main() {
	// Inicijalizacija logera
	logger := log.New(os.Stdout, "[saobracajna-policija] ", log.LstdFlags)

	// Inicijalizacija MongoDB klijenta
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logger.Fatal("Failed to create MongoDB client: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB: ", err)
	}
	defer mongoClient.Disconnect(ctx)

	// Inicijalizacija repositorijuma
	repo := data.NewSaobracajnaPolicijaRepo(ctx, logger)

	// Inicijalizacija handlera
	handler := handler.NewSaobracajnaPolicijaHandler(repo, logger)

	// Inicijalizacija rutera
	router := mux.NewRouter()

	// Rute za rukovanje nesrećama
	router.HandleFunc("/nesreca/new", handler.CreateNesreca).Methods(http.MethodPost)

	// Rute za rukovanje prekršajima
	router.HandleFunc("/prekrsaj/new", handler.CreatePrekrsaj).Methods(http.MethodPost)

	router.HandleFunc("/nesreca", handler.GetNesrece).Methods(http.MethodGet)

	router.HandleFunc("/prekrsaji", handler.GetPrekrsaji).Methods(http.MethodGet)

	router.HandleFunc("/nesreca/{id}", handler.DeleteNesreca).Methods(http.MethodDelete)
	router.HandleFunc("/prekrsaj/{id}", handler.DeletePrekrsaj).Methods(http.MethodDelete)

	// Kreiranje HTTP servera
	server := &http.Server{
		Addr:         ":8084",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Pokretanje HTTP servera u zasebnoj rutini
	go func() {
		logger.Println("Starting server on port 8080")
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Očekivanje signala za zaustavljanje servera
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received termination signal: ", sig)

	// Zaustavljanje servera sa timeout-om
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(timeoutCtx)
	logger.Println("Server stopped")
}

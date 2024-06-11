package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"law_court/client"
	"law_court/court_handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	dburi := os.Getenv("MONGO_DB_URI")
	clientOptions := options.Client().ApplyURI(dburi)

	dbClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = dbClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8082"
	}
	auth := client.NewAuthClient(&http.Client{}, "http://auth_service:8082")
	r := mux.NewRouter()
	r.HandleFunc("/get_user", func(w http.ResponseWriter, r *http.Request) {

		userID := r.URL.Query().Get("id")

		user, err := auth.GetUserByID(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		userJSON, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(userJSON)
		if err != nil {
			log.Println("Error writing response:", err)
		}
	})
	r.HandleFunc("/legal_entities", court_handlers.CreateLegalEntity(dbClient, auth)).Methods("POST")
	r.HandleFunc("/legal_entities/{id}", court_handlers.GetLegalEntity(dbClient)).Methods("GET")
	r.HandleFunc("/legal_entities/user/{id}", court_handlers.GetLegalEntitiesByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/legal_entities/{id}", court_handlers.UpdateLegalEntity(dbClient)).Methods("PUT")
	r.HandleFunc("/legal_entities/{id}", court_handlers.DeleteLegalEntity(dbClient)).Methods("DELETE")
	r.HandleFunc("/house_search_warrants", court_handlers.CreateHouseSearchWarrant(dbClient)).Methods("POST")
	r.HandleFunc("/house_search_warrants/{id}", court_handlers.GetHouseSearchWarrant(dbClient)).Methods("GET")
	r.HandleFunc("/house_search_warrants/user/{id}", court_handlers.GetHouseSearchWarrantsByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/house_search_warrants/{id}", court_handlers.UpdateHouseSearchWarrant(dbClient)).Methods("PUT")
	r.HandleFunc("/house_search_warrants/{id}", court_handlers.DeleteHouseSearchWarrant(dbClient)).Methods("DELETE")
	r.HandleFunc("/hearings", court_handlers.ScheduleHearing(dbClient)).Methods("POST")
	r.HandleFunc("/hearings/entity/{id}", court_handlers.GetHearingsByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/hearings/{id}", court_handlers.GetHearing(dbClient)).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"}) // Update with your Angular app's origin

	// Apply CORS middleware
	handlerWithCORS := handlers.CORS(headers, methods, origins)(r)

	http.Handle("/", handlerWithCORS)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      handlerWithCORS,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	<-sigCh // Wait for a signal

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	mupvozila := client.NewMupvozilaClient(&http.Client{}, "http://mupvozila_service:8081")
	saobracajnaPolicijaClient := client.NewSaobracajnaPolicijaClient("http://saobracajna_policija:8084")

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

	r.HandleFunc("/legal_entities", court_handlers.CreateLegalEntity(dbClient, auth, mupvozila)).Methods("POST")
	r.HandleFunc("/legal_entities/{id}", court_handlers.GetLegalEntity(dbClient)).Methods("GET")
	r.HandleFunc("/legal_entities/user/{id}", court_handlers.GetLegalEntitiesByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/legal_entities/{id}", court_handlers.UpdateLegalEntity(dbClient)).Methods("PUT")
	r.HandleFunc("/legal_entities/{id}", court_handlers.DeleteLegalEntity(dbClient)).Methods("DELETE")
	r.HandleFunc("/house_search_warrants", court_handlers.CreateHouseSearchWarrant(dbClient, auth)).Methods("POST")
	r.HandleFunc("/house_search_warrants/{id}", court_handlers.GetHouseSearchWarrant(dbClient)).Methods("GET")
	r.HandleFunc("/house_search_warrants/user/{id}", court_handlers.GetHouseSearchWarrantsByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/house_search_warrants/{id}", court_handlers.UpdateHouseSearchWarrant(dbClient)).Methods("PUT")
	r.HandleFunc("/house_search_warrants/{id}", court_handlers.DeleteHouseSearchWarrant(dbClient)).Methods("DELETE")
	r.HandleFunc("/hearings", court_handlers.ScheduleHearing(dbClient)).Methods("POST")
	r.HandleFunc("/hearings/entity/{id}", court_handlers.GetHearingsByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/hearings/{id}", court_handlers.GetHearing(dbClient)).Methods("GET")
	r.HandleFunc("/legal_requests", court_handlers.CreateLegalRequest(dbClient, saobracajnaPolicijaClient)).Methods("POST")
	r.HandleFunc("/legal_requests/user/{id}", court_handlers.GetLegalRequestsByUserID(dbClient)).Methods("GET")
	r.HandleFunc("/house_search_warrants/all", court_handlers.GetAllHouseSearchWarrants(dbClient)).Methods("GET")
	r.HandleFunc("/legal_requests/all", court_handlers.GetAllLegalRequests(dbClient)).Methods("GET")
	r.HandleFunc("/hearings/all", court_handlers.GetAllHearings(dbClient)).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{
		Handler:      handlers.CORS(headers, methods, origins)(r),
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("ListenAndServe error: %v\n", err)
		}
	}()

	log.Printf("Server running on port %s\n", port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Println("shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	dbClient.Disconnect(ctx)
}

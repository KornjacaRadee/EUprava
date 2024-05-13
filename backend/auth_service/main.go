package main

import (
	"auth_service/authHandlers"
	"auth_service/config"
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	logger := config.NewLogger("./logging/log.log")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8082"
	}

	r := mux.NewRouter()
	r.HandleFunc("/register", authHandlers.HandleRegister(dbClient)).Methods("POST")
	r.HandleFunc("/login", authHandlers.HandleLogin(dbClient)).Methods("POST")
	r.HandleFunc("/users", authHandlers.HandleGetAllUsers(dbClient)).Methods("GET")
	//r.HandleFunc("/user", authHandlers.HandleDeleteUser(dbClient, reservation, accommodation, profile)).Methods("DELETE")
	//r.HandleFunc("/users/{id}", authHandlers.HandleGetUserByID(dbClient)).Methods("GET")
	// change user passwrod
	//r.HandleFunc("/change-password", authHandlers.HandleChangePassword(dbClient)).Methods("POST")

	// Enable CORS
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"}) // Update with your Angular app's origin

	// Apply CORS middleware
	handlerWithCORS := handlers.CORS(headers, methods, origins)(r)

	http.Handle("/", handlerWithCORS)

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      handlerWithCORS,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Panicf("Panic on auth-service during listening")
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shut down gracefully
	if server.Shutdown(context.TODO()) != nil {
		logger.Fatalf("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}

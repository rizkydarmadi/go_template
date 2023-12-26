package routes

import (
	"fmt"
	"go-crud-api/middleware"
	"go-crud-api/src/api/users"
	"go-crud-api/src/config"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//create router

func Routes() {

	db := config.ConnectDBGen()

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/users", users.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", users.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", users.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", users.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", users.DeleteUser(db)).Methods("DELETE")

	//start server
	router.Use(middleware.LoggingMiddleware)
	// Start the server
	numb_port := os.Getenv("PORT")
	port := fmt.Sprintf(":%s", numb_port)
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, router)
	log.Fatal(http.ListenAndServe(port, middleware.JsonContentTypeMiddleware(router)))

}

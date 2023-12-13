package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"go-crud-api/api/users"
	"go-crud-api/middleware"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// connect database
	db, err := sql.Open("postgres", "postgres://postgres:12qwaszx@localhost:5432/demo_go?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//create router
	router := mux.NewRouter()
	router.HandleFunc("/users", users.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", users.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", users.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", users.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("users/{id}", users.DeleteUser(db)).Methods("DELETE")

	//start server
	router.Use(middleware.LoggingMiddleware)
	// Start the server
	port := ":8000"
	fmt.Printf("Server listening on port %s...\n", port)
	http.ListenAndServe(port, router)
	log.Fatal(http.ListenAndServe(port, middleware.JsonContentTypeMiddleware(router)))
}

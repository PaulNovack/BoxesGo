package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var db *sql.DB

func main() {

	db := dbconnect()
	db.Ping()
	r := mux.NewRouter()
	// Handle Boxes endpoints
	r.HandleFunc("/boxes", postBox).Methods("POST")
	r.HandleFunc("/boxes/{id}", putBox).Methods("PUT")
	r.HandleFunc("/boxes", getAllBoxes).Methods("GET")
	// Handle Items endpoints
	r.HandleFunc("/items/{box_id}", postItem).Methods("POST")
	r.HandleFunc("/items/{box_id}", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", putItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	//Handle User endpoints
	r.HandleFunc("/user/{id}", getUser).Methods("GET")
	r.HandleFunc("/user/", createUser).Methods("POST")
	// Endpoint react front end can check to see backend is up
	r.HandleFunc("/up", func(w http.ResponseWriter, r *http.Request) {
		// Just check server is up
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	user, err := dbgetUserId(db, "paulnovack", "paulnovack")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", user)
	log.Fatal(srv.ListenAndServe())

}

package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var db *sql.DB
var ZMQAlert bool

func main() {
	ZMQAlert = false
	db, _ = sql.Open("mysql", "boxes:boxes@(127.0.0.1:3306)/boxes?parseTime=true")
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	r := mux.NewRouter()
	// Handle Boxes endpoints
	r.HandleFunc("/boxes", postBox).Methods("POST")
	r.HandleFunc("/boxes", putBox).Methods("PUT")
	r.HandleFunc("/boxes", getAllBoxes).Methods("GET")
	r.HandleFunc("/boxes/{box_id}", deleteBox).Methods("DELETE")
	// Handle Items endpoints
	r.HandleFunc("/items", postItem).Methods("POST")
	r.HandleFunc("/items", putItem).Methods("PUT")
	r.HandleFunc("/items/{box_id}", getItems).Methods("GET")
	r.HandleFunc("/items/{item_id}", deleteItem).Methods("DELETE")
	//Handle User endpoints
	r.HandleFunc("/login", getUser).Methods("GET")
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

	log.Fatal(srv.ListenAndServe())

}

package main

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
)

func dbconnect() *sql.DB {
	db, err := sql.Open("mysql", "boxes:boxes@(127.0.0.1:3306)/boxes?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func createSessionKey() string {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		// handle error here
	}
	s := fmt.Sprintf("%x", key)
	//fmt.Println(s)
	return s
}

func dbgetUserId(db *sql.DB, username string, password string) (User, error) {
	query := `select id,created_at from users where username = ? and password = ?`
	res, err := db.Query(query, username, password)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	for res.Next() {
		err := res.Scan(&user.Id, &user.Createdate)
		if err != nil {
			log.Fatal(err)
		}
	}
	sessionKey := createSessionKey()
	query = "update users set authkey = ? where id = ?"
	_, err = db.Query(query, user.Id, sessionKey)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func dbgetBoxes(db *sql.DB, user_id int) ([]Box, error) {
	query := `SELECT id, user_id, name, picture, created_at from boxes where user_id = ?`
	res, err := db.Query(query, user_id)
	if err != nil {
		log.Fatal(err)
	}
	var boxes []Box
	for res.Next() {
		var box Box
		err := res.Scan(&box.Id, &box.Userid, &box.Name, &box.Picture, &box.Createdate)
		if err != nil {
			log.Fatal(err)
		}
		boxes = append(boxes, box)
	}
	return boxes, err
}

func dbgetItems(db *sql.DB, box_id int) ([]Item, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	res, err := db.Query(query, box_id)
	if err != nil {
		log.Fatal(err)
	}
	var items []Item
	for res.Next() {
		var item Item
		err := res.Scan(&item.Id, &item.Userid, &item.Boxid, &item.Name, &item.Picture, &item.Createdate)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	return items, err
}

func dbputItem(db *sql.DB, item Item) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, item.Id)
	return 0, err
}
func dbputBox(db *sql.DB, box Box) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, box.Id)
	return 0, err
}

func dbputUser(db *sql.DB, user User) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, user.Id)
	return 0, err
}

func dbpostItem(db *sql.DB, item Item) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, item.Id)
	return 0, err
}
func dbpostBox(db *sql.DB, box Box) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, box.Id)
	return 0, err
}
func dbdeleteItem(db *sql.DB, user_id int, item_id int) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, item_id)
	return 0, err
}
func dbdeleteBox(db *sql.DB, user_id int, box_id int) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, box_id)
	return 0, err
}

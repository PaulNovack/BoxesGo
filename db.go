package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "boxes:boxes@(127.0.0.1:3306)/boxes?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func getUserId(db *sql.DB, username string, password string) (User, error) {
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
	return user, err
}

func getBoxes(db *sql.DB, user_id int) ([]Box, error) {
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

func getItems(db *sql.DB, box_id int) ([]Item, error) {
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

func putItem(db *sql.DB, item Item) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, item.Id)
	return 0, err
}
func putBox(db *sql.DB, box Box) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, box.Id)
	return 0, err
}

func putUser(db *sql.DB, user User) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, user.Id)
	return 0, err
}

func postItem(db *sql.DB, item Item) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, item.Id)
	return 0, err
}
func postBox(db *sql.DB, box Box) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, box.Id)
	return 0, err
}

package main

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func createSessionKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		// handle error here
	}
	s := fmt.Sprintf("%x", key)
	//fmt.Println(s)
	return s
}
func dbGetUserFromAuthKey(AuthKey string) (id int) {
	query := `select id from users where authKey = ?`
	res, err := db.Query(query, AuthKey)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	for res.Next() {
		err := res.Scan(&user.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	res.Close()
	return user.Id
}
func dbcreateUser(username string, password string) (user User, err error) {
	query := `select id,authkey,created_at from users where username = ?`
	res, err := db.Query(query, username)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		err := res.Scan(&user.Id, &user.Authkey, &user.Createdate)
		if err != nil {
			log.Fatal(err)
		}
	}
	res.Close()
	if user.Id != 0 {
		err = fmt.Errorf(`the user exists password is incorrect`)
	} else {
		authKey := createSessionKey()
		insertquery := `insert into users (username,password,authkey) values(?,?,?)`
		_, err := db.Exec(insertquery, username, password, authKey)
		if err != nil {
			log.Fatal(err)
		}
		query = `select last_insert_id()`
		res, err := db.Query(query)
		for res.Next() {
			err := res.Scan(&user.Id)
			if err != nil {
				log.Fatal(err)
			}
		}
		user.Username = username
		user.Authkey = authKey
		user.Password = password
		res.Close()
	}
	return user, err
}

func dbgetUser(username string, password string) (User, error) {
	query := `select id,authkey,created_at from users where username = ? and password = ?`
	res, err := db.Query(query, username, password)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	for res.Next() {
		err := res.Scan(&user.Id, &user.Authkey, &user.Createdate)
		if err != nil {
			log.Fatal(err)
		}
	}
	res.Close()
	if user.Id == 0 {
		dbcreateUser(username, password)
	}

	res.Close()
	sessionKey := createSessionKey()
	query = "update users set authkey = ? where id = ?"
	_, err = db.Exec(query, sessionKey, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	query = `select id,username,authkey,created_at from users where username = ? and password = ?`
	res, err = db.Query(query, username, password)
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		err := res.Scan(&user.Id, &user.Username, &user.Authkey, &user.Createdate)
		if err != nil {
			log.Fatal(err)
		}
	}
	res.Close()
	return user, err
}

func dbgetBoxes(db *sql.DB, user_id int) ([]Box, error) {
	query := `SELECT id, user_id, name,weight, picture, created_at from boxes where user_id = ?`
	res, err := db.Query(query, user_id)
	if err != nil {
		log.Fatal(err)
	}
	var boxes []Box
	for res.Next() {
		var box Box
		err := res.Scan(&box.Id, &box.Userid, &box.Name, &box.Weight, &box.Picture, &box.Createdate)
		if err != nil {
			log.Fatal(err)
		}
		boxes = append(boxes, box)
	}
	return boxes, err
}

func dbgetItems(db *sql.DB, user_id int, box_id string) ([]Item, error) {
	query := `SELECT id, user_id, box_id, name,quantity, picture, created_at 
				from items where user_id = ? and  box_id = ?  `
	res, err := db.Query(query, user_id, box_id)
	if err != nil {
		log.Fatal(err)
	}
	var items []Item
	for res.Next() {
		var item Item
		err := res.Scan(&item.Id, &item.Userid, &item.Boxid, &item.Name,
			&item.Quantity, &item.Picture, &item.Createdate)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	return items, err
}
func dbisMyItem(db *sql.DB, userId int, item Item) (Item, error) {
	query := `SELECT id
				from items where user_id = ? and  id = ?  `
	res, err := db.Query(query, userId, item.Id)
	if err != nil {
		log.Fatal(err)
	}
	// must be able to query the box id
	item.Id = 0
	for res.Next() {
		err := res.Scan(&item.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return item, err
}
func dbputItem(db *sql.DB, userId int, item Item) (Item, error) {
	item, err := dbisMyItem(db, userId, item)
	if item.Id == 0 {
		err := errors.New("can not update item its not yours")
		return item, err
	}
	var query string
	var box Box
	box.Id = item.Boxid
	box, err = dbisMyBox(db, userId, box)
	if box.Id != 0 {
		// Own this box not going to return error right now just update other fields
		// if box.Id is 0 do not own box going to skip moving it to new box for now
		if item.Boxid != 0 {
			query = "update items set box_id = ? where id = ?"
			_, err = db.Exec(query, item.Boxid, item.Id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	if item.Name != "" {
		query = "update items set name = ? where id = ?"
		_, err = db.Exec(query, item.Name, item.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	if item.Picture != "" {
		query = "update items set picture = ? where id = ?"
		_, err = db.Exec(query, item.Picture, item.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	if item.Quantity != 0 {
		query = "update items set quantity = ? where id = ?"
		_, err = db.Exec(query, item.Quantity, item.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	query = `SELECT id, user_id, box_id,name,quantity, picture, created_at from items where id = ?`
	res, err := db.Query(query, item.Id)

	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {

		err := res.Scan(&item.Id, &item.Userid, &item.Boxid, &item.Name, &item.Quantity, &item.Picture, &item.Createdate)
		if err != nil {
			log.Fatal(err)
		}
	}
	return item, err
}
func dbisMyBox(db *sql.DB, userId int, box Box) (Box, error) {
	query := `SELECT id
				from boxes where user_id = ? and  id = ?  `
	res, err := db.Query(query, userId, box.Id)
	if err != nil {
		log.Fatal(err)
	}
	// must be able to query the box id
	box.Id = 0
	for res.Next() {
		err := res.Scan(&box.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return box, err
}
func dbputBox(db *sql.DB, userId int, box Box) (Box, error) {
	box, err := dbisMyBox(db, userId, box)
	if box.Id == 0 {
		err := errors.New("can not update box its not yours")
		return box, err
	}
	var query string
	if box.Name != "" {
		query = "update boxes set name = ? where id = ?"
		_, err = db.Exec(query, box.Name, box.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	if box.Picture != "" {
		query = "update boxes set picture = ? where id = ?"
		_, err = db.Exec(query, box.Picture, box.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	if box.Weight != 0 {
		query = "update boxes set weight = ? where id = ?"
		_, err = db.Exec(query, box.Weight, box.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	query = `SELECT id, user_id, name,weight, picture, created_at from boxes where id = ?`
	res, err := db.Query(query, box.Id)

	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {

		err := res.Scan(&box.Id, &box.Userid, &box.Name, &box.Weight, &box.Picture, &box.Createdate)
		if err != nil {
			log.Fatal(err)
		}
	}
	return box, err
}

func dbputUser(db *sql.DB, user User) (int, error) {
	query := `SELECT id, user_id, box_id, quantity,name, picture, created_at from items where box_id = ?`
	_, err := db.Query(query, user.Id)
	return 0, err
}

func dbpostItem(db *sql.DB, userId int, item Item) (newItem Item, err error) {
	var box Box
	box.Id = item.Boxid
	box, err = dbisMyBox(db, userId, box)
	if box.Id != 0 {
		item.Userid = userId
		query := `insert into items (user_id,box_id,name,quantity,picture) values (?,?,?,?,?)`
		_, err := db.Exec(query, item.Userid, item.Boxid, item.Name, item.Quantity, item.Picture)
		if err != nil {
			log.Fatal(err)
		}
		query = `select last_insert_id()`
		res, err := db.Query(query)
		for res.Next() {
			err := res.Scan(&item.Id)
			if err != nil {
				log.Fatal(err)
			}
		}
		return item, err
	} else {
		var itemfail Item
		return itemfail, err
	}

	return item, err
}
func dbpostBox(db *sql.DB, userId int, box Box) (Box, error) {
	query := `insert into boxes (user_id,name,weight,picture) values (?,?,?,?)`
	_, err := db.Exec(query, box.Userid, box.Name, box.Weight, box.Picture)
	if err != nil {
		log.Fatal(err)
	}
	query = `select last_insert_id()`
	res, err := db.Query(query)
	for res.Next() {
		err := res.Scan(&box.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return box, err
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

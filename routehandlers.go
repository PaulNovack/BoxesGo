package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func checkSession(w http.ResponseWriter, r *http.Request) (userid int) {
	w.Header().Set("Content-Type", "application/json")
	authKey, _ := r.Cookie("authToken")
	if authKey == nil {
		w.Write([]byte(`{"error":"you must first login"}`))
		return 0
	}
	var user User
	user.Id = dbGetUserFromAuthKey(authKey.Value)
	if user.Id == 0 {
		w.Write([]byte(`{"error":"you must first login invalid cookie"}`))
		return 0
	}
	// refresh cookie allow 30 minutes more with each request
	cookie := &http.Cookie{
		Name:   "authToken",
		Value:  authKey.Value,
		MaxAge: 1800,
	}
	http.SetCookie(w, cookie)
	return user.Id
}

func postBox(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var box Box
	err = json.Unmarshal([]byte(body), &box)

	box, err = dbpostBox(db, userId, box)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(box)
}
func putBox(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var box Box
	err = json.Unmarshal([]byte(body), &box)

	box, err = dbputBox(db, userId, box)
	if err != nil {
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	} else {
		json.NewEncoder(w).Encode(box)
	}

}
func deleteBox(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func getAllBoxes(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}

	boxes, err := dbgetBoxes(db, userId)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(boxes)
}
func postItem(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var item Item
	err = json.Unmarshal([]byte(body), &item)

	item, err = dbpostItem(db, userId, item)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(item)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	boxId := params["box_id"]
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}

	items, err := dbgetItems(db, userId, boxId)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(items)
}

func putItem(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var item Item
	err = json.Unmarshal([]byte(body), &item)
	if err != nil {
		log.Fatal(err)
	}
	item, err = dbputItem(db, userId, item)
	if err != nil {
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
	} else {
		json.NewEncoder(w).Encode(item)
	}
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	userId := checkSession(w, r)
	if userId == 0 {
		return
	}
	if false {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"error": "you do not have permission to delete this item"})
	}

}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	if `` == username || `` == password {
		w.Write([]byte(`{"error":"username and password are required"}`))
		return
	}
	user, err := dbgetUser(username, password)
	if err != nil {
		log.Fatal(err)
	}
	sessionKey := user.Authkey
	cookie := &http.Cookie{
		Name:   "authToken",
		Value:  sessionKey,
		MaxAge: 1800,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode(user)
}

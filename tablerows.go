package main

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Authkey    string    `json:"authkey"`
	Createdate time.Time `json:"created_at"`
}

type Box struct {
	Id         int       `json:"id"`
	Userid     int       `json:"user_id"`
	BoxId      int       `json:"box_id"`
	Name       string    `json:"name"`
	Picture    string    `json:"picture"`
	Createdate time.Time `json:"created_at"`
}

type Item struct {
	Id         int       `json:"id"`
	Boxid      int       `json:"box_id"`
	Userid     int       `json:"user_id"`
	Name       string    `json:"name"`
	Quantity   int       `json:"quantity"`
	Picture    string    `json:"picture"`
	Createdate time.Time `json:"created_at"`
}

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Authkey    string
	Createdate time.Time
}

type Box struct {
	Id         int
	Userid     int
	Name       string
	Picture    string
	Createdate time.Time
}

type Item struct {
	Id         int
	Boxid      int
	Userid     int
	Name       string
	Quantity   int
	Picture    string
	Createdate time.Time
}

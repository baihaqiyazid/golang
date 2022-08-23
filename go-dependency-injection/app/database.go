package app

import (
	"database/sql"
	"go-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go-restful")
	helper.Panic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(20 * time.Minute)

	return db
}
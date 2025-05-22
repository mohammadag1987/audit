package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // ثبت درایور
)

func openSQLiteB(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	//defer db.Close() // مطمئن شوید که اتصال بسته می‌شود

	// 3. تست اتصال (Ping)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Successfully connected to DB!")

	return db, nil
}

func (app *application) ConnectToSQLite() (*sql.DB, error) {

	connection, err := openSQLiteB(app.SQLiteDSN)
	if err != nil {
		return nil, err
	}
	connection.SetMaxOpenConns(1)

	log.Println("Connected to DB!")
	return connection, nil
}

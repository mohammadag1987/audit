package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // وارد کردن درایور SQL Server
)

func openSQLServerDB(connectionString string) (*sql.DB, error) {

	db, err := sql.Open("sqlserver", connectionString)
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

func (app *application) ConnectToSQLServer() (*sql.DB, error) {

	connection, err := openSQLServerDB(app.SQLServeDSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB!")
	return connection, nil
}
func (app *application) ConnectToMachineSQLServer() (*sql.DB, error) {

	connection, err := openSQLServerDB(app.SQLServeMachineDSN)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB!")
	return connection, nil
}

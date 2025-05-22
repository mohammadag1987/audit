package main

import (
	"audit/internal/repository"
	"audit/internal/repository/dbrepo"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	SQLServeDSN                  string
	SQLServeMachineDSN           string
	SQLiteDSN                    string
	Domain                       string
	SQLServerDBConnection        repository.SQLServerDatabaseRepo
	SQLServerMachineDBConnection repository.SQLServerDatabaseRepo
	SQLiteDBConnection           repository.SQLiteDatabaseRepo
	//auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	APIKey       string
}

func main() {
	// set application config
	var app application
	app.SQLServeDSN = "sqlserver://sa:1@localhost:14330?database=CloudMonitoring&encrypt=disable"
	app.SQLiteDSN = "./SGTroubleshooterDB.s3db?cache=shared&mode=memory"
	app.JWTSecret = "verysecret"
	app.JWTIssuer = "example.com"
	app.JWTAudience = "example.com"
	app.CookieDomain = "" //"localhost"
	app.Domain = "example.com"
	app.APIKey = "e8bd48681c008942dd85a1596b2e4692"
	//flag.StringVar(&app.DSN, "dsn", "host=localhost port=45600 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	//flag.Parse()

	// connect to the database
	con, err := app.ConnectToSQLServer()
	if err != nil {
		log.Fatal(err)
	}
	app.SQLServerDBConnection = &dbrepo.SQLServerDBRepo{DB: con}
	defer app.SQLServerDBConnection.Connection().Close()

	con2, err2 := app.ConnectToSQLite()
	if err2 != nil {
		log.Fatal(err2)
	}
	app.SQLiteDBConnection = &dbrepo.SQLiteDBRepo{DB: con2}
	defer app.SQLServerDBConnection.Connection().Close()
	/*
		app.auth = Auth{
			Issuer:        app.JWTIssuer,
			Audience:      app.JWTAudience,
			Secret:        app.JWTSecret,
			TokenExpiry:   time.Minute * 15,
			RefreshExpiry: time.Hour * 24,
			CookiePath:    "/",
			CookieName:    "refresh_token",
			CookieDomain:  app.CookieDomain,
		}
	*/
	// start a web server
	log.Printf("Starting application on port: %v", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}

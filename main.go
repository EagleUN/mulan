package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jschavesr/mulan/app"
	"github.com/jschavesr/mulan/db"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}

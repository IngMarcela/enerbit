package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/handlers"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	handlers.ConfigureHandlers(router, db)

	log.Println("Servidor corriendo")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

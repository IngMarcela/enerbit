package handlers

import (
	"github.com/gorilla/mux"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/handlers/customer"
	"testUser/testEnerBit/handlers/workOrder"
)

func ConfigureHandlers(router *mux.Router, db *database.DB) {
	router.HandleFunc("/customers/active/", customer.ActiveHandler(db)).Methods("GET")
	router.HandleFunc("/customers/store", customer.StoreHandler(db)).Methods("POST")
	router.HandleFunc("/workOrder/store", workOrder.StoreHandler(db)).Methods("POST")
	router.HandleFunc("/workOrder/date/", workOrder.DateHandler(db)).Methods("GET")
	router.HandleFunc("/workOrder/customer/{id}", workOrder.CustomerOrderHandler(db)).Methods("GET")
	router.HandleFunc("/workOrder/customerOrder/{id}", workOrder.CustomerOrderTogetherHandler(db)).Methods("GET")
}

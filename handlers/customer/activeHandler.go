package customer

import (
	"encoding/json"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/models"
)

func FindActiveCustomers(db *database.DB) ([]models.Customer, error) {
	var customers []models.Customer
	if err := db.Find(&customers, "is_active = ?", true).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ActiveHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customers, err := FindActiveCustomers(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseJSON(w, customers)
	}
}

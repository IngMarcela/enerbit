package customer

import (
	"encoding/json"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/models"
)

func StoreHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer
		err := json.NewDecoder(r.Body).Decode(&customer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := db.Exec("INSERT INTO customers (first_name, last_name, address, end_date, is_active, created_at) VALUES (?, ?, ?, ?, ?, ?)", customer.FirstName, customer.LastName, customer.Address, customer.EndDate, customer.IsActive, customer.CreatedAt)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

package workOrder

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/models"
)

func CustomerOrderHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		customerID := vars["id"]

		var workOrder []models.WorkOrder
		result := db.Find(&workOrder, "customer_id = uuid_to_bin(?)", customerID)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(workOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

}

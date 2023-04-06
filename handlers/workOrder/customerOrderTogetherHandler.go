package workOrder

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/models"
)

type OrderInfo struct {
	WorkOrder models.WorkOrder
	Customer  models.Customer
}

func CustomerOrderTogetherHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		customerID := vars["id"]

		var orderInfo OrderInfo
		result := db.
			Joins("JOIN customers c ON work_orders.customer_id = c.id").
			Where("work_orders.id = uuid_to_bin(?)", customerID).
			First(&orderInfo.WorkOrder).
			Scan(&orderInfo)

		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(orderInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

}

package workOrder

import (
	"encoding/json"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/models"
	"time"
)

type dateRange struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func DateHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var workOrder []models.WorkOrder
		var dateRange dateRange
		err := json.NewDecoder(r.Body).Decode(&dateRange)

		db.Find(&workOrder, "created_at >= ? || created_at <= ?", dateRange.StartDate, dateRange.EndDate)

		jsonData, err := json.Marshal(workOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

package workOrder

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"net/http"
	"testUser/testEnerBit/database"
	"testUser/testEnerBit/models"
	"time"
)

func StoreHandler(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var workOrder models.WorkOrder
		err := json.NewDecoder(r.Body).Decode(&workOrder)
		isNotActive, err := validateDate(workOrder.PlannedDateBegin, workOrder.PlannedDateEnd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if isNotActive == true {
			changeIsNotActive(db, workOrder.CustomerID)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := db.Exec("INSERT INTO work_orders (customer_id, title, planned_date_begin, planned_date_end, status) VALUES (uuid_to_bin(?), ?, ?, ?, ?)", workOrder.CustomerID, workOrder.Title, workOrder.PlannedDateBegin, workOrder.PlannedDateEnd, workOrder.Status)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		updateCustomerIsActive(db, workOrder.CustomerID)

		w.WriteHeader(http.StatusCreated)
	}
}

func changeIsNotActive(db *database.DB, id uuid.UUID) {
	db.Model(&models.Customer{}).Where("id = uuid_to_bin(?)", id).Updates(map[string]interface{}{
		"is_active": false,
		"end_date":  time.Now(),
	})
}

func validateDate(begin time.Time, end time.Time) (bool, error) {
	if begin.After(end) {
		return false, errors.New("La fecha de inicio no puede ser posterior a la fecha final")
	}
	diff := end.Sub(begin)
	if diff > 2*time.Hour {
		return false, errors.New("La diferencia entre las fechas no puede ser mayor a dos horas")
	}
	if diff == 0*time.Hour {
		return true, nil
	}
	return false, nil
}

func updateCustomerIsActive(db *database.DB, id uuid.UUID) {
	db.Model(&models.Customer{}).Where("id = uuid_to_bin(?)", id).Updates(map[string]interface{}{
		"is_active":  true,
		"start_date": time.Now(),
	})
}

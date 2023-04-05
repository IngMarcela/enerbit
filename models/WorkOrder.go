package models

import (
	"github.com/google/uuid"
	"time"
)

type WorkOrder struct {
	ID               uuid.UUID `gorm:"type:binary(16);primary_key;default:(UUID_TO_BIN(UUID()))" json:"id"`
	CustomerID       uuid.UUID `gorm:"type:binary(16);primary_key" json:"customer_id"`
	Title            string    `gorm:"not null" json:"title"`
	PlannedDateBegin time.Time `gorm:"not null" json:"planned_date_begin"`
	PlannedDateEnd   time.Time `gorm:"not null" json:"planned_date_end"`
	Status           string    `gorm:"not null;enum:new,done,cancelled" json:"status"`
	CreatedAt        time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
}

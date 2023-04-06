package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"testUser/testEnerBit/models"
)

const dbHostEnvVar = "DB_HOST"

type DB struct {
	*gorm.DB
}

func NewDB() (*DB, error) {
	dbHost := os.Getenv(dbHostEnvVar)
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dsn := "user:password@tcp(" + dbHost + ":3306)/enerbit?timeout=15s&charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Migrate() error {
	err := db.AutoMigrate(&models.Customer{}, &models.WorkOrder{})
	if err != nil {
		return fmt.Errorf("Error al migrar tablas: %v", err)
	}
	return nil
}

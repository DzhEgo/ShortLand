package db

import (
	"ShortLand/internal/common/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
}

var db *gorm.DB

func Connect(conn string) (*gorm.DB, error) {
	var err error

	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	migrate()
	log.Println("Connected to database")

	return db, nil
}

func migrate() {
	db.AutoMigrate(&model.LinkTable{})
}

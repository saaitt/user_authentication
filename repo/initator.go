package repo

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {
	dbHost := "0.0.0.0"
	dbPort := "5432"
	pass := "123123"
	url := fmt.Sprintf("host=%s port=%s user=postgres dbname=auth_db password=%s sslmode=disable", dbHost, dbPort, pass)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	return db, err
}
package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionParams struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
}

func New(params ConnectionParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", params.Host, params.User, params.Password, params.DatabaseName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
